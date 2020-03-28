// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"errors"
	"fmt"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

type Listener struct {
	cfg           *core.ChainConfig
	conn          *Connection
	sub           *state.StorageSubscription // Subscription to all events
	subscriptions map[eventName]eventHandler // Handlers for specific events
	router        chains.Router
}

func NewListener(conn *Connection, cfg *core.ChainConfig) *Listener {
	return &Listener{
		cfg:           cfg,
		conn:          conn,
		subscriptions: make(map[eventName]eventHandler),
	}
}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

// Start creates the initial subscription for all events
func (l *Listener) Start() error {
	log15.Info("Starting substrate listener...", "chainID", l.cfg.Id, "subs", Subscriptions)

	for _, sub := range Subscriptions {
		switch sub {
		case "nfts":
			err := l.RegisterEventHandler("nfts", nftHandler)
			if err != nil {
				return err
			}
		case AssetTx:
			err := l.RegisterEventHandler(AssetTx, assetTransferHandler)
			if err != nil {
				return err
			}
		case ValidatorAdded:
			err := l.RegisterEventHandler(ValidatorAdded, validatorAddedHandler)
			if err != nil {
				return err
			}
		case ValidatorRemoved:
			err := l.RegisterEventHandler(ValidatorRemoved, validatorRemovedHandler)
			if err != nil {
				return err
			}
		case VoteFor:
			err := l.RegisterEventHandler(VoteFor, voteForHandler)
			if err != nil {
				return err
			}
		case VoteAgainst:
			err := l.RegisterEventHandler(VoteAgainst, voteAgainstHandler)
			if err != nil {
				return err
			}
		case ProposalSucceeded:
			err := l.RegisterEventHandler(ProposalSucceeded, proposalSucceededHandler)
			if err != nil {
				return err
			}
		case ProposalFailed:
			err := l.RegisterEventHandler(ProposalFailed, proposalFailedHandler)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unrecognized event: %s", sub)
		}

	}

	log15.Trace("Registered event handlers", "events", l.subscriptions)

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
	if l.sub == nil {
		if l.subscriptions[name] != nil {
			return fmt.Errorf("event %s already registered", name)
		}
		l.subscriptions[name] = handler
		return nil
	}
	return fmt.Errorf("can't register handler once listener is started")

}

var ErrBlockNotReady = errors.New("required result to be 32 bytes, but got 0")
var BlockRetryInterval = time.Second * 1

func (l *Listener) pollBlocks() error {
	var latestBlock uint64 = 0
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

func (l *Listener) processEvents(hash types.Hash) error {
	log15.Trace("Fetching block", "hash", hash)
	data := l.conn.getMetadata()
	key, err := types.CreateStorageKey(&data, "System", "Events", nil, nil)
	if err != nil {
		return err
	}

	var records types.EventRecordsRaw
	_, err = l.conn.api.RPC.State.GetStorage(key, &records, hash)
	if err != nil {
		return err
	}

	e := Events{}
	err = records.DecodeEventRecords(&data, &e)
	if err != nil {
		return err
	}

	l.handleEvents(e)
	log15.Trace("Finished processing events", "block", hash)

	return nil
}

// watchForEvents is a blocking function that watches the subscription's event and error channels.
// New events are handled by calling handleEvents. Errors are simply logged.
func (l *Listener) watchForEvents(sub *state.StorageSubscription) { //nolint:unused
	for {
		select {
		case evt := <-sub.Chan():
			log15.Trace("Received new block", "chainID", l.cfg.Id)
			for _, chng := range evt.Changes {
				events := Events{}
				meta, err := l.conn.api.RPC.State.GetMetadataLatest()
				if err != nil {
					log15.Error("Failed to get metadata", "err", err)
				}
				err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(meta, &events)
				if err != nil {
					panic(err)
				}
				l.handleEvents(events)
			}
		case err := <-sub.Err():
			// TODO: Re-try connection
			if err != nil {
				log15.Error("subscription error", "sub", sub, "err", err)
			}
		}
	}
}

// handleEvents calls the associated handler for all registered event types
func (l *Listener) handleEvents(evts Events) {
	if l.subscriptions[AssetTx] != nil {
		for _, assetTx := range evts.Bridge_AssetTransfer {
			log15.Trace("Handling AssetTransfer event")
			_ = l.subscriptions[AssetTx](assetTx)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}

	if l.subscriptions[ValidatorAdded] != nil {
		for _, v := range evts.Bridge_ValidatorAdded {
			log15.Trace("Handling ValidatorAdded event")
			_ = l.subscriptions[ValidatorAdded](v)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}
	if l.subscriptions[ValidatorRemoved] != nil {
		for _, v := range evts.Bridge_ValidatorRemoved {
			log15.Trace("Handling ValidatorRemoved event")
			_ = l.subscriptions[ValidatorRemoved](v)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}
	if l.subscriptions[VoteFor] != nil {
		for _, e := range evts.Bridge_VoteFor {
			log15.Trace("Handling AssetTransfer event")
			_ = l.subscriptions[VoteFor](e)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}
	if l.subscriptions[VoteAgainst] != nil {
		for _, e := range evts.Bridge_VoteAgainst {
			log15.Trace("Handling AssetTransfer event")
			_ = l.subscriptions[VoteAgainst](e)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}
	if l.subscriptions[ProposalSucceeded] != nil {
		for _, e := range evts.Bridge_ProposalSucceeded {
			log15.Trace("Handling ProposalSucceeded event")
			_ = l.subscriptions[ProposalSucceeded](e)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}
	if l.subscriptions[ProposalFailed] != nil {
		for _, e := range evts.Bridge_ProposalFailed {
			log15.Trace("Handling ProposalFailed event")
			_ = l.subscriptions[ProposalFailed](e)
			//err := l.router.Send(msg)
			//if err != nil {
			//	log15.Error("failed to process event", "err", err)
			//}
		}
	}
}

func (l *Listener) Stop() error {
	l.sub.Unsubscribe()
	return nil
}
