package ethereum

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

var TestEthereumEndpoint = "wss://goerli.infura.io/ws/v3/b0a01296903f4812b5ec2cf26cbded48"

func TestConnect(t *testing.T) {
	ctx := context.Background()
	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: TestEthereumEndpoint,
	}
	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

func TestSendTx(t *testing.T) {
	ctx := context.Background()

	kp, err := secp256k1.GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	signer := ethtypes.MakeSigner(ethparams.GoerliChainConfig, ethparams.GoerliChainConfig.IstanbulBlock)

	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: TestEthereumEndpoint,
		Keypair:  kp,
		Signer:   signer,
	}

	conn := NewConnection(cfg)
	err = conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	tx := ethtypes.NewTransaction(
		1,
		ethcommon.Address([20]byte{}),
		big.NewInt(0),
		1,
		big.NewInt(1),
		[]byte{},
	)

	data, err := tx.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	err = conn.SubmitTx(data)
	if err != nil && strings.Compare(err.Error(), "insufficient funds for gas * price + value") != 0 {
		t.Fatal(err)
	}
}

func TestSubscribe(t *testing.T) {
	ctx := context.Background()
	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: TestEthereumEndpoint,
	}

	conn := NewConnection(cfg)
	chainCfg := core.ChainConfig{
		Id:            msg.EthereumId,
		Endpoint:      TestEthereumEndpoint,
		Home:          "",
		Away:          "",
		From:          "",
		Subscriptions: nil,
	}
	l := NewListener(conn, chainCfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	q := eth.FilterQuery{}

	_, err = l.conn.subscribeToEvent(q, "func(uint256)")
	if err != nil {
		t.Fatal(err)
	}
}
