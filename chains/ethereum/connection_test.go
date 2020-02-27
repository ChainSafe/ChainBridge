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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"
const TestPrivateKey = "39a9ea0dce63086c64a80ce045b796bebed2006554e3992d92601515c7b19807"
const TestPrivateKey2 = "5de3b6992e5ad40dc346cfd3b00595f58bd16ea38b43511e7a00a2a60d380225"

var TestAddress = ethcmn.HexToAddress("34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4")
var TestAddress2 = ethcmn.HexToAddress("0a4c3620AF8f3F182e203609f90f7133e018Bf5D")
var TestCentrifugeContractAddress = ethcmn.HexToAddress("70486404e42d17298c57b046Aa162Dc3aCc075f0")
var TestReceiverContractAddress = ethcmn.HexToAddress("705D4Fa884AF2Ae59C7780A0f201109947E2Bf6D")
var TestEmitterContractAddress = ethcmn.HexToAddress("60F9363AaF4993ABA818D5438db5E64bCe6E612b")

const TestTimeout = time.Second * 10

func TestConnect(t *testing.T) {
	cfg := &Config{
		endpoint: TestEndpoint,
		from:     "ethereum",
		keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
	}
	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

func TestSendTx(t *testing.T) {
	conn := newLocalConnection(t, ethcmn.HexToAddress("0x0"))
	defer conn.Close()

	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	TestAddr := keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].(*secp256k1.Keypair).Public().Address()
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
		keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
		from:     "ethereum",
	}

	conn := NewConnection(cfg)
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

func createTestAuth(t *testing.T, conn *Connection) *bind.TransactOpts {
	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	TestAddr := keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].(*secp256k1.Keypair).Public().Address()
	nonce, err := conn.NonceAt(ethcmn.HexToAddress(TestAddr), currBlock.Number())

	if err != nil {
		t.Fatal(err)
	}

	privateKey := conn.kp.Private().(*secp256k1.PrivateKey).Key()
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = big.NewInt(10)

	return auth
}
