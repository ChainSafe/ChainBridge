// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subutils "github.com/ChainSafe/ChainBridge/shared/substrate"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/common"
)

// Random recipient
var SteveSubKp = signature.KeyringPair{
	URI:       "//Steve",
	Address:   "5F7UdYV6noJeqbMfiiJKfHDTGYnBiqhMgY2EBi3nFAqcL33w",
	PublicKey: []byte{0x86, 0xd1, 0xe9, 0xb3, 0x79, 0x37, 0x23, 0x39, 0x02, 0xca, 0xe0, 0x62, 0xf5, 0xd0, 0x1c, 0xae, 0x46, 0x38, 0x58, 0x42, 0xe7, 0xec, 0x9d, 0x1c, 0xeb, 0x6b, 0x1b, 0x10, 0x67, 0x0e, 0x30, 0x27},
}
var SteveEthAddr = common.HexToAddress("0x880fd09782C3183c489595111637bb3bc906f215")

// Final recipients
var subRecipient = SteveSubKp.PublicKey
var ethRecipient = SteveEthAddr

// Tx's per configuration
var numberOfTxs = 5

// Value per transaction
var amountPerTest = types.NewU128(*big.NewInt(500))

// Expected overall balance change for recipients
var balanceDelta = big.NewInt(0).Mul(amountPerTest.Int, big.NewInt(int64(numberOfTxs)))

// Substrate block time
var blockTime = time.Second * 5

// Delay between tx's to allow more interleaving
var txInterval = time.Millisecond * 200

