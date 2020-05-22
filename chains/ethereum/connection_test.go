// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"testing"

	"github.com/ChainSafe/ChainBridge/blockstore"
	eth "github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
)

func TestConnect(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)

	_ = deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
	conn := newLocalConnection(t, aliceTestConfig)
	conn.Close()
}

func TestSubscribe(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)

	_ = deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
	conn := newLocalConnection(t, aliceTestConfig)
	l := NewListener(conn, aliceTestConfig, TestLogger, &blockstore.EmptyStore{}, make(chan int), make(chan error))
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
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)

	contracts := deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
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
