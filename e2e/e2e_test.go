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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const EthChainId = msg.ChainId(0)
const SubChainId = msg.ChainId(1)

var logFiles = []string{}

type test struct {
	name string
	fn   func(*testing.T, *testContext)
}

var tests = []test{
	{"Erc20ToSubstrate", testErc20ToSubstrate},
	{"SubstrateToErc20", testSubstrateToErc20},
	{"SubstrateHashToGenericHandler", testSubstrateHashToGenericHandler},
}

type testContext struct {
	contracts *ethutils.DeployedContracts
	ethClient *ethclient.Client
	opts      *bind.TransactOpts
	subClient *subutils.Client
}

func createAndStartBridge(t *testing.T, name, bridgeAddress, erc20HandlerAddres, genericHandlerAddress string) *core.Core {
	// Create logger to write to a file, and store the log file name in global var
	logger := log.Root().New()
	fileName := name + ".output"
	logFiles = append(logFiles, fileName)
	logger.SetHandler(log.Must.FileHandler(fileName, log.TerminalFormat()))

	ethCfg := eth.CreateConfig(name, EthChainId, bridgeAddress, erc20HandlerAddres, genericHandlerAddress)
	eth, err := ethChain.InitializeChain(ethCfg, logger.New("relayer", name))
	if err != nil {
		t.Fatal(err)
	}

	subCfg := sub.CreateConfig(name, SubChainId)
	sub, err := subChain.InitializeChain(subCfg, logger.New("relayer", name))
	if err != nil {
		t.Fatal(err)
	}

	bridge := core.NewCore()
	bridge.AddChain(eth)
	bridge.AddChain(sub)

	err = eth.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = sub.Start()
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

	// Deploy contracts, mint, approve
	contracts := eth.DeployTestContracts(t, EthChainId, big.NewInt(int64(threshold)))
	ethClient, opts := eth.CreateEthClient(t)

	// Setup substrate client, register resource, add relayers
	subClient := subtest.CreateClient(t, sub.AliceKp.AsKeyringPair(), sub.TestSubEndpoint)
	subtest.EnsureInitializedChain(subClient, sub.RelayerSet, []msg.ChainId{EthChainId}, uint32(threshold))

	// Create and start three bridges with both chains
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "charlie", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	ctx := &testContext{
		contracts: contracts,
		ethClient: ethClient,
		opts:      opts,
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
	erc20Contract := ethtest.DeployMintApproveErc20(t, ctx.ethClient, ctx.opts, ctx.contracts.ERC20HandlerAddress, big.NewInt(100))
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Contract.Bytes(), 31), 0))
	ethtest.RegisterErc20Resource(t, ctx.ethClient, ctx.opts, ctx.contracts.ERC20HandlerAddress, resourceId, erc20Contract)
	subtest.RegisterResource(t, ctx.subClient, resourceId, string(subutils.ExampleTransferMethod))

	numberOfTxs := 5

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			log.Info("Submitting transaction", "number", i)
			// Initiate transfer
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			eth.CreateErc20Deposit(t, ctx.ethClient, ctx.opts, SubChainId, sub.BobKp.AsKeyringPair().PublicKey, amount, ctx.contracts, resourceId)

			// Check for success event
			sub.WaitForProposalSuccessOrFail(t, ctx.subClient, types.NewU64(uint64(i)), types.U8(EthChainId))
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
	erc20Contract := eth.DeployErc20AndAddMinter(t, ctx.ethClient, ctx.opts, ctx.contracts.ERC20HandlerAddress)
	ethtest.RegisterErc20Resource(t, ctx.ethClient, ctx.opts, ctx.contracts.ERC20HandlerAddress, resourceId, erc20Contract)

	numberOfTxs := 5
	expectedBalance := big.NewInt(0)
	recipient := eth.CharlieEthAddr
	ethtest.AssertBalance(t, ctx.ethClient, expectedBalance, erc20Contract, recipient)

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			// Get latest eth block
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethClient)

			// Execute transfer
			amount := types.NewU32(uint32(i * 5))
			subtest.InitiateSubstrateNativeTransfer(t, ctx.subClient, amount, recipient.Bytes(), EthChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ctx.ethClient, ctx.contracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethClient, ctx.contracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify balance change
			expectedBalance.Add(expectedBalance, big.NewInt(int64(amount)))
			ethtest.AssertBalance(t, ctx.ethClient, expectedBalance, erc20Contract, recipient)
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
	ethtest.RegisterGenericResource(t, ctx.ethClient, ctx.opts, ctx.contracts.CentrifugeHandlerAddress, resourceId, common.Address{})

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			latestEthBlock := ethtest.GetLatestBlock(t, ctx.ethClient)

			// Execute transfer
			hash := sub.HashInt(i)
			subtest.InitiateHashTransfer(t, ctx.subClient, hash, EthChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ctx.ethClient, ctx.contracts.BridgeAddress, ethutils.DepositProposalCreated, latestEthBlock)
			eth.WaitForEthereumEvent(t, ctx.ethClient, ctx.contracts.BridgeAddress, ethutils.DepositProposalExecuted, latestEthBlock)

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethClient, hash, ctx.contracts.CentrifugeHandlerAddress)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
}
