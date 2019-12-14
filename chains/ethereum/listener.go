package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"

	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	geth "github.com/ethereum/go-ethereum/mobile"
)

var _ core.Listener = &Listener{}

type Listener struct {
	cfg           core.ChainConfig
	conn          *Connection
	home          common.Address
	away          common.Address
	subscriptions map[EventSig]*Subscription
	//handlers      map[EventSig](func())
}

func NewListener(conn *Connection, cfg core.ChainConfig) *Listener {
	return &Listener{
		cfg:  cfg,
		conn: conn,
		home: common.HexToAddress(cfg.Home),
		away: common.HexToAddress(cfg.Away),
	}
}

func (l *Listener) Start() error {
	log15.Info("Starting ethereum listener...", "chainID", l.cfg.Id, "subs", l.cfg.Subscriptions)
	for _, sub := range l.cfg.Subscriptions {
		err := l.RegisterEventHandler(sub, func(evtI interface{}) msg.Message {
			evt := evtI.(*geth.Log)
			log15.Info("Got event!", "evt", evt)
			return msg.Message{}
		})
		if err != nil {
			log15.Error("failed to register event handler", "err", err)
		}
	}
	return nil
}

type Subscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}

// buildQuery constructs a query for the contract by hashing sig to get the event topic
func (l *Listener) buildQuery(contract common.Address, sig EventSig) eth.FilterQuery {
	query := eth.FilterQuery{
		// TODO: Might want current block
		FromBlock: nil,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}


func (l *Listener) RegisterEventHandler(sig string, handler func(interface{}) msg.Message) error {
	log15.Info("Registering event handler", "sig", sig)
	evt := EventSig(sig)
	query := l.buildQuery(l.home, evt)
	sub, err := l.conn.subscribeToEvent(query, evt)
	if err != nil {
		return err
	}
	// TODO: Should be go routine
	watchEvent(sub, handler)
	return nil
}

func watchEvent(sub *Subscription, handler func(interface{}) msg.Message) {
	for {
		select {
		case evt := <-sub.ch:
			handler(evt)
		case err := <-sub.sub.Err():
			log15.Error("subscription error", "sub", sub, "err", err)
		}
	}
}

func (l *Listener) Unsubscribe(sig EventSig) {
	l.subscriptions[sig].sub.Unsubscribe()
}

func (l *Listener) Stop() error {
	for _, sub := range l.subscriptions {
		sub.sub.Unsubscribe()
	}
	return nil
}
