package centrifuge

import (
	"context"
	"math/big"
	"testing"

	"ChainBridgeV2/types"

	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var TestPolkadotEndpoint = "wss://poc3-rpc.polkadot.io"

func TestConnect(t *testing.T) {
	ctx := context.Background()
	conn := NewConnection(TestPolkadotEndpoint)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

//func TestSendTx(t *testing.T) {
//	ctx := context.Background()
//	conn := NewConnection(ctx, TestPolkadotEndpoint)
//	err := conn.Connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer conn.Close()
//
//	tx := ethtypes.NewTransaction(
//		1,
//		ethcommon.Address([20]byte{}),
//		big.NewInt(0),
//		1,
//		big.NewInt(1),
//		[]byte{},
//	)
//
//	data, err := tx.MarshalJSON()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// TODO: this fails since we don't since the tx
//	err = conn.SubmitTx(data)
//	if err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestSubscribe(t *testing.T) {
//	conn := NewConnection(TestPolkadotEndpoint)
//	err := conn.Connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer conn.Close()
//
//	q := eth.FilterQuery{}
//
//	// TODO: this fails with "notifications not supported"
//	_, err = conn.Subscribe(types.TransferEvent, q)
//	if err != nil {
//		t.Fatal(err)
//	}
//}
