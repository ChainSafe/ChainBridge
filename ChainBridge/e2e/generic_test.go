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
	log "github.com/ChainSafe/log15"
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
			eth.WaitForProposalActive(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethA.Client, hash, ctx.ethA.TestContracts.AssetStoreSub)
		})
		if !ok {
			return
		}
	}
}

func testEthereumHashToGenericHandler(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, EthBChainId) + 1
	//ethtest.AssertGenericResourceAddress(t, ctx.ethA.Client, ctx.ethA.BaseContracts.GenericHandlerAddress, ctx.EthGenericResourceId, ctx.ethA.TestContracts.AssetStoreEth)
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			hash := sub.HashInt(i)
			log.Info("Submitting transaction", "number", i, "hash", hash.Hex(), "resourceId", ctx.EthGenericResourceId.Hex(), "from", ctx.ethA.Client.Opts.From, "handler", ctx.ethA.BaseContracts.GenericHandlerAddress)
			eth.CreateGenericDeposit(t, ctx.ethA.Client, EthBChainId, hash[:], ctx.ethB.BaseContracts, ctx.EthGenericResourceId)

			// Wait for event
			eth.WaitForProposalActive(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethB.Client, hash, ctx.ethB.TestContracts.AssetStoreEth)
		})
		if !ok {
			return
		}
	}
}
