// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	"github.com/ChainSafe/ChainBridge/blockstore"
	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("listener block polling failed")

type ActiveSubscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription // subscribe to specific events
}

type listener struct {
	cfg                    Config
	conn                   *Connection
	router                 chains.Router
	bridgeContract         *Bridge.Bridge // instance of bound bridge contract
	erc20HandlerContract   *ERC20Handler.ERC20Handler
	erc721HandlerContract  *ERC721Handler.ERC721Handler
	genericHandlerContract *GenericHandler.GenericHandler
	log                    log15.Logger
	blockstore             blockstore.Blockstorer
	stop                   <-chan int
	sysErr                 chan<- error // Reports fatal error to core
}

// NewListener creates and returns a listener
func NewListener(conn *Connection, cfg *Config, log log15.Logger, bs blockstore.Blockstorer, stop <-chan int, sysErr chan<- error) *listener {
	return &listener{
		cfg:        *cfg,
		conn:       conn,
		log:        log,
		blockstore: bs,
		stop:       stop,
		sysErr:     sysErr,
	}
}

// setContracts sets the listener with the appropriate contracts
func (l *listener) setContracts(bridge *Bridge.Bridge, erc20Handler *ERC20Handler.ERC20Handler, erc721Handler *ERC721Handler.ERC721Handler, genericHandler *GenericHandler.GenericHandler) {
	l.bridgeContract = bridge
	l.erc20HandlerContract = erc20Handler
	l.erc721HandlerContract = erc721Handler
	l.genericHandlerContract = genericHandler
}

// sets the router
func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start registers all subscriptions provided by the config
func (l *listener) start() error {
	l.log.Debug("Starting listener...")

	go func() {
		err := l.pollBlocks()
		if err != nil {
			l.log.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	l.log.Debug("Polling Blocks...")
	var latestBlock = l.cfg.startBlock
	var retry = BlockRetryLimit
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded")
				l.sysErr <- ErrFatalPolling
				return nil
			}

			currBlock, err := l.conn.latestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", latestBlock, "err", err)
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			// Sleep if the current block > latest
			if currBlock.Cmp(latestBlock) == -1 {
				l.log.Info("Block not ready, will retry", "target", latestBlock, "current", currBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			// Parse out events
			err = l.getDepositEventsForBlock(latestBlock)
			if err != nil {
				l.log.Error("Failed to get events for block", "block", latestBlock, "err", err)
				retry--
				continue
			}

			// Write to block store. Not a critical operation, no need to retry
			err = l.blockstore.StoreBlock(latestBlock)
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", latestBlock, "err", err)
			}

			// Goto next block and reset retry counter
			latestBlock.Add(latestBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Info("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(l.cfg.bridgeContract, utils.Deposit, latestBlock, latestBlock)

	// querying for logs
	logs, err := l.conn.conn.FilterLogs(l.conn.ctx, query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %s", err)
	}

	// read through the log events and handle their deposit event if handler is recognized
	for _, log := range logs {
		var m msg.Message
		addr := ethcommon.BytesToAddress(log.Topics[2].Bytes())
		destId := msg.ChainId(log.Topics[1].Big().Uint64())
		nonce := msg.Nonce(log.Topics[3].Big().Uint64())

		if addr == l.cfg.erc20HandlerContract {
			m, err = l.handleErc20DepositedEvent(destId, nonce)
		} else if addr == l.cfg.erc721HandlerContract {
			m, err = l.handleErc721DepositedEvent(destId, nonce)
		} else if addr == l.cfg.genericHandlerContract {
			m, err = l.handleGenericDepositedEvent(destId, nonce)
		} else {
			return fmt.Errorf("Event has unrecognized handler, handler %s", addr.Hex())
		}

		if err != nil {
			return err
		}

		err = l.router.Send(m)
		if err != nil {
			l.log.Error("subscription error: failed to route message", "err", err)
		}
	}

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig utils.EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
