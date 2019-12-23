package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ chains.Listener = &Listener{}

type Listener struct {
	cfg           *core.ChainConfig
	conn          *Connection
	receiver      common.Address
	emitter       common.Address
	subscriptions map[EventSig]*Subscription
	router        chains.Router
	//handlers      map[EventSig](func())
}

func NewListener(conn *Connection, cfg *core.ChainConfig) *Listener {
	return &Listener{
		cfg:           cfg,
		conn:          conn,
		receiver:      common.HexToAddress(cfg.Receiver),
		emitter:       common.HexToAddress(cfg.Emitter),
		subscriptions: make(map[EventSig]*Subscription),
	}
}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

func (l *Listener) Start() error {
	log15.Info("Starting listener...", "chainID", l.cfg.Id, "subs", l.cfg.Subscriptions)
	for _, sub := range l.cfg.Subscriptions {
		err := l.RegisterEventHandler(sub, func(evtI interface{}) msg.Message {
			evt := evtI.(ethtypes.Log)
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
	evt := EventSig(sig)
	query := l.buildQuery(l.emitter, evt)
	sub, err := l.conn.subscribeToEvent(query)
	if err != nil {
		return err
	}
	l.subscriptions[EventSig(sig)] = sub
	go l.watchEvent(sub, handler)
	log15.Info("Registered event handler", "chainID", l.cfg.Id, "contract", l.emitter, "sig", sig)
	return nil
}

func (l *Listener) watchEvent(sub *Subscription, handler func(interface{}) msg.Message) {
	for {
		select {
		case evt := <-sub.ch:
			m := handler(evt)
			err := l.router.Send(m)
			if err != nil {
				log15.Error("subscription error: cannot send message", "sub", sub, "err", err)
			}
		case err := <-sub.sub.Err():
			if err != nil {
				log15.Error("subscription error", "sub", sub, "err", err)
			}
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
