package e2e

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subutils "github.com/ChainSafe/ChainBridge/shared/substrate"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var subRecipient = sub.DaveKp.AsKeyringPair().PublicKey
var ethRecipient = eth.DaveKp.CommonAddress()
var numberOfTxs = 5
var amountPerTest = types.NewU128(*big.NewInt(100))
var balanceDelta = big.NewInt(0).Mul(amountPerTest.Int, big.NewInt(int64(numberOfTxs)))
var blockTime = time.Second * 5

func testThreeChainsParallel(t *testing.T, ctx *testContext) {
	// Creat some additional clients for thread safety
	subClient := ctx.subClient
	ethAClientA := ethtest.NewClient(t, eth.EthAEndpoint, eth.AliceKp)
	ethAClientB := ethtest.NewClient(t, eth.EthAEndpoint, eth.AliceKp)
	ethBClientA := ethtest.NewClient(t, eth.EthBEndpoint, eth.AliceKp)

	// Get current balances
	subRecipientBalance := subtest.BalanceOf(t, subClient, subRecipient)
	ethASubRecipientBalance := ethtest.Erc20BalanceOf(t, ethAClientA.Client, ctx.ethA.TestContracts.Erc20Sub, ethRecipient)

	ethAEthRecipientBalance := ethtest.Erc20BalanceOf(t, ethAClientA.Client, ctx.ethA.TestContracts.Erc20Eth, ethRecipient)
	ethBRecipientBalance := ethtest.Erc20BalanceOf(t, ethBClientA.Client, ctx.ethB.TestContracts.Erc20Eth, ethRecipient)

	// Execute several sub-tests that call t.Paralell

	// EthA -> Substrate
	t.Run("Submit Eth to Sub", func(t *testing.T) {
		testEthToSubParallel(t, ctx, ethAClientA)
	})
	// Substrate -> EthA
	t.Run("Submit Eth to Sub", func(t *testing.T) {
		testSubToEthParallel(t, ctx, subClient)
	})
	// EthA -> EthB
	t.Run("Submit EthA to EthB", func(t *testing.T) {
		testEthAToEthBParallel(t, ctx, ethAClientB)
	})
	// EthB -> EthA
	t.Run("Submit EthB to EthA", func(t *testing.T) {
		testEthBToEthAParallel(t, ctx, ethBClientA)
	})


	// Calculate and verify expected results
	time.Sleep(blockTime * 2)

	subtest.AssertBalanceOf(t, subClient, subRecipient, subRecipientBalance.Add(subRecipientBalance, balanceDelta))
	ethtest.Erc20AssertBalance(t, ethAClientA.Client, ethASubRecipientBalance.Add(ethASubRecipientBalance, balanceDelta), ctx.ethA.TestContracts.Erc20Sub, ethRecipient)
	ethtest.Erc20AssertBalance(t, ethAClientA.Client, ethAEthRecipientBalance.Add(ethAEthRecipientBalance, balanceDelta), ctx.ethA.TestContracts.Erc20Eth, ethRecipient)
	ethtest.Erc20AssertBalance(t, ethBClientA.Client, ethBRecipientBalance.Add(ethBRecipientBalance, balanceDelta), ctx.ethB.TestContracts.Erc20Eth, ethRecipient)
}

func testEthToSubParallel(t *testing.T, ctx *testContext, ethClient *ethutils.Client) {
	//t.Parallel()

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		t.Run(fmt.Sprintf("Eth to Substrate Transfer %d", i), func(t *testing.T) {
			// Initiate transfer
			log.Info("Submitting transaction", "number", i, "amount", amountPerTest)
			eth.CreateErc20Deposit(t, ethClient.Client, ethClient.Opts, SubChainId, subRecipient, amountPerTest.Int, ctx.ethA.BaseContracts, ctx.EthSubErc20ResourceId)

			time.Sleep(time.Millisecond * 100)
		})
	}
}

func testSubToEthParallel(t *testing.T, ctx *testContext, client *subutils.Client) {
	//t.Parallel()
	var calls []types.Call
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		t.Run(fmt.Sprintf("Substrate to Eth Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			call := subtest.NewNativeTransferCall(t, client, amountPerTest, types.NewAccountID(ethRecipient.Bytes()), EthAChainId)
			calls = append(calls, call)
		})
	}
	log.Info("Submitting transactions", "numberOfTxs", numberOfTxs, "amount/tx", amountPerTest)
	err := subutils.BatchSubmit(client, calls)
	if err != nil {
		t.Fatal(err)
	}
}

func testEthAToEthBParallel(t *testing.T, ctx *testContext, client *ethutils.Client) {
	//t.Parallel()
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			log.Info("Submitting transaction", "number", i, "recipient", ethRecipient, "resourceId", ctx.EthEthErc20ResourceId.Hex(), "amount", amountPerTest, "from", ctx.ethA.Opts.From, "handler", ctx.ethA.BaseContracts.ERC20HandlerAddress)
			eth.CreateErc20Deposit(t, client.Client, client.Opts, EthBChainId, ethRecipient.Bytes(), amountPerTest.Int, ctx.ethA.BaseContracts, ctx.EthEthErc20ResourceId)

			time.Sleep(time.Millisecond * 100)
		})
	}
}

func testEthBToEthAParallel(t *testing.T, ctx *testContext, client *ethutils.Client) {
	//t.Parallel()
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			log.Info("Submitting transaction", "number", i, "recipient", ethRecipient, "resourceId", ctx.EthEthErc20ResourceId.Hex(), "amount", amountPerTest, "from", ctx.ethA.Opts.From, "handler", ctx.ethA.BaseContracts.ERC20HandlerAddress)
			eth.CreateErc20Deposit(t, client.Client, client.Opts, EthAChainId, ethRecipient.Bytes(), amountPerTest.Int, ctx.ethA.BaseContracts, ctx.EthEthErc20ResourceId)

			time.Sleep(time.Millisecond * 100)
		})
	}
}