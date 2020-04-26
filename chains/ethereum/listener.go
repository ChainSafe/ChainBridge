// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
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

var BlockRetryInterval = time.Second * 2

type ActiveSubscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
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
}

func NewListener(conn *Connection, cfg *Config, log log15.Logger, bs blockstore.Blockstorer) *listener {
	return &listener{
		cfg:        *cfg,
		conn:       conn,
		log:        log,
		blockstore: bs,
	}
}

func (l *listener) setContracts(bridge *Bridge.Bridge, erc20Handler *ERC20Handler.ERC20Handler, erc721Handler *ERC721Handler.ERC721Handler, genericHandler *GenericHandler.GenericHandler) {
	l.bridgeContract = bridge
	l.erc20HandlerContract = erc20Handler
	l.erc721HandlerContract = erc721Handler
	l.genericHandlerContract = genericHandler
}

func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// Start registers all subscriptions provided by the config
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

//pollBlocks continously check the blocks for subscription logs, and sends messages to the router if logs are encountered
// stops where there are no subscriptions, and sleeps if we are at the current block
func (l *listener) pollBlocks() error {
	l.log.Debug("Polling Blocks...")
	var latestBlock = l.cfg.startBlock
	for {
		currBlock, err := l.conn.conn.BlockByNumber(l.conn.ctx, nil)
		if err != nil {
			return fmt.Errorf("unable to get latest block: %s", err)
		}
		if currBlock.Number().Cmp(latestBlock) < 0 {
			time.Sleep(BlockRetryInterval)
			continue
		}

		err = l.getDepositEventsForBlock(latestBlock)
		if err != nil {
			return err
		}

		err = l.blockstore.StoreBlock(latestBlock)
		if err != nil {
			return err
		}
		latestBlock.Add(latestBlock, big.NewInt(1))
	}
}

func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	query := buildQuery(l.cfg.bridgeContract, utils.Deposit, latestBlock, latestBlock)

	logs, err := l.conn.conn.FilterLogs(l.conn.ctx, query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %s", err)
	}

	for _, log := range logs {
		var m msg.Message
		addr := ethcommon.BytesToAddress(log.Topics[2].Bytes())
		destId := msg.ChainId(log.Topics[1].Big().Uint64())
		nonce := msg.Nonce(log.Topics[3].Big().Uint64())

		if addr == l.cfg.erc20HandlerContract {
			m = l.handleErc20DepositedEvent(destId, nonce)
		} else if addr == l.cfg.erc721HandlerContract {
			m = l.handleErc721DepositedEvent(destId, nonce)
		} else if addr == l.cfg.genericHandlerContract {
			m = l.handleGenericDepositedEvent(destId, nonce)
		} else {
			l.log.Error("Event has unrecognized handler", "handler", addr)
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

func (l *listener) stop() error {
	return nil
}
