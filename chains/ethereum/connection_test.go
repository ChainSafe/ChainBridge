// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"strings"
	"testing"

	eth "github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func TestConnect(t *testing.T) {
	_ = deployTestContracts(t, aliceTestConfig.id)
	conn := newLocalConnection(t, aliceTestConfig)
	conn.Close()
}

func TestSendTx(t *testing.T) {
	_ = deployTestContracts(t, aliceTestConfig.id)
	conn := newLocalConnection(t, aliceTestConfig)
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
	_ = deployTestContracts(t, aliceTestConfig.id)
	conn := newLocalConnection(t, aliceTestConfig)
	l := NewListener(conn, aliceTestConfig, TestLogger)
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
	contracts := deployTestContracts(t, aliceTestConfig.id)
	conn := newLocalConnection(t, aliceTestConfig)
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	err := conn.checkBridgeContract(contracts.BridgeAddress)
	if err != nil {
		t.Fatal(err)
	}

}
