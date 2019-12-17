package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/types"
)

var _ chains.Listener = &Listener{}

type Listener struct {
	conn chains.Connection
}

func NewListener(conn chains.Connection) *Listener {
	return &Listener{conn: conn}
}

func (l *Listener) RegisterEventHandler(name types.EventName, handler func()) {
	panic("not implemented")
}
