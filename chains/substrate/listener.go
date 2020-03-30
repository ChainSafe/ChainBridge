// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"errors"
	"fmt"
	"time"

	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type Listener struct {
	name          string
	chainId       msg.ChainId
	startBlock    uint64
	conn          *Connection
	subscriptions map[eventName]eventHandler // Handlers for specific events
	router        chains.Router
}

// Frequency of polling for a new block
var BlockRetryInterval = time.Second * 2

func NewListener(conn *Connection, name string, id msg.ChainId, startBlock uint64) *Listener {
	return &Listener{
		name:          name,
		chainId:       id,
		startBlock:    startBlock,
		conn:          conn,
		subscriptions: make(map[eventName]eventHandler),
	}
}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

// Start creates the initial subscription for all events
func (l *Listener) Start() error {
	log15.Info("Starting listener...", "chain", l.name)

	// Check that whether latest is less than starting block
	header, err := l.conn.api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		return err
	}
	if uint64(header.Number) < l.startBlock {
		return fmt.Errorf("starting block (%d) is greater than latest known block (%d)", l.startBlock, header.Number)
	}

	for _, sub := range Subscriptions {
		err := l.RegisterEventHandler(sub.name, sub.handler)
		if err != nil {
			return err
		}
	}

	go func() {
		err := l.pollBlocks()
		if err != nil {
			log15.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

// RegisterEventHandler enables a handler for a given event. This cannot be used after Start is called.
func (l *Listener) RegisterEventHandler(name eventName, handler eventHandler) error {
	if l.subscriptions[name] != nil {
		return fmt.Errorf("event %s already registered", name)
	}
	l.subscriptions[name] = handler
	return nil
}

var ErrBlockNotReady = errors.New("required result to be 32 bytes, but got 0")

func (l *Listener) pollBlocks() error {
	var latestBlock = l.startBlock
	for {
		log15.Trace("Polling for block", "number", latestBlock)
		hash, err := l.conn.api.RPC.Chain.GetBlockHash(latestBlock)
		if err != nil && err.Error() == ErrBlockNotReady.Error() {
			log15.Trace("Block not ready, sleeping...", "interval", BlockRetryInterval)
			time.Sleep(BlockRetryInterval)
			continue
		} else if err != nil {
			return err
		}
		err = l.processEvents(hash)
		if err != nil {
			return err
		}
		latestBlock++
	}
}

// processEvents fetches a block and parses out the events, calling Listener.handleEvents()
func (l *Listener) processEvents(hash types.Hash) error {
	log15.Trace("Fetching block", "hash", hash.Hex())
	meta := l.conn.getMetadata()
	key, err := types.CreateStorageKey(&meta, "System", "Events", nil, nil)
	if err != nil {
		return err
	}

	var records types.EventRecordsRaw
	_, err = l.conn.api.RPC.State.GetStorage(key, &records, hash)
	if err != nil {
		return err
	}

	e := Events{}
	err = records.DecodeEventRecords(&meta, &e)
	if err != nil {
		return err
	}

	l.handleEvents(e)
	log15.Trace("Finished processing events", "block", hash.Hex())

	return nil
}

// handleEvents calls the associated handler for all registered event types
func (l *Listener) handleEvents(evts Events) {
	if l.subscriptions[RelayerThresholdSet] != nil {
		for _, evt := range evts.Bridge_RelayerThresholdSet {
			log15.Trace("Received RelayerThreshold event", "threshold", evt.Threshold)
		}
	}
	if l.subscriptions[ChainWhitelisted] != nil {
		for _, evt := range evts.Bridge_ChainWhitelisted {
			log15.Trace("Received ChainWhitelisted event", "chainId", evt.ChainId)
		}
	}
	if l.subscriptions[RelayerAdded] != nil {
		for _, evt := range evts.Bridge_RelayerAdded {
			log15.Trace("Received RelayerAdded event", "relayer", evt.Relayer.Hex())
		}
	}
	if l.subscriptions[RelayerRemoved] != nil {
		for _, evt := range evts.Bridge_RelayerRemoved {
			log15.Trace("Received RelayerRemoved event", "relayer", evt.Relayer.Hex())
		}
	}
	if l.subscriptions[AssetTransfer] != nil {
		for _, evt := range evts.Bridge_AssetTransfer {
			log15.Trace("Handling AssetTransfer event")
			l.submitMessage(l.subscriptions[AssetTransfer](evt))
		}
	}
	if l.subscriptions[VoteFor] != nil {
		for _, evt := range evts.Bridge_VoteFor {
			log15.Trace("Received VoteFor event", "depositNonce", evt.DepositNonce)
		}
	}
	if l.subscriptions[VoteAgainst] != nil {
		for _, evt := range evts.Bridge_VoteAgainst {
			log15.Trace("Received VoteAgainst event", "depositNonce", evt.DepositNonce)
		}
	}
	if l.subscriptions[ProposalApproved] != nil {
		for _, evt := range evts.Bridge_ProposalApproved {
			log15.Trace("Received ProposalApproved event", "depositNonce", evt.DepositNonce)
		}
	}
	if l.subscriptions[ProposalRejected] != nil {
		for _, evt := range evts.Bridge_ProposalRejected {
			log15.Trace("Received ProposalRejected event", "depositNonce", evt.DepositNonce)
		}
	}
	if l.subscriptions[ProposalSucceeded] != nil {
		for _, evt := range evts.Bridge_ProposalSucceeded {
			log15.Trace("Received ProposalSucceeded event", "depositNonce", evt.DepositNonce)
		}
	}
	if l.subscriptions[ProposalFailed] != nil {
		for _, evt := range evts.Bridge_ProposalFailed {
			log15.Trace("Received ProposalFailed event", "depositNonce", evt.DepositNonce)
		}
	}
}

// submitMessage inserts the chainId into the msg and sends it to the router
func (l *Listener) submitMessage(m msg.Message) {
	m.Source = l.chainId
	err := l.router.Send(m)
	if err != nil {
		log15.Error("failed to process event", "err", err)
	}
}

func (l *Listener) Stop() error {
	return nil
}
