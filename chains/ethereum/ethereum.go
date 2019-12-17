package ethereum

import (
	"context"
	"fmt"
	"syscall"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/ssh/terminal"
)

func InitializeChain(id msg.ChainId, cfg *core.ChainConfig) (*core.Chain, error) {
	ctx := context.Background()
	c := core.NewChain(id, cfg)

	path := cfg.From
	pswd := getPassword(fmt.Sprintf("Enter password for key %s:", path))
	priv, err := keystore.ReadFromFileAndDecrypt(path, pswd)
	if err != nil {
		return nil, err
	}

	kp, err := keystore.PrivateKeyToKeypair(priv)
	if err != nil {
		return nil, err
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
	c.SetListener(NewListener(conn))
	c.SetWriter(NewWriter(c.Connection()))
	return c, nil
}

// prompt user to enter password for encrypted keystore
func getPassword(msg string) []byte {
	for {
		fmt.Println(msg)
		fmt.Print("> ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("invalid input: %s\n", err)
		} else {
			fmt.Printf("\n")
			return password
		}
	}
}
