package ethereum

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
	"context"
)

func InitializeChain(id msg.ChainId, cfg *core.ChainConfig) *core.Chain {
	ctx := context.Background()
	c := core.NewChain(id, cfg)

	conn := NewConnection(ctx, cfg.Endpoint)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}
