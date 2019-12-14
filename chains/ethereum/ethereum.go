package ethereum

import (
	"context"

	"github.com/ChainSafe/ChainBridgeV2/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

func InitializeChain(cfg *core.ChainConfig) *core.Chain {
	ctx := context.Background()
	c := core.NewChain(cfg)

	// TODO: add network to use to config
	signer := ethtypes.MakeSigner(ethparams.MainnetChainConfig, ethparams.MainnetChainConfig.IstanbulBlock)

	conncfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: cfg.Endpoint,
		Signer:   signer,
		// TODO: keypair
	}

	conn := NewConnection(conncfg)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn, *cfg))
	c.SetWriter(NewWriter(c.Connection()))

	return c
}
