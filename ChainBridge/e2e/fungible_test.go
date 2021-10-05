// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"math/big"
	"testing"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func testErc20ToSubstrate(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, SubChainId) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			log.Info("Submitting transaction", "number", i, "from", ctx.ethA.Client.Opts.From, "handler", ctx.ethA.BaseContracts.ERC20HandlerAddress.String(), "amount", amount.String())

			err := utils.UpdateNonce(ctx.ethA.Client)
			if err != nil {
				t.Fatal(err)
			}
			// Initiate transfer
			eth.CreateErc20Deposit(t, ctx.ethA.Client, SubChainId, sub.BobKp.AsKeyringPair().PublicKey, amount, ctx.ethA.BaseContracts, ctx.EthSubErc20ResourceId)

			// Check for success event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.NewU64(nonce), types.U8(EthAChainId))
			nonce++
		})
		if !ok {
			return
		}
	}
}

func testSubstrateToErc20(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	expectedBalance := big.NewInt(0)
	recipient := eth.CharlieKp.CommonAddress()
	ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedBalance, ctx.ethA.TestContracts.Erc20Sub, recipient)
	nonce := subtest.GetDepositNonce(t, ctx.subClient, EthAChainId) + 1

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			amount := types.NewU128(*big.NewInt(int64(i * 5)))
			subtest.InitiateNativeTransfer(t, ctx.subClient, amount, recipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForProposalActive(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify balance change
			expectedBalance.Add(expectedBalance, amount.Int)
			ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedBalance, ctx.ethA.TestContracts.Erc20Sub, recipient)
		})
		if !ok {
			return
		}
	}
}

func testErc20ToErc20(t *testing.T, ctx *testContext) {
	recipient := eth.CharlieKp.CommonAddress()
	expectedBalance := big.NewInt(0)
	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, EthBChainId) + 1

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			log.Info("Submitting transaction", "number", i, "recipient", recipient, "resourcId", ctx.EthEthErc20ResourceId.Hex(), "amount", amount.String(), "from", ctx.ethA.Client.Opts.From, "handler", ctx.ethA.BaseContracts.ERC20HandlerAddress)

			err := utils.UpdateNonce(ctx.ethA.Client)
			if err != nil {
				t.Fatal(err)
			}

			eth.CreateErc20Deposit(t, ctx.ethA.Client, EthBChainId, recipient.Bytes(), amount, ctx.ethA.BaseContracts, ctx.EthEthErc20ResourceId)

			eth.WaitForProposalActive(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify balance change
			expectedBalance.Add(expectedBalance, amount)
			ethtest.Erc20AssertBalance(t, ctx.ethB.Client, expectedBalance, ctx.ethB.TestContracts.Erc20Eth, recipient)
		})
		if !ok {
			return
		}
	}
}

func testErc20SubstrateRoundTrip(t *testing.T, ctx *testContext) {
	// Transfer params
	numberOfTxs := 5
	subRecipient := sub.AliceKp.AsKeyringPair().PublicKey
	ethRecipient := eth.AliceKp.CommonAddress()

	expectedSubBalance := subtest.BalanceOf(t, ctx.subClient, subRecipient)

	initialEthBalance := ethtest.Erc20BalanceOf(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	expectedEthBalance := ethtest.Erc20BalanceOf(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, SubChainId) + 1

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Eth to Substrate Transfer %d", i), func(t *testing.T) {
			// Initiate transfer
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			log.Info("Submitting transaction", "number", i, "amount", amount.String())

			err := utils.UpdateNonce(ctx.ethA.Client)
			if err != nil {
				t.Fatal(err)
			}

			eth.CreateErc20Deposit(t, ctx.ethA.Client, SubChainId, subRecipient, amount, ctx.ethA.BaseContracts, ctx.EthSubErc20ResourceId)

			// Check for success event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.U64(nonce), types.U8(EthAChainId))
			nonce++
			// Verify balance
			expectedSubBalance.Add(expectedSubBalance, amount)
			subtest.AssertBalanceOf(t, ctx.subClient, subRecipient, expectedSubBalance)

			expectedEthBalance.Sub(expectedEthBalance, amount)
			ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedEthBalance, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
			log.Info("Asserted balance", "owner", ethRecipient, "balance", expectedEthBalance.String())

		})
		if !ok {
			return
		}
	}

	// Repeat the process in the opposite direction
	expectedEthBalance = ethtest.Erc20BalanceOf(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	expectedSubBalance = subtest.BalanceOf(t, ctx.subClient, subRecipient)
	feePerTx := big.NewInt(125000143)
	nonce = subtest.GetDepositNonce(t, ctx.subClient, EthAChainId) + 1

	log.Info("Asserted resulting balance", "owner", ethRecipient, "balance", expectedEthBalance.String())
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Substrate to Eth Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			amount := types.NewU128(*big.NewInt(int64(i * 5)))
			log.Info("Submitting transaction", "number", i, "amount", amount)
			subtest.InitiateNativeTransfer(t, ctx.subClient, amount, ethRecipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForProposalActive(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			eth.WaitForProposalExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, nonce)
			nonce++

			// Verify balance change
			expectedEthBalance.Add(expectedEthBalance, amount.Int)
			ethtest.Erc20AssertBalance(t, ctx.ethA.Client, expectedEthBalance, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
			log.Info("Asserted balance", "owner", ethRecipient, "balance", expectedEthBalance.String())

			expectedSubBalance.Sub(expectedSubBalance, amount.Int)
			expectedSubBalance.Sub(expectedSubBalance, feePerTx)
			subtest.AssertBalanceOf(t, ctx.subClient, subRecipient, expectedSubBalance)
		})
		if !ok {
			return
		}
	}

	ethtest.Erc20AssertBalance(t, ctx.ethA.Client, initialEthBalance, ctx.ethA.TestContracts.Erc20Sub, ctx.ethA.Client.Opts.From)
}
