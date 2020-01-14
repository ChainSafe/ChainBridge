package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

var _ chains.Listener = &Listener{}

type Listener struct {
	conn chains.Connection
}

func NewListener(conn chains.Connection) *Listener {
	return &Listener{conn: conn}
}

func (l *Listener) Start() error {
	panic("not implemented")
}

func (l *Listener) SetRouter(chains.Router) {
	panic("implement me")
}

func (l *Listener) RegisterEventHandler(name string, handler func(interface{}) msg.Message) error {
	panic("not implemented")
}

func (l *Listener) Stop() error {
	panic("not implemented")
}
