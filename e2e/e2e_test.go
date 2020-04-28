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
	{"Erc20toErc20", testErc20ToErc20},
	{"Erc20 to Substrate Round Trip", testErc20SubstrateRoundTrip},

	{"Erc721 to Substrate Round Trip", testErc721ToSubstrateRoundTrip},
	{"Erc721 to Erc721 Round Trip", testErc721EthToEthRoundTrip},

	{"SubstrateHashToGenericHandler", testSubstrateHashToGenericHandler},
	{"Eth to Eth HashToGenericHandler", testEthereumHashToGenericHandler},
}

type testContext struct {
	ethA      *eth.TestContext
	ethB      *eth.TestContext
	subClient *subutils.Client

	EthSubErc20ResourceId  msg.ResourceId
	EthEthErc20ResourceId  msg.ResourceId
	EthSubErc721ResourceId msg.ResourceId
	EthEthErc721ResourceId msg.ResourceId
	GenericHashResourceId  msg.ResourceId
	EthGenericResourceId   msg.ResourceId
}

func createAndStartBridge(t *testing.T, name string, contractsA, contractsB *ethutils.DeployedContracts) (*core.Core, log.Logger) {
	// Create logger to write to a file, and store the log file name in global var
	logger := log.Root().New()

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

	return bridge, logger
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

func setupFungibleTests(t *testing.T, ctx *testContext) {
	mintAmount := big.NewInt(1000)
	approveAmount := big.NewInt(500)

	// Deploy Sub<>Eth erc20 on ethA, register resource ID, add handler as minter
	erc20ContractASub := ethtest.Erc20DeployMint(t, ctx.ethA.Client, ctx.ethA.Opts, mintAmount)
	log.Info("Deployed erc20 contract", "address", erc20ContractASub)
	ethtest.Erc20Approve(t, ctx.ethA.Client, ctx.ethA.Opts, erc20ContractASub, ctx.ethA.BaseContracts.ERC20HandlerAddress, approveAmount)
	ethtest.Erc20AddMinter(t, ctx.ethA.Client, ctx.ethA.Opts, erc20ContractASub, ctx.ethA.BaseContracts.ERC20HandlerAddress)
	ethtest.RegisterResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC20HandlerAddress, ctx.EthSubErc20ResourceId, erc20ContractASub)
	ethtest.SetBurnable(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC20HandlerAddress, erc20ContractASub)

	// Deploy Eth<>Eth erc20 on ethA, register resource ID, add handler as minter
	erc20ContractAEth := ethtest.Erc20DeployMint(t, ctx.ethA.Client, ctx.ethA.Opts, mintAmount)
	log.Info("Deployed erc20 contract", "address", erc20ContractAEth)
	ethtest.Erc20Approve(t, ctx.ethA.Client, ctx.ethA.Opts, erc20ContractAEth, ctx.ethA.BaseContracts.ERC20HandlerAddress, approveAmount)
	ethtest.Erc20AddMinter(t, ctx.ethA.Client, ctx.ethA.Opts, erc20ContractAEth, ctx.ethA.BaseContracts.ERC20HandlerAddress)
	ethErc20ResourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20ContractAEth.Bytes(), 31), 0))
	ethtest.RegisterResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC20HandlerAddress, ethErc20ResourceId, erc20ContractAEth)
	ethtest.SetBurnable(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC20HandlerAddress, erc20ContractAEth)

	// Deploy Eth<>Eth erc20 on ethB, add handler as minter
	erc20ContractBEth := ethtest.Erc20DeployMint(t, ctx.ethB.Client, ctx.ethB.Opts, mintAmount)
	log.Info("Deployed erc20 contract", "address", erc20ContractBEth)
	ethtest.Erc20AddMinter(t, ctx.ethB.Client, ctx.ethB.Opts, erc20ContractBEth, ctx.ethB.BaseContracts.ERC20HandlerAddress)
	ethtest.RegisterResource(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.BaseContracts.BridgeAddress, ctx.ethB.BaseContracts.ERC20HandlerAddress, ethErc20ResourceId, erc20ContractBEth)
	ethtest.SetBurnable(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.BaseContracts.BridgeAddress, ctx.ethB.BaseContracts.ERC20HandlerAddress, erc20ContractBEth)

	ctx.ethA.TestContracts.Erc20Sub = erc20ContractASub
	ctx.ethA.TestContracts.Erc20Eth = erc20ContractAEth
	ctx.ethB.TestContracts.Erc20Eth = erc20ContractBEth
	ctx.EthEthErc20ResourceId = ethErc20ResourceId
}

