package ethereum

import (
	"ChainBridgeV2/core"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ core.Listener = &Listener{}

type Listener struct {
	conn          *Connection
	home          common.Address
	away          common.Address
	subscriptions map[EventSig]*Subscription
	handlers      map[EventSig](func())
}

func NewListener(conn *Connection, cfg core.ChainConfig) *Listener {
	return &Listener{
		conn: conn,
		home: common.HexToAddress(cfg.Home),
		away: common.HexToAddress(cfg.Away),
	}
}

type Subscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}

func (l *Listener) RegisterEventHandler(sig string, handler func()) error {
	evt := EventSig(sig)
	_, _, err := l.conn.subscribeToEvent(l.home, evt)
	if err != nil {
		return err
	}
	l.handlers[evt] = handler
	return nil
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
