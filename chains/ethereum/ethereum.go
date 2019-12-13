package ethereum

import (
	"context"

	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

func InitializeChain(id msg.ChainId, cfg *core.ChainConfig) *core.Chain {
	ctx := context.Background()
	c := core.NewChain(id, cfg)

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
	c.SetListener(NewListener(conn))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}