func setupNonFungibleTests(t *testing.T, ctx *testContext) {

	// Deploy Sub<>Eth erc721 on ethA, register resource ID, set burnable
	erc721ContractASub := ethtest.Erc721Deploy(t, ctx.ethA.Client, ctx.ethB.Opts)
	ethtest.Erc721AddMinter(t, ctx.ethA.Client, ctx.ethA.Opts, erc721ContractASub, ctx.ethA.BaseContracts.ERC721HandlerAddress)
	ethtest.RegisterResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC721HandlerAddress, ctx.EthSubErc721ResourceId, erc721ContractASub)
	ethtest.SetBurnable(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC721HandlerAddress, erc721ContractASub)

	// Deploy Eth<>Eth erc721 on ethA, register resource ID, set burnable
	erc721ContractAEth := ethtest.Erc721Deploy(t, ctx.ethA.Client, ctx.ethB.Opts)
	ethtest.Erc721AddMinter(t, ctx.ethA.Client, ctx.ethA.Opts, erc721ContractAEth, ctx.ethA.BaseContracts.ERC721HandlerAddress)
	ethErc721ResourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc721ContractAEth.Bytes(), 31), byte(EthAChainId)))
	ethtest.RegisterResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC721HandlerAddress, ethErc721ResourceId, erc721ContractAEth)
	ethtest.SetBurnable(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.ERC721HandlerAddress, erc721ContractASub)

	// Deploy Eth<>Eth erc721 on ethB, register resource ID, set burnable
	erc721ContractBEth := ethtest.Erc721Deploy(t, ctx.ethB.Client, ctx.ethB.Opts)
	ethtest.Erc721AddMinter(t, ctx.ethB.Client, ctx.ethB.Opts, erc721ContractBEth, ctx.ethB.BaseContracts.ERC721HandlerAddress)
	ethtest.RegisterResource(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.BaseContracts.BridgeAddress, ctx.ethB.BaseContracts.ERC721HandlerAddress, ethErc721ResourceId, erc721ContractBEth)
	ethtest.SetBurnable(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.BaseContracts.BridgeAddress, ctx.ethB.BaseContracts.ERC721HandlerAddress, erc721ContractBEth)

	ctx.ethA.TestContracts.Erc721Sub = erc721ContractASub
	ctx.ethA.TestContracts.Erc721Eth = erc721ContractAEth
	ctx.ethB.TestContracts.Erc721Eth = erc721ContractBEth
	ctx.EthEthErc721ResourceId = ethErc721ResourceId
}

func setupGenericTests(t *testing.T, ctx *testContext) {
	// Deploy asset store for sub->eth on ethA, register resource ID
	assetStoreContractASub := ethtest.DeployAssetStore(t, ctx.ethA.Client, ctx.ethA.Opts)
	ethtest.RegisterGenericResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.GenericHandlerAddress, ctx.GenericHashResourceId, assetStoreContractASub, [4]byte{}, ethutils.StoreFunctionSig)

	// Deploy asset store for eth->eth on ethB, register resource ID
	assetStoreContractBEth := ethtest.DeployAssetStore(t, ctx.ethB.Client, ctx.ethB.Opts)
	ethGenericResourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(assetStoreContractBEth.Bytes(), 31), byte(EthBChainId)))
	ethtest.RegisterGenericResource(t, ctx.ethB.Client, ctx.ethB.Opts, ctx.ethB.BaseContracts.BridgeAddress, ctx.ethB.BaseContracts.GenericHandlerAddress, ethGenericResourceId, assetStoreContractBEth, [4]byte{}, ethutils.StoreFunctionSig)
	// Register resource on ethA as well for deposit, address used could be anything
	ethtest.RegisterGenericResource(t, ctx.ethA.Client, ctx.ethA.Opts, ctx.ethA.BaseContracts.BridgeAddress, ctx.ethA.BaseContracts.GenericHandlerAddress, ethGenericResourceId, assetStoreContractBEth, [4]byte{}, ethutils.StoreFunctionSig)

	ctx.ethA.TestContracts.AssetStoreSub = assetStoreContractASub
	ctx.ethB.TestContracts.AssetStoreEth = assetStoreContractBEth
	ctx.EthGenericResourceId = ethGenericResourceId
}

