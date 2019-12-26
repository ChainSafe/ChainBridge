package ethereum

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"
const TestPrivateKey = "39a9ea0dce63086c64a80ce045b796bebed2006554e3992d92601515c7b19807"
const TestAddress = "34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4" // Account address
const TestCentrifugeContractAddress = "290f41e61374c715C1127974bf08a3993afd0145"
const TestEmitterContractAddress = "1fA38b0EfccA4228EB9e15112D4d98B0CEe3c600"
const TestTimeout = time.Second * 30

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
