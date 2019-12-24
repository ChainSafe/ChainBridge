package ethereum

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var TestEndpoint = "ws://localhost:8545"
var TestPrivateKey = "ae6a8b4518e3970a0501ecf796a51dc0dab9143a66be75e948bf352582db15d5"
var TestAddress = "C5F737aE7EaBB7226f21121E335b0949d8eA6365"
var TestCentrifugeContractAddress = "F60D9c8AC3B9B88483cee749b25117330F927780"

func TestConnect(t *testing.T) {
	ctx := context.Background()
	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: TestEndpoint,
	}
	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

func TestSendTx(t *testing.T) {
	conn := newLocalConnection(t, "")
	defer conn.Close()

	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	nonce, err := conn.NonceAt(common.StringToAddress(TestAddress), currBlock.Number())
	if err != nil {
		t.Fatal(err)
	}

	tx := ethtypes.NewTransaction(
		nonce,
		ethcommon.Address([20]byte{}),
		big.NewInt(0),
		1000000,
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
		Endpoint: TestEndpoint,
	}

	conn := NewConnection(cfg)
	chainCfg := &core.ChainConfig{
		Id:            msg.EthereumId,
		Endpoint:      TestEndpoint,
		Receiver:      "",
		Emitter:       "",
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

	_, err = l.conn.subscribeToEvent(q)
	if err != nil {
		t.Fatal(err)
	}
}
