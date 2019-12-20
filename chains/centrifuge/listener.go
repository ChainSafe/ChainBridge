package centrifuge

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var _ chains.Listener = &Listener{}

type Listener struct {
	cfg core.ChainConfig
	conn *Connection
	sub *state.StorageSubscription
	subscriptions map[string]chains.HandlerFn
}

type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	Nfts_DepositAsset []EventNFTDeposited //nolint:stylecheck,golint
}

func NewListener(conn *Connection, cfg core.ChainConfig) *Listener {
	return &Listener{
		cfg: cfg,
		conn: conn,
		subscriptions: make(map[string]chains.HandlerFn),
	}
}

// Start creates the initial subscription for all events
func (l *Listener) Start() error {
	log15.Info("Starting centrifuge listener...", "chainID", l.cfg.Id, "subs", l.cfg.Subscriptions)
	sub, err := l.conn.Subscribe()
	if err != nil {
		return err
	}
	l.sub = sub

	return nil
}

func (l *Listener) RegisterEventHandler(name string, handler chains.HandlerFn) error {
	if l.subscriptions[name] != nil {
		return fmt.Errorf("event %s already registered", name)
	}
	l.subscriptions[name] = handler
	return nil
}

func (l *Listener) watchForEvents(sub *state.StorageSubscription, handler func(interface{}) msg.Message) {
	for {
		select {
		case evt := <-sub.Chan():
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
				log15.Info("Got event!", "evt", evt)
			}
		case err := <-sub.Err():
			log15.Error("subscription error", "sub", sub, "err", err)
		}
	}
}

func (l *Listener) handleEvents(evts Events) {
	// TODO: Iterate through events and handle those we are subscribed to

}

func (l *Listener) Stop() error {
	panic("not implemented")
}
