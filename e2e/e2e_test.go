// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"testing"

	ethChain "github.com/ChainSafe/ChainBridge/chains/ethereum"
	subChain "github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/ChainBridge/shared"
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subutils "github.com/ChainSafe/ChainBridge/shared/substrate"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/common"
)

const EthAChainId = msg.ChainId(0)
const SubChainId = msg.ChainId(1)
const EthBChainId = msg.ChainId(2)

const EthAEndpoint = "ws://localhost:8545"
const EthBEndpoint = "ws://localhost:8546"

var logFiles = []string{}

type test struct {
	name string
	fn   func(*testing.T, *testContext)
}

var tests = []test{
	{"Erc20ToSubstrate", testErc20ToSubstrate},
	{"SubstrateToErc20", testSubstrateToErc20},
	{"SubstrateHashToGenericHandler", testSubstrateHashToGenericHandler},
	{"Erc20toErc20", testErc20ToErc20},
}

type testContext struct {
	ethA      *eth.TestContext
	ethB      *eth.TestContext
	subClient *subutils.Client
}

func createAndStartBridge(t *testing.T, name string, contractsA, contractsB *ethutils.DeployedContracts) *core.Core {
	// Create logger to write to a file, and store the log file name in global var
	logger := log.Root().New()
	fileName := name + ".output"
	logFiles = append(logFiles, fileName)
	logger.SetHandler(log.Must.FileHandler(fileName, log.TerminalFormat()))

	ethACfg := eth.CreateConfig(name, EthAChainId, contractsA, EthAEndpoint)
	ethA, err := ethChain.InitializeChain(ethACfg, logger.New("relayer", name, "chain", "ethA"))
	if err != nil {
		t.Fatal(err)
	}

	subCfg := sub.CreateConfig(name, SubChainId)
	sub, err := subChain.InitializeChain(subCfg, logger.New("relayer", name, "chain", "sub"))
	if err != nil {
		t.Fatal(err)
	}

	ethBCfg := eth.CreateConfig(name, EthBChainId, contractsB, EthBEndpoint)
	ethB, err := ethChain.InitializeChain(ethBCfg, logger.New("relayer", name, "chain", "ethB"))
	if err != nil {
		t.Fatal(err)
	}

	bridge := core.NewCore()
	bridge.AddChain(ethA)
	bridge.AddChain(sub)
	bridge.AddChain(ethB)

	err = ethA.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = sub.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = ethB.Start()
	if err != nil {
		t.Fatal(err)
	}

	return bridge
}

func attemptToPrintLogs() {
	for _, fileName := range logFiles {
		dat, err := ioutil.ReadFile(fileName)
		if err != nil {
			continue
		}
		name := strings.Split(fileName, ".")[0]
		fmt.Printf("\n\tOutput from %s:\n\n", name)
		fmt.Print(string(dat))
		os.Remove(fileName)
	}
}

func Test_ThreeRelayers(t *testing.T) {
	shared.SetLogger(log.LvlTrace)
	threshold := 3

	// Deploy contracts, mint, approve (EthA)
	contractsA := eth.DeployTestContracts(t, EthAEndpoint, EthAChainId, big.NewInt(int64(threshold)))
	ethClientA, optsA := eth.CreateEthClient(t, EthAEndpoint)

	// Deploy contracts, mint, approve (EthB)
	contractsB := eth.DeployTestContracts(t, EthBEndpoint, EthBChainId, big.NewInt(int64(threshold)))
	ethClientB, optsB := eth.CreateEthClient(t, EthBEndpoint)

	// Setup substrate client, register resource, add relayers
	subClient := subtest.CreateClient(t, sub.AliceKp.AsKeyringPair(), sub.TestSubEndpoint)
	subtest.EnsureInitializedChain(subClient, sub.RelayerSet, []msg.ChainId{EthAChainId}, uint32(threshold))

	// Create and start three bridges with both chains
	_ = createAndStartBridge(t, "alice", contractsA, contractsB)
	_ = createAndStartBridge(t, "bob", contractsA, contractsB)
	_ = createAndStartBridge(t, "charlie", contractsA, contractsB)

	ctx := &testContext{
		ethA: &eth.TestContext{
			Contracts: contractsA,
			Client:    ethClientA,
			Opts:      optsA,
		},
		ethB: &eth.TestContext{
			Contracts: contractsB,
			Client:    ethClientB,
			Opts:      optsB,
		},
		subClient: subClient,
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.fn(t, ctx)
		})
	}
	attemptToPrintLogs()
}

