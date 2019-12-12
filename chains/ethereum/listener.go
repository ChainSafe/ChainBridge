package ethereum

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"

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
	log15.Info("Starting ethereum listener...", "subs", l.cfg.Subscriptions)
	for _, sub := range l.cfg.Subscriptions {
		err := l.RegisterEventHandler(sub, func(evtI interface{}) {
			evt := evtI.(*geth.Log)
			log15.Info("Got event!", "evt", evt)
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

func (l *Listener) RegisterEventHandler(sig string, handler func(interface{}) msg.Message) error {
	log15.Info("Registering event handler", "sig", sig)
	evt := EventSig(sig)
	sub, err := l.conn.subscribeToEvent(l.home, evt)
	if err != nil {
		return err
	}
	watchEvent(sub, handler)
	return nil
}

func watchEvent(sub *Subscription, handler func(interface{})) {
	for {
		select {
		case evt := <-sub.ch:
			handler(evt)
		case err := <-sub.sub.Err():
			log15.Error("subscription error", "sub", sub, "err", err)
		}
	}
}

//func (l *Listener) subscribe(name types.EventName, q eth.FilterQuery) (*Subscription, error) {
//	logChan := make(chan ethtypes.Log)
//	ethsub, err := l.conn.conn.SubscribeFilterLogs(l.conn.ctx, q, logChan)
//	if err != nil {
//		return nil, err
//	}
//
//	sub := &Subscription{
//		ch:  logChan,
//		sub: ethsub,
//	}
//
//	l.subscriptions[name] = sub
//	return sub, nil
//}

func (l *Listener) Unsubscribe(sig EventSig) {
	l.subscriptions[sig].sub.Unsubscribe()
}

func (l *Listener) Stop() error {
	for _, sub := range l.subscriptions {
		sub.sub.Unsubscribe()
	}
	return nil
}
