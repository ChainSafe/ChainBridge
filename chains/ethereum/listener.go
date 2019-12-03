package ethereum

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

func (l *Listener) RegisterEventHandler() {
	panic("not implemented")
}
