package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
)

func InitializeChain(cfg *core.ChainConfig) *core.Chain {
	c := core.NewChain(cfg)
	conn := NewConnection(cfg.Endpoint)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn, *cfg))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}