// This tests three relayers connected to three chains (2 ethereum, 1 substrate).
//
// EthA:
//  - Native erc20 token
// Eth B:
//  - Synthetic erc20 token
// Substrate:
//  - Synthetic token (native to chain)
//
func Test_ThreeRelayers(t *testing.T) {
	shared.SetLogger(log.LvlTrace)
	threshold := 3

	// Setup test client connections for each chain
	ethClientA, optsA := eth.CreateEthClient(t, EthAEndpoint, eth.AliceKp)
	ethClientB, optsB := eth.CreateEthClient(t, EthBEndpoint, eth.AliceKp)
	subClient := subtest.CreateClient(t, sub.AliceKp.AsKeyringPair(), sub.TestSubEndpoint)

	// First lookup the substrate resource IDs
	var rawRId types.Bytes32
	subtest.QueryConst(t, subClient, "Example", "NativeTokenId", &rawRId)
	subErc20ResourceId := msg.ResourceIdFromSlice(rawRId[:])
	subtest.QueryConst(t, subClient, "Example", "Erc721Id", &rawRId)
	subErc721ResourceId := msg.ResourceIdFromSlice(rawRId[:])
	subtest.QueryConst(t, subClient, "Example", "HashId", &rawRId)
	genericHashResourceId := msg.ResourceIdFromSlice(rawRId[:])

	// Base setup for ethA
	contractsA := eth.DeployTestContracts(t, EthAEndpoint, EthAChainId, big.NewInt(int64(threshold)))
	// Base setup for ethB ERC20 - handler can mint
	contractsB := eth.DeployTestContracts(t, EthBEndpoint, EthBChainId, big.NewInt(int64(threshold)))

	// Create the initial context, added to in setup functions
	ctx := &testContext{
		ethA: &eth.TestContext{
			BaseContracts: contractsA,
			TestContracts: eth.TestContracts{},
			Client:        ethClientA,
			Opts:          optsA,
		},
		ethB: &eth.TestContext{
			BaseContracts: contractsB,
			TestContracts: eth.TestContracts{},
			Client:        ethClientB,
			Opts:          optsB,
		},
		subClient:              subClient,
		EthSubErc20ResourceId:  subErc20ResourceId,
		EthSubErc721ResourceId: subErc721ResourceId,
		GenericHashResourceId:  genericHashResourceId,
	}

	setupFungibleTests(t, ctx)
	setupNonFungibleTests(t, ctx)
	setupGenericTests(t, ctx)

	// Setup substrate client, register resource, add relayers
	resources := map[msg.ResourceId]subutils.Method{
		subErc20ResourceId:    subutils.ExampleTransferMethod,
		subErc721ResourceId:   subutils.ExampleMintErc721Method,
		genericHashResourceId: subutils.ExampleRemarkMethod,
	}
	subtest.EnsureInitializedChain(t, subClient, sub.RelayerSet, []msg.ChainId{EthAChainId}, resources, uint32(threshold))

	// Create and start three bridges with both chains
	_, aliceLog := createAndStartBridge(t, "alice", contractsA, contractsB)
	_, bobLog := createAndStartBridge(t, "bob", contractsA, contractsB)
	_, charlieLog := createAndStartBridge(t, "charlie", contractsA, contractsB)

	for _, tt := range tests {
		tt := tt

		// Swap handler
		fileName := "alice" + ".output"
		logFiles = append(logFiles, fileName)
		aliceLog.SetHandler(log.Must.FileHandler(fileName, log.TerminalFormat()))

		// Swap handler
		fileName = "bob" + ".output"
		logFiles = append(logFiles, fileName)
		bobLog.SetHandler(log.Must.FileHandler(fileName, log.TerminalFormat()))

		// Swap handler
		fileName = "charlie" + ".output"
		logFiles = append(logFiles, fileName)
		charlieLog.SetHandler(log.Must.FileHandler(fileName, log.TerminalFormat()))

		t.Run(tt.name, func(t *testing.T) {
			tt.fn(t, ctx)
		})

		// Flush logs
		attemptToPrintLogs()
	}
}
