// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"testing"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
)

func testSubstrateHashToGenericHandler(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	nonce := subtest.GetDepositNonce(t, ctx.subClient, EthAChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			hash := sub.HashInt(i)
			subtest.InitiateHashTransfer(t, ctx.subClient, hash, EthAChainId)

			// Wait for event
			eth.WaitForDepositCreatedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForDepositExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethA.Client, hash, ctx.ethA.BaseContracts.CentrifugeHandlerAddress)
		})
		if !ok {
			return
		}
	}
}
