// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/ChainBridge/blockstore"
	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type listener struct {
	name          string
	chainId       msg.ChainId
	startBlock    uint64
	blockstore    blockstore.Blockstorer
	conn          *Connection
	subscriptions map[eventName]eventHandler // Handlers for specific events
	router        chains.Router
	log           log15.Logger
}

// Frequency of polling for a new block
var BlockRetryInterval = time.Second * 2
var BlockRetryLimit = 3

func NewListener(conn *Connection, name string, id msg.ChainId, startBlock uint64, log log15.Logger, bs blockstore.Blockstorer) *listener {
	return &listener{
		name:          name,
		chainId:       id,
		startBlock:    startBlock,
		blockstore:    bs,
		conn:          conn,
		subscriptions: make(map[eventName]eventHandler),
		log:           log,
	}
}

func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start creates the initial subscription for all events
func (l *listener) start() error {
	// Check that whether latest is less than starting block
	header, err := l.conn.api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		return err
	}
	if uint64(header.Number) < l.startBlock {
		return fmt.Errorf("starting block (%d) is greater than latest known block (%d)", l.startBlock, header.Number)
	}

	for _, sub := range Subscriptions {
		err := l.registerEventHandler(sub.name, sub.handler)
		if err != nil {
			return err
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

// registerEventHandler enables a handler for a given event. This cannot be used after Start is called.
func (l *listener) registerEventHandler(name eventName, handler eventHandler) error {
	if l.subscriptions[name] != nil {
		return fmt.Errorf("event %s already registered", name)
	}
	l.subscriptions[name] = handler
	return nil
}

var ErrBlockNotReady = errors.New("required result to be 32 bytes, but got 0")

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	var latestBlock = l.startBlock
	var retry = BlockRetryLimit
	for {
		// No more retries, goto next block
		if retry == 0 {
			latestBlock++
			retry = BlockRetryLimit
		}

		// Get hash for latest block, sleep and retry if not ready
		hash, err := l.conn.api.RPC.Chain.GetBlockHash(latestBlock)
		if err != nil && err.Error() == ErrBlockNotReady.Error() {
			time.Sleep(BlockRetryInterval)
			continue
		} else if err != nil {
			l.log.Error("Failed to query latest block", "block", latestBlock, "err", err)
			retry--
			time.Sleep(BlockRetryInterval)
			continue
		}

		err = l.processEvents(hash)
		if err != nil {
			l.log.Error("Failed to process events in block", "block", latestBlock, "err", err)
			retry--
			continue
		}

		// Write to blockstore
		err = l.blockstore.StoreBlock(big.NewInt(0).SetUint64(latestBlock))
		if err != nil {
			l.log.Error("Failed to write to blockstore", "err", err)
		}
		latestBlock++
		retry = BlockRetryLimit
	}
}

// processEvents fetches a block and parses out the events, calling Listener.handleEvents()
func (l *listener) processEvents(hash types.Hash) error {
	l.log.Trace("Fetching block for events", "hash", hash.Hex())
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
	l.log.Trace("Finished processing events", "block", hash.Hex())

	return nil
}

// handleEvents calls the associated handler for all registered event types
func (l *listener) handleEvents(evts Events) {
	if l.subscriptions[FungibleTransfer] != nil {
		for _, evt := range evts.ChainBridge_FungibleTransfer {
			l.log.Trace("Handling FungibleTransfer event")
			l.submitMessage(l.subscriptions[FungibleTransfer](evt, l.log))
		}
	}
	if l.subscriptions[NonFungibleTransfer] != nil {
		for _, evt := range evts.ChainBridge_NonFungibleTransfer {
			l.log.Trace("Handling NonFungibleTransfer event")
			l.submitMessage(l.subscriptions[NonFungibleTransfer](evt, l.log))
		}
	}
	if l.subscriptions[GenericTransfer] != nil {
		for _, evt := range evts.ChainBridge_GenericTransfer {
			l.log.Trace("Handling GenericTransfer event")
			l.submitMessage(l.subscriptions[GenericTransfer](evt, l.log))
		}
	}

	if len(evts.System_CodeUpdated) > 0 {
		l.log.Trace("Received CodeUpdated event")
		err := l.conn.updateMetatdata()
		if err != nil {
			l.log.Error("Unable to update Metadata", "error", err)
		}
	}
}

// submitMessage inserts the chainId into the msg and sends it to the router
func (l *listener) submitMessage(m msg.Message, err error) {
	if err != nil {
		log15.Error("Critical error processing event", "err", err)
		return
	}
	m.Source = l.chainId
	err = l.router.Send(m)
	if err != nil {
		log15.Error("failed to process event", "err", err)
	}
}

func (l *listener) Stop() error {
	return nil
}
