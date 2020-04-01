// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridge/chains"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

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
	subscriptions        map[EventSig]*ActiveSubscription
	router               chains.Router
	bridgeContract       *BridgeContract       // instance of bound bridge contract
	erc20HandlerContract *ERC20HandlerContract // instance of bound erc20 handler
	log                  log15.Logger
}

func NewListener(conn *Connection, cfg *Config, log log15.Logger) *Listener {
	return &Listener{
		cfg:           *cfg,
		conn:          conn,
		subscriptions: make(map[EventSig]*ActiveSubscription),
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
	return nil
}

// buildQuery constructs a query for the contract by hashing sig to get the event topic
// TODO: Start from current block
func buildQuery(contract ethcommon.Address, sig EventSig, startBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

// RegisterEventHandler creates a subscription for the provided event on the bridge contract.
// Handler will be called for every instance of event.
func (l *Listener) RegisterEventHandler(subscription EventSig, handler evtHandlerFn) error {
	evt := subscription
	query := buildQuery(l.cfg.contract, evt, big.NewInt(0))
	eventSubscription, err := l.conn.subscribeToEvent(query)
	if err != nil {
		return err
	}
	l.subscriptions[subscription] = eventSubscription
	go l.watchEvent(eventSubscription, handler)
	l.log.Debug("Registered event handler", "contract", l.cfg.contract, "sig", subscription)
	return nil
}

// watchEvent will call the handler for every occurrence of the corresponding event. It should be run in a separate
// goroutine to monitor the subscription channel.
func (l *Listener) watchEvent(eventSubscription *ActiveSubscription, handler evtHandlerFn) {
	for {
		select {
		case evt := <-eventSubscription.ch:
			m := handler(evt)
			err := l.router.Send(m)
			if err != nil {
				l.log.Error("subscription error: cannot send message", "sub", eventSubscription, "err", err)
			}
		case err := <-eventSubscription.sub.Err():
			if err != nil {
				l.log.Error("subscription error", "err", err)
			}
		}
	}
}

// Unsubscribe cancels a subscription for the given event
func (l *Listener) Unsubscribe(sig EventSig) {
	if _, ok := l.subscriptions[sig]; ok {
		l.subscriptions[sig].sub.Unsubscribe()
	}
}

// Stop cancels all subscriptions. Must be called before Connection.Stop().
func (l *Listener) Stop() error {
	for _, sub := range l.subscriptions {
		sub.sub.Unsubscribe()
	}
	return nil
}
