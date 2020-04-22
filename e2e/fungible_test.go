// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"math/big"
	"testing"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func testErc20ToSubstrate(t *testing.T, ctx *testContext) {
	numberOfTxs := 5

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			log.Info("Submitting transaction", "number", i, "from", ctx.ethA.Opts.From, "handler", ctx.ethA.BaseContracts.ERC20HandlerAddress.String(), "amount", amount.String())
			// Initiate transfer
			eth.CreateErc20Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, SubChainId, sub.BobKp.AsKeyringPair().PublicKey, amount, ctx.ethA.BaseContracts, ctx.EthSubResourceId)

			// Check for success event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.NewU64(uint64(i)), types.U8(EthAChainId))
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}

func testSubstrateToErc20(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	expectedBalance := big.NewInt(0)
	recipient := eth.CharlieAddr
	ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedBalance, ctx.ethA.TestContracts.Erc20Sub, recipient)

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Get latest eth block
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethA.Client)

			// Execute transfer
			amount := types.NewU32(uint32(i * 5))
			subtest.InitiateNativeTransfer(t, ctx.subClient, amount, recipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify balance change
			expectedBalance.Add(expectedBalance, big.NewInt(int64(amount)))
			ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedBalance, ctx.ethA.TestContracts.Erc20Sub, recipient)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}

func testErc20ToErc20(t *testing.T, ctx *testContext) {
	recipient := eth.CharlieAddr
	expectedBalance := big.NewInt(0)

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethB.Client)

			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			log.Info("Submitting transaction", "number", i, "recipient", recipient, "resourcId", ctx.EthEthResourceId.Hex(), "amount", amount.String(), "from", ctx.ethA.Opts.From, "handler", ctx.ethA.BaseContracts.ERC20HandlerAddress)
			// Initiate transfer
			eth.CreateErc20Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, EthBChainId, recipient.Bytes(), amount, ctx.ethA.BaseContracts, ctx.EthEthResourceId)

			eth.WaitForEthereumEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify balance change
			expectedBalance.Add(expectedBalance, amount)
			ethtest.Erc20AssertBalance(t, ctx.ethB.Client, expectedBalance, ctx.ethB.TestContracts.Erc20Eth, recipient)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}

func testErc20SubstrateRoundTrip(t *testing.T, ctx *testContext) {
	// Transfer params
	numberOfTxs := 5
	subRecipient := sub.DaveKp.AsKeyringPair().PublicKey
	ethRecipient := eth.AliceAddr

	expectedSubBalance := subtest.BalanceOf(t, ctx.subClient, subRecipient)
	initialEthBalance := ethtest.Erc20BalanceOf(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	expectedEthBalance := ethtest.Erc20BalanceOf(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Eth to Substrate Transfer %d", i), func(t *testing.T) {
			// Initiate transfer
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			log.Info("Submitting transaction", "number", i, "amount", amount.String())
			eth.CreateErc20Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, SubChainId, subRecipient, amount, ctx.ethA.BaseContracts, ctx.EthSubResourceId)

			// Check for success event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.NewU64(uint64(i+5)), types.U8(EthAChainId))
			// Verify balance
			expectedSubBalance.Add(expectedSubBalance, amount)
			subtest.AssertBalanceOf(t, ctx.subClient, subRecipient, expectedSubBalance)
			expectedEthBalance.Sub(expectedEthBalance, amount)
			ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedEthBalance, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
			log.Info("Asserted balance", "owner", ethRecipient, "balance", expectedEthBalance.String())

		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}

	// Repeat the process in the opposite direction
	expectedEthBalance = ethtest.Erc20BalanceOf(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	log.Info("Asserted resulting balance", "owner", ethRecipient, "balance", expectedEthBalance.String())
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Substrate to Eth Transfer %d", i), func(t *testing.T) {
			// Get latest eth block
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethA.Client)

			// Execute transfer
			amount := types.NewU32(uint32(i * 5))
			log.Info("Submitting transaction", "number", i, "amount", amount)
			subtest.InitiateNativeTransfer(t, ctx.subClient, amount, ethRecipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify balance change
			expectedEthBalance.Add(expectedEthBalance, big.NewInt(int64(amount)))
			ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedEthBalance, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
			log.Info("Asserted balance", "owner", ethRecipient, "balance", expectedEthBalance.String())
			// TODO: Presently unable to take gas costs into consideration.
			//expectedSubBalance.Add(expectedSubBalance, big.NewInt(int64(amount)))
			//subtest.AssertBalanceOf(t, ctx.subClient, subRecipient, expectedSubBalance)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}

	ethtest.Erc20AssertBalance(t, ctx.ethA.Client, initialEthBalance, ctx.ethA.TestContracts.Erc20Sub, ctx.ethA.Opts.From)

}
