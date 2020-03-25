// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	eth "github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]

type deployOpts struct {
	pk               string
	port             string
	numRelayers      int
	relayerThreshold *big.Int
	minCount         uint8
}

var defaultDeployOpts = deployOpts{
	pk:               hexutil.Encode(AliceKp.Encode())[2:],
	port:             "8545",
	numRelayers:      2,
	relayerThreshold: big.NewInt(1),
	minCount:         uint8(0),
}

var emptyDeployOpts = deployOpts{
	pk:               "",
	port:             "",
	numRelayers:      0,
	relayerThreshold: nil,
	minCount:         0,
}

func setOpts(opts deployOpts) deployOpts {
	cfg := defaultDeployOpts
	if opts.pk != "" {
		cfg.pk = opts.pk
	}
	if opts.port != "" {
		cfg.port = opts.port
	}
	if opts.numRelayers != 0 {
		cfg.numRelayers = opts.numRelayers
	}
	if opts.relayerThreshold != nil {
		cfg.relayerThreshold = opts.relayerThreshold
	}
	if opts.minCount != 0 {
		cfg.minCount = opts.minCount
	}
	return cfg
}

func testDeployContracts(t *testing.T, customOpts deployOpts) *Config {
	opts := setOpts(customOpts)
	deployedContracts, err := DeployContracts(opts.pk, opts.port, opts.numRelayers, opts.relayerThreshold, opts.minCount)
	if err != nil {
		t.Fatal(err)
	}
	return &Config{
		id:       msg.EthereumId,
		endpoint: TestEndpoint,
		from:     keystore.AliceKey,
		gasLimit: big.NewInt(6721975),
		gasPrice: big.NewInt(20000000000),
		contract: deployedContracts.BridgeAddress,
	}
}

func newLocalConnection(t *testing.T, cfg *Config) *Connection {
	kp := keystore.TestKeyRing.EthereumKeys[cfg.from]
	conn := NewConnection(cfg, kp)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func TestConnect(t *testing.T) {
	cfg := testDeployContracts(t, emptyDeployOpts)
	conn := newLocalConnection(t, cfg)
	conn.Close()
}

func TestSendTx(t *testing.T) {
	cfg := testDeployContracts(t, emptyDeployOpts)
	conn := newLocalConnection(t, cfg)
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
	cfg := testDeployContracts(t, emptyDeployOpts)
	conn := newLocalConnection(t, cfg)
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
	cfg := testDeployContracts(t, emptyDeployOpts)
	conn := newLocalConnection(t, cfg)
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	err := conn.checkBridgeContract(cfg.contract)
	if err != nil {
		t.Fatal(err)
	}

}
