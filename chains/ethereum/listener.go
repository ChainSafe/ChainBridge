// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/ChainBridge/chains"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var BlockRetryInterval = time.Second * 2

type Subscription struct {
	signature EventSig
	handler   evtHandlerFn
}

type ActiveSubscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}

type Listener struct {
	cfg                  Config
	conn                 *Connection
	subscriptions        map[EventSig]*Subscription
	router               chains.Router
	bridgeContract       *BridgeContract       // instance of bound bridge contract
	erc20HandlerContract *ERC20HandlerContract // instance of bound erc20 handler
	log                  log15.Logger
}

func NewListener(conn *Connection, cfg *Config, log log15.Logger) *Listener {
	return &Listener{
		cfg:           *cfg,
		conn:          conn,
		subscriptions: make(map[EventSig]*Subscription),
		log:           log,
	}
}

func (l *Listener) SetContracts(bridge *BridgeContract, erc20Handler *ERC20HandlerContract) {
	l.bridgeContract = bridge
	l.erc20HandlerContract = erc20Handler

}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

func (l *Listener) GetSubscriptions() []*Subscription {
	return []*Subscription{
		{
			signature: Deposit,
			handler:   l.handleErc20DepositedEvent,
		},
	}

}

// Start registers all subscriptions provided by the config
func (l *Listener) Start() error {
	l.log.Debug("Starting listener...")

	subscriptions := l.GetSubscriptions()
	for _, sub := range subscriptions {
		err := l.RegisterEventHandler(sub.signature, sub.handler)
		if err != nil {
			l.log.Error("failed to register event handler", "err", err)
		}
	}

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
func (l *Listener) pollBlocks() error {
	l.log.Debug("Polling Blocks...")
	var latestBlock = l.cfg.startBlock
	for {
		if len(l.subscriptions) == 0 {
			l.log.Debug("No subscriptions, stopping Polling")
			break
		}
		currBlock, err := l.conn.conn.BlockByNumber(l.conn.ctx, nil)
		if err != nil {
			return fmt.Errorf("Unable to get latest block: %s", err)
		}
		if currBlock.Number().Cmp(latestBlock) < 0 {
			time.Sleep(BlockRetryInterval)
			continue
		}

		err = l.getEventsForBlock(latestBlock)
		if err != nil {
			return err
		}

		latestBlock.Add(latestBlock, big.NewInt(1))
	}

	return nil
}

func (l *Listener) getEventsForBlock(latestBlock *big.Int) error {

	for evt, sub := range l.subscriptions {
		query := buildQuery(l.cfg.bridgeContract, evt, latestBlock, latestBlock)

		logs, err := l.conn.conn.FilterLogs(l.conn.ctx, query)
		if err != nil {
			return fmt.Errorf("Unable to Filter Logs: %s", err)
		}

		for _, log := range logs {
			m := sub.handler(log)
			err = l.router.Send(m)
			if err != nil {
				l.log.Error("subscription error: cannot send message", "sub", sub, "err", err)
			}
		}
	}

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
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

// RegisterEventHandler creates a subscription for the provided event on the bridge bridgeContract.
// Handler will be called for every instance of event.
func (l *Listener) RegisterEventHandler(subscription EventSig, handler evtHandlerFn) error {

	if l.subscriptions[subscription] != nil {
		return fmt.Errorf("event %s already registered", subscription)
	}
	sub := new(Subscription)
	sub.handler = handler
	sub.signature = subscription
	l.subscriptions[subscription] = sub

	return nil

}

// Stop cancels all subscriptions. Must be called before Connection.Stop().
func (l *Listener) Stop() error {
	for sig := range l.subscriptions {
		delete(l.subscriptions, sig)
	}
	return nil
}
