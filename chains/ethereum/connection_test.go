// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"bytes"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	"github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	"github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
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
var TestCentrifugeContractAddress = ethcmn.HexToAddress("0xcB76d991cFCd621b477d705be7DdF5EA69D39C00")
var TestReceiverContractAddress = ethcmn.HexToAddress("0x5842B333910Fe0BfA05F5Ea9F1602a40d1AF3584")
var TestEmitterContractAddress = ethcmn.HexToAddress("0x3c747684333605408F9A4907DA043ee4c1A72D9c")

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

func panicOnTimeout(d time.Duration) {
	<-time.After(d)
	panic("Test timed out")
}

// TestContractCode is used to make sure the contracts are deployed correctly.
// This is probably the least intrusive way to check if the contracts exists
func TestContractCode(t *testing.T) {

	// We set a custom timeout for this test
	// This is because of a ganache bug
	// To reproduce the bug, attempt to create a contract Object (eg Emitter) with a valid Address of a different contract type (eg RecieverAddress)
	// It will cause the test to hang forever, this code is to prevent the hanging in that case.
	go panicOnTimeout(TestTimeout)

	conn := NewConnection(testConfig)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	byteCode, err := conn.getByteCode(TestEmitterContractAddress)
	if err != nil {
		t.Fatal(err)
	}
	if len(byteCode) == 0 {
		t.Fatal("Emitter Contract doesn't exist")
	}

	byteCode, err = conn.getByteCode(TestReceiverContractAddress)
	if err != nil {
		t.Fatal(err)
	}
	if len(byteCode) == 0 {
		t.Fatal("Receiver Contract doesn't exist")
	}

	byteCode, err = conn.getByteCode(TestCentrifugeContractAddress)
	if err != nil {
		t.Fatal(err)
	}
	if len(byteCode) == 0 {
		t.Fatal("BridgeAsset Contract doesn't exist")
	}

	// This section attempts to make the correct Contract type, trying to ensure that the contract at the addresses are the correct types
	emit, err := Emitter.NewEmitter(TestEmitterContractAddress, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	rec, err := Receiver.NewReceiver(TestReceiverContractAddress, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	_, err = BridgeAsset.NewBridgeAsset(TestCentrifugeContractAddress, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	// This following section attempts to make calls on Emitter and Receiver.
	// There is no simple call that can be made on BridgeAsset, so unfortunately it is not called.

	// This is a call on the emitter, it checks if the owner is the same as the address.
	owner, err := emit.Owner(nil)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(TestEmitterContractAddress.Bytes(), owner.Bytes()) {
		t.Fatal("Incorrect Contract")
	}

	// We don't know what value this should be checking agaisnt, so we just make sure this call doesn't error out.
	_, err = rec.TotalValidators(nil)
	if err != nil {
		t.Fatal(err)
	}
}
