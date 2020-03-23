// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	eth "github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"

var Alice = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]

var TestBridgeContractAddress = ethcmn.HexToAddress("0x3167776db165D8eA0f51790CA2bbf44Db5105ADF")

const TestTimeout = time.Second * 10

func newLocalConnection(t *testing.T) *Connection {

	cfg := &Config{
		endpoint: TestEndpoint,
		contract: TestBridgeContractAddress,
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg, Alice)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func TestConnect(t *testing.T) {
	conn := newLocalConnection(t)
	conn.Close()
}

func TestSendTx(t *testing.T) {
	conn := newLocalConnection(t)
	defer conn.Close()

	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	TestAddr := Alice.Address()
	nonce, err := conn.NonceAt(ethcmn.HexToAddress(TestAddr), currBlock.Number())
	if err != nil {
		t.Fatal(err)
	}

	tx := ethtypes.NewTransaction(
		nonce,
		ethcmn.Address([20]byte{}),
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
	cfg := &Config{
		id:       msg.EthereumId,
		endpoint: TestEndpoint,
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg, Alice)
	l := NewListener(conn, cfg)
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

// TestContractCode is used to make sure the contracts are deployed correctly.
// This is probably the least intrusive way to check if the contracts exists
func TestContractCode(t *testing.T) {
	cfg := &Config{
		id:       msg.EthereumId,
		endpoint: TestEndpoint,
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg, Alice)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	err = conn.checkBridgeContract(TestBridgeContractAddress)
	if err != nil {
		t.Fatal(err)
	}

}

// Unused, may be useful in the future
//func createTestAuth(t *testing.T, conn *Connection) *bind.TransactOpts {
//	currBlock, err := conn.LatestBlock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	TestAddr := keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].(*secp256k1.Keypair).Public().Address()
//	nonce, err := conn.NonceAt(ethcmn.HexToAddress(TestAddr), currBlock.Number())
//
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	privateKey := conn.kp.Private().(*secp256k1.PrivateKey).Key()
//	auth := bind.NewKeyedTransactor(privateKey)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)     // in wei
//	auth.GasLimit = uint64(300000) // in units
//	auth.GasPrice = big.NewInt(10)
//
//	return auth
//}
