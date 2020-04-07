// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/common"
)

const TestTimeout = time.Second * 30
const TestEthEndpoint = "ws://localhost:8545"
const TestSubEndpoint = "ws://localhost:9944"
const EthChainId = msg.ChainId(0)
const SubChainId = msg.ChainId(1)

var AliceEthKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var CharlieEthKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var AliceSubKp = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey]
var BobSubKp = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey]
var CharlieEthAddr = common.HexToAddress(CharlieEthKp.Address())

func setLogger(lvl log.Lvl) {
	logger := log.Root()
	handler := logger.GetHandler()
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))
}

func createEthConfig(key, bridgeAddress, erc20HandlerAddress, genericHandler string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("ethereum(%s)", key),
		Id:           EthChainId,
		Endpoint:     TestEthEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts: map[string]string{
			"bridge":         bridgeAddress,
			"erc20Handler":   erc20HandlerAddress,
			"genericHandler": genericHandler,
		},
	}
}

func createSubConfig(key string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("substrate(%s)", key),
		Id:           SubChainId,
		Endpoint:     TestSubEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts:         map[string]string{},
	}
}

func createAndStartBridge(t *testing.T, name, bridgeAddress, erc20HandlerAddres, genericHandlerAddress string) *core.Core {
	eth, err := ethereum.InitializeChain(createEthConfig(name, bridgeAddress, erc20HandlerAddres, genericHandlerAddress))
	if err != nil {
		t.Fatal(err)
	}
	sub, err := substrate.InitializeChain(createSubConfig(name))
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

func TestErc20ToSubstrate(t *testing.T) {
	setLogger(log.LvlTrace)

	// Deploy contracts, mint, approve
	contracts := deployTestContracts(t, EthChainId)
	ethClient, opts := createEthClient(t)
	erc20Contract := deployMintApproveErc20(t, ethClient, opts, contracts.ERC20HandlerAddress)

	// Setup substrate client
	subClient := createSubClient(t, AliceSubKp.AsKeyringPair())
	success := make(chan bool)
	fail := make(chan error)

	// Create and start two bridges with both chains
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5

	for i := 1; i <= numberOfTxs; i++ {
		log.Info("Submitting transaction", "number", i)
		// Initiate transfer
		createErc20Deposit(t, ethClient, opts, contracts, erc20Contract)

		// Check for success event
		go watchForProposalSuccessOrFail(subClient, types.NewU32(uint32(i)), success, fail)

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
	setLogger(log.LvlInfo)

	// Whitelist chain
	subClient := createSubClient(t, AliceSubKp.AsKeyringPair())
	whitelistChain(t, subClient, EthChainId)

	// Deploy contracts, mint, approve
	contracts := deployTestContracts(t, EthChainId)
	ethClient, opts := createEthClient(t)
	erc20Contract := deployAndFundErc20(t, ethClient, opts, contracts.ERC20HandlerAddress)

	// Start two bridges
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5
	amount := 10
	expectedBalance := big.NewInt(0)
	recipient := CharlieEthAddr
	assertBalance(t, ethClient, expectedBalance, erc20Contract, recipient)
	for i := 1; i <= numberOfTxs; i++ {
		// Execute transfer
		initiateSubstrateNativeTransfer(t, subClient, amount, recipient.Bytes(), EthChainId)

		// Wait for event
		waitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethereum.DepositProposalCreated)
		waitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethereum.DepositProposalExecuted)

		// Verify balance change
		expectedBalance.Add(expectedBalance, big.NewInt(int64(amount)))
		assertBalance(t, ethClient, expectedBalance, erc20Contract, recipient)
	}
}

func TestHashToGenericHandler(t *testing.T) {
	setLogger(log.LvlTrace)

	// Deploy contracts (incl. handler)
	contracts := deployTestContracts(t, EthChainId)
	ethClient, _ := createEthClient(t)

	// Whitelist chain
	subClient := createSubClient(t, AliceSubKp.AsKeyringPair())
	whitelistChain(t, subClient, EthChainId)

	// Create two bridges
	_ = createAndStartBridge(t, "alice", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())
	_ = createAndStartBridge(t, "bob", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex(), contracts.CentrifugeHandlerAddress.Hex())

	numberOfTxs := 5
	for i := 1; i <= numberOfTxs; i++ {
		// Execute transfer
		hash := hashInt(i)

		initiateHashTransfer(t, subClient, hash, EthChainId)

		// Wait for event
		waitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethereum.DepositProposalCreated)
		waitForEthereumEvent(t, ethClient, contracts.BridgeAddress, ethereum.DepositProposalExecuted)

		// Verify hash is available
		assertHashExistence(t, ethClient, hash, contracts.CentrifugeHandlerAddress)
	}

}
