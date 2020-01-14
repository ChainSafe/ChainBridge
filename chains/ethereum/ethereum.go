package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
)

func InitializeChain(chainCfg *core.ChainConfig) *core.Chain {
	cfg := ParseChainConfig(chainCfg)

	c := core.NewChain(chainCfg)
	conn := NewConnection(cfg)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn, cfg))
	c.SetWriter(NewWriter(conn, cfg))

	return c
}
