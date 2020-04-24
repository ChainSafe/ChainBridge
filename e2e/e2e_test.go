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
	{"Erc20 to Substrate Round Trip", testErc20SubstrateRoundTrip},
	{"Erc721 to Substrate Round Trip", testErc721ToSubstrateRoundTrip},
	{"Erc721 to Erc721 Round Trip", testErc721EthToEthRoundTrip},
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
	mintAmount := big.NewInt(1000)
	fundAmount := big.NewInt(500)
	approveAmount := big.NewInt(500)

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

	// ERC20 - bridge is funded and depositer has tokens approved for transfer
	// Deploy erc20 for substrate transfers
	erc20ContractASub := ethtest.Erc20DeployMint(t, ethClientA, optsA, mintAmount)
	ethtest.FundErc20Handler(t, ethClientA, optsA, contractsA.ERC20HandlerAddress, erc20ContractASub, fundAmount)
	ethtest.Erc20Approve(t, ethClientA, optsA, erc20ContractASub, contractsA.ERC20HandlerAddress, approveAmount)
	// Create another erc20 for eth to eth
	erc20ContractAEth := ethtest.Erc20DeployMint(t, ethClientA, optsA, mintAmount)
	ethtest.FundErc20Handler(t, ethClientA, optsA, contractsA.ERC20HandlerAddress, erc20ContractAEth, fundAmount)
	ethtest.Erc20Approve(t, ethClientA, optsA, erc20ContractAEth, contractsA.ERC20HandlerAddress, approveAmount)
	// Register both erc20 contracts with resource IDs
	ethErc20ResourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20ContractAEth.Bytes(), 31), 0))
	ethtest.RegisterErc20Resource(t, ethClientA, optsA, contractsA.ERC20HandlerAddress, subErc20ResourceId, erc20ContractASub)
	ethtest.RegisterErc20Resource(t, ethClientA, optsA, contractsA.ERC20HandlerAddress, ethErc20ResourceId, erc20ContractAEth)

	// ERC721 for Eth <-> Substrate - bridge is funded and depositer has tokens approved for transfer
	erc721ContractASub := ethtest.Erc721Deploy(t, ethClientA, optsA)
	ethtest.Erc721MintMany(t, ethClientA, optsA, erc721ContractASub, 1, 10)
	ethtest.Erc721ApproveMany(t, ethClientA, optsA, erc721ContractASub, contractsA.ERC721HandlerAddress, 1, 5)
	ethtest.Erc721FundHandlerMany(t, ethClientA, optsA, contractsA.ERC721HandlerAddress, erc721ContractASub, 6, 5)
	ethtest.Erc721AddMinter(t, ethClientA, optsA, erc721ContractASub, contractsA.ERC721HandlerAddress)
	ethtest.RegisterErc721Resource(t, ethClientA, optsA, contractsA.ERC721HandlerAddress, subErc721ResourceId, erc721ContractASub)
	ethtest.Erc721SetBurnable(t, ethClientA, optsA, contractsA.ERC721HandlerAddress, erc721ContractASub)

	// Erc721 for Eth <-> Eth - bridge is funded and depositer has tokens approved for transfer
	// Enable mint/burn so tokens can only exist on one chain
	erc721ContractAEth := ethtest.Erc721Deploy(t, ethClientA, optsA)
	ethtest.Erc721MintMany(t, ethClientA, optsA, erc721ContractAEth, 11, 10)
	ethtest.Erc721ApproveMany(t, ethClientA, optsA, erc721ContractAEth, contractsA.ERC721HandlerAddress, 11, 5)
	ethtest.Erc721FundHandlerMany(t, ethClientA, optsA, contractsA.ERC721HandlerAddress, erc721ContractAEth, 16, 5)
	ethtest.Erc721AddMinter(t, ethClientA, optsA, erc721ContractAEth, contractsA.ERC721HandlerAddress)
	ethErc721ResourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc721ContractAEth.Bytes(), 31), byte(EthAChainId)))
	ethtest.RegisterErc721Resource(t, ethClientA, optsA, contractsA.ERC721HandlerAddress, ethErc721ResourceId, erc721ContractAEth)
	ethtest.Erc721SetBurnable(t, ethClientA, optsA, contractsA.ERC721HandlerAddress, erc721ContractAEth)

	// Register the the generic hash resource ID (TODO: Register an actual address)
	ethtest.RegisterGenericResource(t, ethClientA, optsA, contractsA.CentrifugeHandlerAddress, genericHashResourceId, contractsA.CentrifugeHandlerAddress)

	// Base setup for ethB ERC20 - handler can mint
	contractsB := eth.DeployTestContracts(t, EthBEndpoint, EthBChainId, big.NewInt(int64(threshold)))
	// Deploy erc20 for eth to eth transfers
	erc20ContractBEth := ethtest.Erc20DeployMint(t, ethClientB, optsB, mintAmount)
	ethtest.Erc20Approve(t, ethClientB, optsB, erc20ContractBEth, contractsB.ERC20HandlerAddress, approveAmount)
	ethtest.Erc20AddMinter(t, ethClientB, optsB, erc20ContractBEth, contractsB.ERC20HandlerAddress)
	ethtest.RegisterErc20Resource(t, ethClientB, optsB, contractsB.ERC20HandlerAddress, ethErc20ResourceId, erc20ContractBEth)
	ethtest.Erc20SetBurnable(t, ethClientB, optsB, contractsB.ERC20HandlerAddress, erc20ContractBEth)

	erc721ContractBEth := ethtest.Erc721Deploy(t, ethClientB, optsB)
	ethtest.Erc721AddMinter(t, ethClientB, optsB, erc721ContractBEth, contractsB.ERC721HandlerAddress)
	ethtest.RegisterErc721Resource(t, ethClientB, optsB, contractsB.ERC721HandlerAddress, ethErc721ResourceId, erc721ContractBEth)
	ethtest.Erc721SetBurnable(t, ethClientB, optsB, contractsB.ERC721HandlerAddress, erc721ContractBEth)

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

	ctx := &testContext{
		ethA: &eth.TestContext{
			BaseContracts: contractsA,
			TestContracts: eth.TestContracts{
				Erc20Sub:  erc20ContractASub,
				Erc20Eth:  erc20ContractAEth,
				Erc721Sub: erc721ContractASub,
				Erc721Eth: erc721ContractAEth,
			},
			Client: ethClientA,
			Opts:   optsA,
		},
		ethB: &eth.TestContext{
			BaseContracts: contractsB,
			TestContracts: eth.TestContracts{
				Erc20Sub:  common.Address{}, // Not used
				Erc20Eth:  erc20ContractBEth,
				Erc721Sub: common.Address{}, // Not used
				Erc721Eth: erc721ContractBEth,
			},
			Client: ethClientB,
			Opts:   optsB,
		},
		subClient:              subClient,
		EthSubErc20ResourceId:  subErc20ResourceId,
		EthEthErc20ResourceId:  ethErc20ResourceId,
		EthSubErc721ResourceId: subErc721ResourceId,
		EthEthErc721ResourceId: ethErc721ResourceId,
		GenericHashResourceId:  genericHashResourceId,
	}

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
