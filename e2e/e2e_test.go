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
	logger.SetHandler(log.Must.FileHandler(name+".output", log.TerminalFormat()))

	ethCfg := eth.CreateConfig(name, EthChainId, bridgeAddress, erc20HandlerAddres, genericHandlerAddress)
	eth, err := ethChain.InitializeChain(ethCfg, logger.New("chain", ethCfg.Name))
	if err != nil {
		t.Fatal(err)
	}

	subCfg := sub.CreateConfig(name, SubChainId)
	sub, err := subChain.InitializeChain(subCfg, logger.New("chain", subCfg.Name))
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
			panic(err)
		}
		name := strings.Split(fileName, ".")[0]
		fmt.Printf("\n\tOutput from %s:\n\n", name)
		fmt.Print(string(dat))
		os.Remove(fileName)
	}
}

func TestMain(m *testing.M) {
	res := m.Run()
	attemptToPrintLogs()
	os.Exit(res)
}

func TestErc20ToSubstrate(t *testing.T) {
	utils.SetLogger(log.LvlTrace)

	// Deploy contracts, mint, approve
	contracts := eth.DeployTestContracts(t, EthChainId)
	ethClient, opts := eth.CreateEthClient(t)
	erc20Contract := eth.DeployMintApproveErc20(t, ethClient, opts, contracts.ERC20HandlerAddress)
	resourceId := append(common.LeftPadBytes(erc20Contract.Bytes(), 31), 0)
	eth.WhitelistResourceId(t, ethClient, opts, msg.ResourceIdFromSlice(resourceId), erc20Contract)

	// Setup substrate client
	subClient := sub.CreateSubClient(t, sub.AliceKp.AsKeyringPair())
	sub.RegisterResource(t, subClient, msg.ResourceIdFromSlice(resourceId), subChain.ExampleTransfer.String())
	success := make(chan bool)
	fail := make(chan error)

	// Create and start two bridges with both chains
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5

	for i := 1; i <= numberOfTxs; i++ {
		log.Info("Submitting transaction", "number", i)
		// Initiate transfer
		amount := big.NewInt(0).Mul(big.NewInt(int64(i)), big.NewInt(5))
		eth.CreateErc20Deposit(t, ethClient, opts, SubChainId, sub.BobKp.AsKeyringPair().PublicKey, amount, contracts, erc20Contract)

		// Check for success event
		go sub.WatchForProposalSuccessOrFail(subClient, types.NewU64(uint64(i)), success, fail)

		select {
		case <-success:
			// Do nothing
		case e := <-fail:
			t.Fatal(e)
		case <-time.After(TestTimeout):
			t.Fatalf("test timed out")
		}
	}

}

func TestSubstrateToErc20(t *testing.T) {
	// TODO: Remove once example pallet has transfer out function (https://github.com/ChainSafe/chainbridge-substrate/issues/28)
	t.Skip()
	utils.SetLogger(log.LvlInfo)

	// Whitelist chain
	subClient := sub.CreateSubClient(t, sub.AliceKp.AsKeyringPair())
	sub.WhitelistChain(t, subClient, EthChainId)

	// Deploy contracts, mint, approve
	contracts := eth.DeployTestContracts(t, EthChainId)
	ethClient, opts := eth.CreateEthClient(t)
	erc20Contract := eth.DeployAndFundErc20(t, ethClient, opts, contracts.ERC20HandlerAddress)

	// Start two bridges
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5
	amount := 10
	expectedBalance := big.NewInt(0)
	recipient := eth.CharlieEthAddr
	eth.AssertBalance(t, ethClient, expectedBalance, erc20Contract, recipient)
	for i := 1; i <= numberOfTxs; i++ {
		// Execute transfer
		sub.InitiateSubstrateNativeTransfer(t, subClient, amount, recipient.Bytes(), EthChainId)

		// Wait for event
		eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalCreated)
		eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalExecuted)

		// Verify balance change
		expectedBalance.Add(expectedBalance, big.NewInt(int64(amount)))
		eth.AssertBalance(t, ethClient, expectedBalance, erc20Contract, recipient)
	}
}

func TestHashToGenericHandler(t *testing.T) {
	utils.SetLogger(log.LvlTrace)

	// Deploy contracts (incl. handler)
	contracts := eth.DeployTestContracts(t, EthChainId)
	ethClient, _ := eth.CreateEthClient(t)

	// Whitelist chain
	subClient := sub.CreateSubClient(t, sub.AliceKp.AsKeyringPair())
	sub.WhitelistChain(t, subClient, EthChainId)

	// Create two bridges
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		// Execute transfer
		hash := sub.HashInt(i)

		sub.InitiateHashTransfer(t, subClient, hash, EthChainId)

		// Wait for event
		eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalCreated)
		eth.WaitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethChain.DepositProposalExecuted)

		// Verify hash is available
		eth.AssertHashExistence(t, ethClient, hash, contracts.CentrifugeHandlerAddress)
	}

}
