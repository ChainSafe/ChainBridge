package ethereum

import (
	"context"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

func InitializeChain(id msg.ChainId, cfg *core.ChainConfig, kp crypto.Keypair) *core.Chain {
	ctx := context.Background()
	c := core.NewChain(id, cfg)

	conn := NewConnection(ctx, cfg.Endpoint, kp)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}
