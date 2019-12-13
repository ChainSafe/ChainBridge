package ethereum

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/types"

	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

var TestEthereumEndpoint = "https://rinkeby.infura.io/v3/b0a01296903f4812b5ec2cf26cbded48"

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

	signer := ethtypes.MakeSigner(ethparams.RinkebyChainConfig, ethparams.RinkebyChainConfig.IstanbulBlock)

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
	t.Skip()
	ctx := context.Background()
	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: TestEthereumEndpoint,
	}

	conn := NewConnection(cfg)
	l := NewListener(conn)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	q := eth.FilterQuery{}

	// TODO: this fails with "notifications not supported"
	_, err = l.subscribe(types.TransferEvent, q)
	if err != nil {
		t.Fatal(err)
	}
}
