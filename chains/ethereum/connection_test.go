// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	eth "github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"

var Alice = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].(*secp256k1.Keypair)

var TestAddress = ethcmn.HexToAddress("34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4")
var TestAddress2 = ethcmn.HexToAddress("0a4c3620AF8f3F182e203609f90f7133e018Bf5D")
var TestBridgeContractAddress = ethcmn.HexToAddress("0xcB76d991cFCd621b477d705be7DdF5EA69D39C00")
var TestReceiverContractAddress = ethcmn.HexToAddress("0x5842B333910Fe0BfA05F5Ea9F1602a40d1AF3584")
var TestEmitterContractAddress = ethcmn.HexToAddress("0x3c747684333605408F9A4907DA043ee4c1A72D9c")

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
	err = conn.CheckBridgeContract(TestBridgeContractAddress)
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
