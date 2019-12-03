package ethereum

import (
	"ChainBridgeV2/core"
	"ChainBridgeV2/types"

	eth "github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ core.Listener = &Listener{}

type Listener struct {
	conn          *Connection
	subscriptions map[types.EventName]*Subscription
	handlers      map[types.EventName](func())
}

func NewListener(conn *Connection) *Listener {
	return &Listener{conn: conn}
}

type Subscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}

func (l *Listener) RegisterEventHandler(name types.EventName, handler func()) {
	l.handlers[name] = handler
}

func (l *Listener) subscribe(name types.EventName, q eth.FilterQuery) (*Subscription, error) {
	logChan := make(chan ethtypes.Log)
	ethsub, err := l.conn.conn.SubscribeFilterLogs(l.conn.ctx, q, logChan)
	if err != nil {
		return nil, err
	}

	sub := &Subscription{
		ch:  logChan,
		sub: ethsub,
	}

	l.subscriptions[name] = sub
	return sub, nil
}

func (l *Listener) Unsubscribe(name types.EventName) {
	l.subscriptions[name].sub.Unsubscribe()
}
