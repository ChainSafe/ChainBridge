package ethereum

import (
	"context"

	"github.com/ChainSafe/ChainBridgeV2/core"
)

func InitializeChain(cfg *core.ChainConfig) *core.Chain {
	ctx := context.Background()
	c := core.NewChain(cfg)

	conn := NewConnection(ctx, cfg.Endpoint)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn, *cfg))
	c.SetWriter(NewWriter(c.Connection()))

	return c
}
