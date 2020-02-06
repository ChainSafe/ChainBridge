package centrifuge

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var _ chains.Listener = &Listener{}

type Listener struct {
	cfg           core.ChainConfig
	conn          *Connection
	sub           *state.StorageSubscription     // Subscription to all events
	subscriptions map[string]chains.EvtHandlerFn // Handlers for specific events
	router        chains.Router
}

func NewListener(conn *Connection, cfg core.ChainConfig) *Listener {
	return &Listener{
		cfg:           cfg,
		conn:          conn,
		subscriptions: make(map[string]chains.EvtHandlerFn),
	}
}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

// Start creates the initial subscription for all events
func (l *Listener) Start() error {
	log15.Info("Starting centrifuge listener...", "chainID", l.cfg.Id, "subs", l.cfg.Subscriptions)

	for _, sub := range l.cfg.Subscriptions {
		switch sub {
		case "nfts":
			err := l.RegisterEventHandler("nfts", nftHandler)

			if err != nil {
				return err
			}
		case "assetTx":
			err := l.RegisterEventHandler("assetTx", assetTransferHandler)

			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("unrecognized event: %s", sub)
		}

	}

	sub, err := l.conn.Subscribe()
	if err != nil {
		return err
	}
	l.sub = sub

	go l.watchForEvents(sub)
	return nil
}

// RegisterEventHandler enables a handler for a given event. This cannot be used after Start is called.
func (l *Listener) RegisterEventHandler(name string, handler chains.EvtHandlerFn) error {
	if l.sub == nil {
		if l.subscriptions[name] != nil {
			return fmt.Errorf("event %s already registered", name)
		}
		l.subscriptions[name] = handler
		return nil
	}
	return fmt.Errorf("can't register handler once listener is started")

}

// watchForEvents is a blocking function that watches the subscription's event and error channels.
// New events are handled by calling handleEvents. Errors are simply logged.
func (l *Listener) watchForEvents(sub *state.StorageSubscription) {
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
	// TODO: Handle all event types
	if l.subscriptions["nfts"] != nil {
		for _, nft := range evts.Nfts_DepositAsset {
			log15.Trace("Handling NFT event")
			msg := l.subscriptions["nfts"](nft)
			err := l.router.Send(msg)
			if err != nil {
				log15.Error("failed to process event", "err", err)
			}
		}
	}

	if l.subscriptions["assetTx"] != nil {
		for _, assetTx := range evts.Bridge_AssetTransfer {
			log15.Trace("Handling AssetTransfer event")
			msg := l.subscriptions["assetTx"](assetTx)
			err := l.router.Send(msg)
			if err != nil {
				log15.Error("failed to process event", "err", err)
			}
		}
	}

}

func (l *Listener) Stop() error {
	l.sub.Unsubscribe()
	return nil
}
