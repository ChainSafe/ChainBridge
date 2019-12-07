package ethereum

import (
	"context"
	"math/big"
	"testing"

	"ChainBridgeV2/types"

	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var TestEthereumEndpoint = "https://rinkeby.infura.io/v3/b0a01296903f4812b5ec2cf26cbded48"

func TestConnect(t *testing.T) {
	ctx := context.Background()
	conn := NewConnection(ctx, TestEthereumEndpoint)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

func TestSendTx(t *testing.T) {
	ctx := context.Background()
	conn := NewConnection(ctx, TestEthereumEndpoint)
	err := conn.Connect()
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

	// TODO: this fails since we don't sign the tx
	err = conn.SubmitTx(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscribe(t *testing.T) {
	ctx := context.Background()
	conn := NewConnection(ctx, TestEthereumEndpoint)
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
