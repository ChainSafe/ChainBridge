// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"strings"
	"testing"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const TestEndpoint = "ws://localhost:8545"

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]

var defaultDeployOpts = DeployOpts{
	pk:               hexutil.Encode(AliceKp.Encode())[2:],
	url:              "http://localhost:8545",
	numRelayers:      2,
	relayerThreshold: big.NewInt(1),
	minCount:         uint8(0),
}

type DeployOpts struct {
	pk               string
	url              string
	numRelayers      int
	relayerThreshold *big.Int
	minCount         uint8
}

func setOpts(opts DeployOpts) DeployOpts {
	cfg := defaultDeployOpts
	if opts.pk != "" {
		cfg.pk = opts.pk
	}
	if opts.url != "" {
		cfg.url = opts.url
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

func deployContracts(t *testing.T, customOpts DeployOpts) (*Config, *DeployedContracts) {
	opts := setOpts(customOpts)
	deployedContracts, err := DeployContracts(opts.pk, opts.url, opts.numRelayers, opts.relayerThreshold, opts.minCount)
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
		},
		deployedContracts
}

func createBridgeInstance(t *testing.T, connection *Connection, address common.Address) BridgeContract {
	bridgeInstance, err := bridge.NewBridge(address, connection.conn)
	if err != nil {
		t.Fatal(err)
	}

	raw := &bridge.BridgeRaw{
		Contract: bridgeInstance,
	}

	bridgeContract := BridgeContract{
		BridgeRaw:        raw,
		BridgeCaller:     &bridgeInstance.BridgeCaller,
		BridgeTransactor: &bridgeInstance.BridgeTransactor,
	}
	return bridgeContract
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
	cfg, _ := deployContracts(t, defaultDeployOpts)
	conn := newLocalConnection(t, cfg)
	conn.Close()
}

func TestSendTx(t *testing.T) {
	cfg, _ := deployContracts(t, defaultDeployOpts)
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
	cfg, _ := deployContracts(t, defaultDeployOpts)
	conn := newLocalConnection(t, cfg)
	l := NewListener(conn, cfg)
	defer conn.Close()

	q := eth.FilterQuery{}

	_, err := l.conn.subscribeToEvent(q)
	if err != nil {
		t.Fatal(err)
	}
}

// TestContractCode is used to make sure the contracts are deployed correctly.
// This is probably the least intrusive way to check if the contracts exists
func TestContractCode(t *testing.T) {
	cfg, _ := deployContracts(t, defaultDeployOpts)
	conn := newLocalConnection(t, cfg)
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	err := conn.checkBridgeContract(cfg.contract)
	if err != nil {
		t.Fatal(err)
	}

}
