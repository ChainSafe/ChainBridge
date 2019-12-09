package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

func InitializeChain(id msg.ChainId, cfg *core.ChainConfig) *core.Chain {
	c := core.NewChain(id, cfg)
	c.SetConnection(NewConnection(cfg.Endpoint))
	c.SetListener(NewListener(c.Connection()))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}
