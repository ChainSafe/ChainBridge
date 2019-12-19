package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
)

func InitializeChain(cfg *core.ChainConfig) *core.Chain {
	c := core.NewChain(cfg)
	c.SetConnection(NewConnection(cfg.Endpoint))
	c.SetListener(NewListener(c.Connection()))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}
