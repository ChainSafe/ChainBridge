// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ chains.Listener = &Listener{}

type Subscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}

type Listener struct {
	cfg             Config
	conn            *Connection
	subscriptions   map[EventSig]*Subscription
	router          chains.Router
	emitterContract EmitterContract // instance of bound emitter contract
}

func NewListener(conn *Connection, cfg *Config) *Listener {
	return &Listener{
		cfg:           *cfg,
		conn:          conn,
		subscriptions: make(map[EventSig]*Subscription),
	}
}

func (l *Listener) SetEmitterContract(ec EmitterContract) {
	l.emitterContract = ec
}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

// Start registers all subscriptions provided by the config
func (l *Listener) Start() error {
	log15.Debug("Starting listener...", "chainID", l.cfg.id, "subs", l.cfg.subscriptions)
	for _, subscription := range l.cfg.subscriptions {
		sub := subscription
		switch sub {
		case ErcTransferSignature, NftTransferSignature:
			err := l.RegisterEventHandler(sub, l.handleTransferEvent)
			if err != nil {
				log15.Error("failed to register event handler", "err", err)
			}
		case DepositAssetSignature:
			err := l.RegisterEventHandler(sub, l.handleTestDeposit)
			if err != nil {
				log15.Error("failed to register event handler", "err", err)
			}
		default:
			return fmt.Errorf("Unrecognized event: %s", sub)
		}
	}
	return nil
}

// buildQuery constructs a query for the contract by hashing sig to get the event topic
// TODO: Start from current block
func (l *Listener) buildQuery(contract ethcommon.Address, sig EventSig) eth.FilterQuery {

	query := eth.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

// RegisterEventHandler creates a subscription for the provided event on the emitter contract.
// Handler will be called for every instance of event.
func (l *Listener) RegisterEventHandler(subscription string, handler chains.EvtHandlerFn) error {
	evt := EventSig(subscription)
	query := l.buildQuery(l.cfg.emitter, evt)
	eventSubscription, err := l.conn.subscribeToEvent(query)
	if err != nil {
		return err
	}
	l.subscriptions[EventSig(subscription)] = eventSubscription
	go l.watchEvent(eventSubscription, handler)
	log15.Debug("Registered event handler", "chainID", l.cfg.id, "contract", l.cfg.emitter, "sig", subscription)
	return nil
}

// watchEvent will call the handler for every occurrence of the corresponding event. It should be run in a separate
// goroutine to monitor the subscription channel.
func (l *Listener) watchEvent(eventSubscription *Subscription, handler func(interface{}) msg.Message) {
	for {
		select {
		case evt := <-eventSubscription.ch:
			log15.Trace("Event found")
			m := handler(evt)
			err := l.router.Send(m)
			if err != nil {
				log15.Error("subscription error: cannot send message", "sub", eventSubscription, "err", err)
			}
		case err := <-eventSubscription.sub.Err():
			if err != nil {
				log15.Error("subscription error", "sub", eventSubscription, "err", err)
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
