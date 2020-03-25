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
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]

const TestTimeout = time.Second * 10

var connectionTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
}

func newLocalConnection(t *testing.T, cfg *Config) *Connection {
	conn := NewConnection(cfg, AliceKp)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func conn_deployContracts(t *testing.T) {
	port := "8545"
	numRelayers := 2
	relayerThreshold := big.NewInt(1)
	pk := hexutil.Encode(AliceKp.Encode())[2:]
	DeployedContracts, err := DeployContracts(pk, port, numRelayers, relayerThreshold, uint8(0))
	if err != nil {
		t.Fatal(err)
	}

	connectionTestConfig.contract = DeployedContracts.BridgeAddress
}

func TestConnect(t *testing.T) {
	conn_deployContracts(t)
	conn := newLocalConnection(t, connectionTestConfig)
	conn.Close()
}

func TestSendTx(t *testing.T) {
	conn_deployContracts(t)
	conn := newLocalConnection(t, connectionTestConfig)
	defer conn.Close()

	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	TestAddr := AliceKp.Address()
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
	conn_deployContracts(t)
	conn := newLocalConnection(t, connectionTestConfig)
	l := NewListener(conn, connectionTestConfig)
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
	conn_deployContracts(t)
	conn := newLocalConnection(t, connectionTestConfig)
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	err := conn.checkBridgeContract(connectionTestConfig.contract)
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