func testThreeChainsParallel(t *testing.T, ctx *testContext) {
	// Creat some additional clients for thread safety
	subClient := ctx.subClient
	ethAClientA := ethtest.NewClient(t, eth.EthAEndpoint, eth.AliceKp)
	ethAClientB := ethtest.NewClient(t, eth.EthAEndpoint, eth.EveKp)
	ethBClientA := ethtest.NewClient(t, eth.EthBEndpoint, eth.AliceKp)

	// Mint tokens to eve and dave, approve handlers
	ethtest.Erc20Mint(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Sub, eth.AliceKp.CommonAddress(), balanceDelta)
	ethtest.Erc20Mint(t, ctx.ethA.Client, ctx.ethA.TestContracts.Erc20Eth, eth.EveKp.CommonAddress(), balanceDelta)
	ethtest.Erc20Mint(t, ctx.ethB.Client, ctx.ethB.TestContracts.Erc20Eth, eth.AliceKp.CommonAddress(), balanceDelta)
	ethtest.Erc20Approve(t, ethAClientA, ctx.ethA.TestContracts.Erc20Sub, ctx.ethA.BaseContracts.ERC20HandlerAddress, balanceDelta)
	ethtest.Erc20Approve(t, ethAClientB, ctx.ethA.TestContracts.Erc20Eth, ctx.ethA.BaseContracts.ERC20HandlerAddress, balanceDelta)
	ethtest.Erc20Approve(t, ethBClientA, ctx.ethB.TestContracts.Erc20Eth, ctx.ethB.BaseContracts.ERC20HandlerAddress, balanceDelta)

	// Get current balances
	subRecipientBalance := subtest.BalanceOf(t, subClient, subRecipient)
	ethASubRecipientBalance := ethtest.Erc20BalanceOf(t, ethAClientA, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	ethAEthRecipientBalance := ethtest.Erc20BalanceOf(t, ethAClientA, ctx.ethA.TestContracts.Erc20Eth, ethRecipient)
	ethBRecipientBalance := ethtest.Erc20BalanceOf(t, ethBClientA, ctx.ethB.TestContracts.Erc20Eth, ethRecipient)

	// Execute several sub-tests that call t.Paralell
	t.Run("Parallel tx submission", func(t *testing.T) {

		// Substrate -> EthA
		t.Run("Submit Sub to Eth", func(t *testing.T) {
			// TODO: Need to run this in sequence, doesn't return if parallel
			//t.Parallel()
			submitSubToEth(t, subClient, EthAChainId, ethRecipient.Bytes(), amountPerTest, ctx.EthSubErc20ResourceId)

		})
		// EthA -> Substrate
		t.Run("Submit Eth to Sub", func(t *testing.T) {
			t.Parallel()
			submitEthDeposit("Eth to Sub", t, ctx.ethA, ethAClientA, SubChainId, subRecipient, amountPerTest.Int, ctx.EthSubErc20ResourceId)
		})
		// EthA -> EthB
		t.Run("Submit EthA to EthB", func(t *testing.T) {
			t.Parallel()
			submitEthDeposit("EthA to EthB", t, ctx.ethA, ethAClientB, EthBChainId, ethRecipient.Bytes(), amountPerTest.Int, ctx.EthEthErc20ResourceId)
		})
		// EthB -> EthA
		t.Run("Submit EthB to EthA", func(t *testing.T) {
			t.Parallel()
			submitEthDeposit("EthB to EthA", t, ctx.ethB, ethBClientA, EthAChainId, ethRecipient.Bytes(), amountPerTest.Int, ctx.EthEthErc20ResourceId)
		})
	})

	// Must wait long enough for processing to complete
	time.Sleep(blockTime * 10)

	// Calculate and verify expected results
	t.Run("Assert Sub balance", func(t *testing.T) {
		subtest.AssertBalanceOf(t, subClient, subRecipient, subRecipientBalance.Add(subRecipientBalance, balanceDelta))
	})
	t.Run("Assert EthA Sub balance", func(t *testing.T) {
		ethtest.Erc20AssertBalance(t, ethAClientA, ethASubRecipientBalance.Add(ethASubRecipientBalance, balanceDelta), ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	})
	t.Run("Assert EthB balance", func(t *testing.T) {
		ethtest.Erc20AssertBalance(t, ethBClientA, ethBRecipientBalance.Add(ethBRecipientBalance, balanceDelta), ctx.ethB.TestContracts.Erc20Eth, ethRecipient)
	})
	t.Run("Assert EthA Eth balance", func(t *testing.T) {
		ethtest.Erc20AssertBalance(t, ethAClientA, ethAEthRecipientBalance.Add(ethAEthRecipientBalance, balanceDelta), ctx.ethA.TestContracts.Erc20Eth, ethRecipient)
	})

}

func submitEthDeposit(name string, t *testing.T, ethCtx *eth.TestContext, client *ethutils.Client, destId msg.ChainId, recipient []byte, amount *big.Int, rId msg.ResourceId) {
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		t.Run(fmt.Sprintf("%s Transfer %d", name, i), func(t *testing.T) {
			// Initiate transfer
			log.Info("Submitting transaction", "number", i, "recipient", recipient, "amount", amount, "rId", rId.Hex())
			ethtest.LockNonceAndUpdate(t, client)
			eth.CreateErc20Deposit(t, client, destId, recipient, amount, ethCtx.BaseContracts, rId)
			client.UnlockNonce()
			time.Sleep(txInterval)
		})
	}
}

func submitSubToEth(t *testing.T, client *subutils.Client, destId msg.ChainId, recipient []byte, amount types.U128, rId msg.ResourceId) {
	var calls []types.Call
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		t.Run(fmt.Sprintf("Substrate to Eth Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			log.Info("Creating transaction", "number", i, "recipient", recipient, "amount", amount, "rId", rId.Hex())
			call := subtest.NewNativeTransferCall(t, client, amount, recipient, destId)
			calls = append(calls, call)
		})
	}
	log.Info("Submitting transactions", "numberOfTxs", numberOfTxs, "amount/tx", amountPerTest)

	err := subutils.BatchSubmit(client, calls)
	if err != nil {
		t.Fatal(err)
	}
}
