// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"testing"

	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	ethcmn "github.com/ethereum/go-ethereum/common"
)

func TestConnect(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)

	_ = deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
	conn := newLocalConnection(t, aliceTestConfig)
	conn.Close()
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