func testErc20ToSubstrate(t *testing.T, ctx *testContext) {
	erc20Contract := ethtest.DeployMintApproveErc20(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.ERC20HandlerAddress, big.NewInt(100))
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Contract.Bytes(), 31), 0))
	ethtest.RegisterErc20Resource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.ERC20HandlerAddress, resourceId, erc20Contract)
	subtest.RegisterResource(t, ctx.subClient, resourceId, string(subutils.ExampleTransferMethod))

	numberOfTxs := 5

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			log.Info("Submitting transaction", "number", i)
			// Initiate transfer
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			eth.CreateErc20Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, SubChainId, sub.BobKp.AsKeyringPair().PublicKey, amount, ctx.ethA.Contracts, resourceId)

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
	// Fetch the resource ID
	var rawRId types.Bytes32
	subtest.QueryConst(t, ctx.subClient, "Example", "NativeTokenId", &rawRId)
	resourceId := msg.ResourceIdFromSlice(rawRId[:])

	// Deploy erc20, mint, approve and register the resulting address under the resource ID
	erc20Contract := eth.DeployErc20AndAddMinter(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.ERC20HandlerAddress)
	ethtest.RegisterErc20Resource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.ERC20HandlerAddress, resourceId, erc20Contract)

	numberOfTxs := 5
	expectedBalance := big.NewInt(0)
	recipient := eth.CharlieEthAddr
	ethtest.AssertBalance(t, ctx.ethA.Client, expectedBalance, erc20Contract, recipient)

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			// Get latest eth block
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethA.Client)

			// Execute transfer
			amount := types.NewU32(uint32(i * 5))
			subtest.InitiateSubstrateNativeTransfer(t, ctx.subClient, amount, recipient.Bytes(), EthAChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.Contracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.Contracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify balance change
			expectedBalance.Add(expectedBalance, big.NewInt(int64(amount)))
			ethtest.AssertBalance(t, ctx.ethA.Client, expectedBalance, erc20Contract, recipient)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}

func testSubstrateHashToGenericHandler(t *testing.T, ctx *testContext) {
	// Fetch the resource ID and register it
	var rawRId types.Bytes32
	subtest.QueryConst(t, ctx.subClient, "Example", "NativeTokenId", &rawRId)
	resourceId := msg.ResourceIdFromSlice(rawRId[:])
	ethtest.RegisterGenericResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.CentrifugeHandlerAddress, resourceId, common.Address{})

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethA.Client)

			// Execute transfer
			hash := sub.HashInt(i)
			subtest.InitiateHashTransfer(t, ctx.subClient, hash, EthAChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.Contracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethA.Client, ctx.ethA.Contracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethA.Client, hash, ctx.ethA.Contracts.CentrifugeHandlerAddress)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}

func testErc20ToErc20(t *testing.T, ctx *testContext) {
	ethAErc20 := ethtest.DeployMintApproveErc20(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.ERC20HandlerAddress, big.NewInt(100))
	ethBErc20 := eth.DeployErc20AndAddMinter(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.Contracts.ERC20HandlerAddress)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(ethAErc20.Bytes(), 31), 0))
	ethtest.RegisterErc20Resource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.Contracts.ERC20HandlerAddress, resourceId, ethAErc20)
	ethtest.RegisterErc20Resource(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.Contracts.ERC20HandlerAddress, resourceId, ethBErc20)

	recipient := eth.CharlieEthAddr
	expectedBalance := big.NewInt(0)

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethB.Client)

			log.Info("Submitting transaction", "number", i, "recipient", recipient, "resourcId", resourceId.Hex())
			// Initiate transfer
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			eth.CreateErc20Deposit(t, ctx.ethA.Client, ctx.ethA.Opts, EthBChainId, recipient.Bytes(), amount, ctx.ethA.Contracts, resourceId)

			eth.WaitForEthereumEvent(t, ctx.ethB.Client, ctx.ethB.Contracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethB.Client, ctx.ethB.Contracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify balance change
			expectedBalance.Add(expectedBalance, amount)
			ethtest.AssertBalance(t, ctx.ethB.Client, expectedBalance, ethBErc20, recipient)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}
