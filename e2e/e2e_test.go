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
	"time"

	ethChain "github.com/ChainSafe/ChainBridge/chains/ethereum"
	subChain "github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/testutils"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/common"
)

const TestTimeout = time.Second * 15
const EthChainId = msg.ChainId(0)
const SubChainId = msg.ChainId(1)

var logFiles = []string{}

func createAndStartBridge(t *testing.T, name, bridgeAddress, erc20HandlerAddres, genericHandlerAddress string) *core.Core {
	// Create logger to write to a file, and store the log file name in global var
	logger := log.Root().New("chain", name)
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

func TestMain(m *testing.M) {
	res := m.Run()
	os.Exit(res)
}

func TestErc20ToSubstrate(t *testing.T) {
	utils.SetLogger(log.LvlTrace)

	// Deploy contracts, mint, approve
	contracts := eth.DeployTestContracts(t, EthChainId)
	ethClient, opts := eth.CreateEthClient(t)
	erc20Contract := eth.DeployMintApproveErc20(t, ethClient, opts, contracts.ERC20HandlerAddress)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Contract.Bytes(), 31), 0))
	eth.RegisterErc20Resource(t, ethClient, opts, contracts.ERC20HandlerAddress, resourceId, erc20Contract)

	// Setup substrate client, register resource, add relayers
	subClient := sub.CreateSubClient(t, sub.AliceKp.AsKeyringPair())
	sub.TrySetupChain(t, subClient, sub.RelayerSet, []msg.ChainId{EthChainId})
	sub.RegisterResource(t, subClient, resourceId, subChain.ExampleTransfer.String())

	// Create and start two bridges with both chains
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5

	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			log.Info("Submitting transaction", "number", i)
			// Initiate transfer
			amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
			eth.CreateErc20Deposit(t, ethClient, opts, EthChainId, SubChainId, sub.BobKp.AsKeyringPair().PublicKey, amount, contracts, resourceId)

			// Check for success event
			sub.WaitForProposalSuccessOrFail(t, subClient, types.NewU64(uint64(i)), types.U8(EthChainId))
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
	attemptToPrintLogs()
}

func TestSubstrateToErc20(t *testing.T) {
	utils.SetLogger(log.LvlInfo)

	// Setup substrate client and chain
	subClient := sub.CreateSubClient(t, sub.AliceKp.AsKeyringPair())
	sub.TrySetupChain(t, subClient, sub.RelayerSet, []msg.ChainId{EthChainId})

	// Fetch the resource ID
	var rawRId types.Bytes32
	sub.QueryConst(t, subClient, "Example", "NativeTokenId", &rawRId)
	resourceId := msg.ResourceIdFromSlice(rawRId[:])

	// Deploy contracts, mint, approve
	contracts := eth.DeployTestContracts(t, EthChainId)
	ethClient, opts := eth.CreateEthClient(t)
	erc20Contract := eth.DeployErc20AndAddMinter(t, ethClient, opts, contracts.ERC20HandlerAddress)
	eth.RegisterErc20Resource(t, ethClient, opts, contracts.ERC20HandlerAddress, resourceId, erc20Contract)

	// Start two bridges
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5
	amount := types.U32(10)
	expectedBalance := big.NewInt(0)
	recipient := eth.CharlieEthAddr
	eth.AssertBalance(t, ethClient, expectedBalance, erc20Contract, recipient)

	for i := 1; i <= numberOfTxs; i++ {
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			// Execute transfer
			sub.InitiateSubstrateNativeTransfer(t, subClient, amount, recipient.Bytes(), EthChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalCreated)
			eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalExecuted)

			// Verify balance change
			expectedBalance.Add(expectedBalance, big.NewInt(int64(amount)))
			eth.AssertBalance(t, ethClient, expectedBalance, erc20Contract, recipient)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
	attemptToPrintLogs()
}

func TestSubstrateHashToGenericHandler(t *testing.T) {
	utils.SetLogger(log.LvlTrace)
	defer attemptToPrintLogs()

	// Whitelist chain
	subClient := sub.CreateSubClient(t, sub.AliceKp.AsKeyringPair())
	sub.TrySetupChain(t, subClient, sub.RelayerSet, []msg.ChainId{EthChainId})

	// Fetch the resource ID
	var rawRId types.Bytes32
	sub.QueryConst(t, subClient, "Example", "NativeTokenId", &rawRId)
	resourceId := msg.ResourceIdFromSlice(rawRId[:])

	// Deploy contracts (incl. handler)
	contracts := eth.DeployTestContracts(t, EthChainId)
	ethClient, opts := eth.CreateEthClient(t)
	eth.RegisterGenericResource(t, ethClient, opts, contracts.CentrifugeHandlerAddress, resourceId, common.Address{})

	// Create two bridges
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Deposit %d", i), func(t *testing.T) {
			// Execute transfer
			hash := sub.HashInt(i)

			sub.InitiateHashTransfer(t, subClient, hash, EthChainId)

			// Wait for event
			eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalCreated)
			eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalExecuted)

			// Verify hash is available
			eth.AssertHashExistence(t, ethClient, hash, contracts.CentrifugeHandlerAddress)
		})
		if !ok {
			attemptToPrintLogs()
			return
		}
	}
	attemptToPrintLogs()
}
