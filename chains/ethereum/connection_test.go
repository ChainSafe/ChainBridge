// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"testing"

	"github.com/ChainSafe/ChainBridge/blockstore"
	eth "github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
)

func TestConnect(t *testing.T) {
	_ = deployTestContracts(t, aliceTestConfig.id, AliceKp)
	conn := newLocalConnection(t, aliceTestConfig)
	conn.Close()
}

func TestSubscribe(t *testing.T) {
	_ = deployTestContracts(t, aliceTestConfig.id, AliceKp)
	conn := newLocalConnection(t, aliceTestConfig)
	l := NewListener(conn, aliceTestConfig, TestLogger, &blockstore.EmptyStore{})
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
	contracts := deployTestContracts(t, aliceTestConfig.id, AliceKp)
	conn := newLocalConnection(t, aliceTestConfig)
	defer conn.Close()

	// The following section checks if the byteCode exists on the chain at the specificed Addresses
	err := conn.ensureHasBytecode(contracts.BridgeAddress)
	if err != nil {
		t.Fatal(err)
	}

	err = conn.ensureHasBytecode(ethcmn.HexToAddress("0x0"))
	if err == nil {
		t.Fatal("should detect no bytecode")
	}

}
