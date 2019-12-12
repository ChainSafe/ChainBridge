package centrifuge

import (
	"ChainBridgeV2/core"
)

var _ core.Listener = &Listener{}

type Listener struct {
	conn core.Connection
}

func NewListener(conn core.Connection) *Listener {
	return &Listener{conn: conn}
}

func (l *Listener) Start() error {
	panic("not implemented")
}

func (l *Listener) RegisterEventHandler(name string, handler func(interface{})) error {
	panic("not implemented")
}

func (l *Listener) Stop() error {
	panic("not implemented")
}
