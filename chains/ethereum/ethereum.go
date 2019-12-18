package ethereum

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

func InitializeChain(cfg *core.ChainConfig) (*core.Chain, error) {
	ctx := context.Background()
	c := core.NewChain(cfg)

	var kp crypto.Keypair
	path := cfg.From
	fmt.Println(path)
	if path != "" {
		pswd := []byte(cfg.Password)
		if bytes.Equal(pswd, []byte{}) {
			pswd = common.GetPassword(fmt.Sprintf("Enter password for key %s:", path))
		}

		priv, err := keystore.ReadFromFileAndDecrypt(path, pswd)
		if err != nil {
			return nil, err
		}

		kp, err = keystore.PrivateKeyToKeypair(priv)
		if err != nil {
			return nil, err
		}
	}

	// TODO: add network to use to config
	signer := ethtypes.MakeSigner(ethparams.MainnetChainConfig, ethparams.MainnetChainConfig.IstanbulBlock)

	conncfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: cfg.Endpoint,
		Signer:   signer,
		Keypair:  kp,
	}

	conn := NewConnection(conncfg)
	c.SetConnection(conn)
	c.SetListener(NewListener(conn, *cfg))
	c.SetWriter(NewWriter(c.Connection()))
	return c, nil
}
