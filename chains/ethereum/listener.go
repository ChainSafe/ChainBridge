package ethereum

import (
	"ChainBridgeV2/core"
	eth "github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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

type Subscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}
