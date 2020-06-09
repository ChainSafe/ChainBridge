// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func createWriters(t *testing.T, client *utils.Client, contracts *utils.DeployedContracts) (*writer, *writer, func(), func(), chan error, chan error) {
	latestBlock := ethtest.GetLatestBlock(t, client)
	errA := make(chan error)
	writerA, stopA := createTestWriter(t, createConfig("bob", latestBlock, contracts), errA)
	errB := make(chan error)
	writerB, stopB := createTestWriter(t, createConfig("charlie", latestBlock, contracts), errB)
	return writerA, writerB, stopA, stopB, errA, errB
}

func createTestWriter(t *testing.T, cfg *Config, errs chan<- error) (*writer, func()) {

	conn := newLocalConnection(t, cfg)
	stop := make(chan int)
	writer := NewWriter(conn, cfg, newTestLogger(cfg.name), stop, errs)

	bridge, err := Bridge.NewBridge(cfg.bridgeContract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	writer.setContract(bridge)

	err = writer.start()
	if err != nil {
		t.Fatal(err)
	}
	return writer, func() { close(stop) }
}

func watchEvent(client *utils.Client, bridge common.Address, subStr utils.EventSig) {
	fmt.Printf("Watching for event: %s\n", subStr)
	query := eth.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{bridge},
		Topics: [][]common.Hash{
			{subStr.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.Client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		log15.Error("Failed to subscribe to event", "event", subStr)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case evt := <-ch:
			fmt.Printf("%s: %#v\n", subStr, evt.Topics)

		case err := <-sub.Err():
			if err != nil {
				log15.Error("Subscription error", "event", subStr, "err", err)
				return
			}
		}
	}
}

func routeMessageAndWait(t *testing.T, client *utils.Client, alice, bob *writer, m msg.Message, aliceErr, bobErr chan error) {
	// Watch for executed event
	query := eth.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{alice.cfg.bridgeContract},
		Topics: [][]common.Hash{
			{utils.ProposalExecuted.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.Client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()

	if err != nil {
		log15.Error("Failed to subscribe to finalization event", "err", err)
	}

	// Alice processes the message, then waits to execute
	if ok := alice.ResolveMessage(m); !ok {
		t.Fatal("Alice failed to resolve the message")
	}

	// Now Bob receives the same message and also waits to execute
	if ok := bob.ResolveMessage(m); !ok {
		t.Fatal("Bob failed to resolve the message")
	}

	for {
		select {
		case evt := <-ch:
			sourceId := evt.Topics[1].Big().Uint64()
			destId := evt.Topics[2].Big().Uint64()
			depositNonce := evt.Topics[3].Big().Uint64()

			if m.Source == msg.ChainId(sourceId) &&
				m.Destination == msg.ChainId(destId) &&
				uint64(m.DepositNonce) == depositNonce {
				return
			}

		case err = <-sub.Err():
			if err != nil {
				t.Fatal(err)
			}
		case err = <-aliceErr:
			t.Fatalf("Fatal error: %s", err)
		case err = <-bobErr:
			t.Fatalf("Fatal error: %s", err)
		case <-time.After(TestTimeout):
			t.Fatal("test timed out")
		}
	}
}

func Test_ContainsVote(t *testing.T) {
	votes := []common.Address{AliceKp.CommonAddress()}

	if !containsVote(votes, AliceKp.CommonAddress()) {
		t.Error("Voter has voted")
	}

	if containsVote(votes, BobKp.CommonAddress()) {
		t.Error("Voter has not voted")
	}
}

func TestWriter_start_stop(t *testing.T) {
	conn := newLocalConnection(t, aliceTestConfig)
	defer conn.Close()

	stop := make(chan int)
	writer := NewWriter(conn, aliceTestConfig, TestLogger, stop, nil)

	err := writer.start()
	if err != nil {
		t.Fatal(err)
	}

	// Initiate shutdown
	close(stop)
}

func TestCreateAndExecuteErc20DepositProposal(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, TestChainId, AliceKp)
	writerA, writerB, stopA, stopB, errA, errB := createWriters(t, client, contracts)

	defer stopA()
	defer stopB()
	defer writerA.conn.Close()
	defer writerB.conn.Close()

	erc20Address := ethtest.DeployMintApproveErc20(t, client, contracts.ERC20HandlerAddress, big.NewInt(100))
	ethtest.FundErc20Handler(t, client, contracts.ERC20HandlerAddress, erc20Address, big.NewInt(100))

	// Create initial transfer message
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Address.Bytes(), 31), 0))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)
	amount := big.NewInt(10)
	m := msg.NewFungibleTransfer(1, 0, 0, amount, resourceId, recipient.Bytes())
	ethtest.RegisterResource(t, client, contracts.BridgeAddress, contracts.ERC20HandlerAddress, resourceId, erc20Address)
	// Helpful for debugging
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalExecuted)

	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.Erc20AssertBalance(t, client, amount, erc20Address, recipient)
}

func TestCreateAndExecuteErc721Proposal(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, TestChainId, AliceKp)
	writerA, writerB, stopA, stopB, errA, errB := createWriters(t, client, contracts)

	defer stopA()
	defer stopB()
	defer writerA.conn.Close()
	defer writerB.conn.Close()

	// We'll use alice to setup the erc721
	tokenId := big.NewInt(1)
	erc721Contract := ethtest.Erc721Deploy(t, client)
	ethtest.Erc721Mint(t, client, erc721Contract, tokenId, []byte{})
	ethtest.Erc721FundHandler(t, client, contracts.ERC721HandlerAddress, erc721Contract, tokenId)

	// Create initial transfer message
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc721Contract.Bytes(), 31), 0))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)
	m := msg.NewNonFungibleTransfer(1, 0, 0, resourceId, tokenId, recipient.Bytes(), []byte{})
	ethtest.RegisterResource(t, client, contracts.BridgeAddress, contracts.ERC721HandlerAddress, resourceId, erc721Contract)
	// Helpful for debugging
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalExecuted)

	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.Erc721AssertOwner(t, client, erc721Contract, tokenId, recipient)
}

func TestCreateAndExecuteGenericProposal(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, TestChainId, AliceKp)
	writerA, writerB, stopA, stopB, errA, errB := createWriters(t, client, contracts)

	defer stopA()
	defer stopB()
	defer writerA.conn.Close()
	defer writerB.conn.Close()

	assetStoreAddr, err := utils.DeployAssetStore(client)
	if err != nil {
		t.Fatal(err)
	}

	rId := msg.ResourceIdFromSlice(common.LeftPadBytes(assetStoreAddr.Bytes(), 32))
	depositSig := utils.CreateFunctionSignature("")
	executeSig := utils.CreateFunctionSignature("store(bytes32)")

	ethtest.RegisterGenericResource(t, client, contracts.BridgeAddress, contracts.GenericHandlerAddress, rId, assetStoreAddr, depositSig, executeSig)
	// Create initial transfer message
	hash := common.HexToHash("0xf0a8748d2b102eb4e0e116047753b9beff0396d81b830693b19a1376ac4b14e8")
	m := msg.Message{
		Source:       1,
		Destination:  0,
		Type:         msg.GenericTransfer,
		DepositNonce: 0,
		ResourceId:   rId,
		Payload: []interface{}{
			hash.Bytes(),
		},
	}

	// Helpful for debugging
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalExecuted)

	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.AssertHashExistence(t, client, hash, assetStoreAddr)
}

func TestDuplicateMessage(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, TestChainId, AliceKp)
	writerA, writerB, stopA, stopB, errA, errB := createWriters(t, client, contracts)

	defer stopA()
	defer stopB()
	defer writerA.conn.Close()
	defer writerB.conn.Close()

	erc20Address := ethtest.DeployMintApproveErc20(t, client, contracts.ERC20HandlerAddress, big.NewInt(100))
	ethtest.FundErc20Handler(t, client, contracts.ERC20HandlerAddress, erc20Address, big.NewInt(100))

	// Create initial transfer message
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(erc20Address.Bytes(), 31), 0))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)
	amount := big.NewInt(10)
	m := msg.NewFungibleTransfer(1, 0, 0, amount, resourceId, recipient.Bytes())
	ethtest.RegisterResource(t, client, contracts.BridgeAddress, contracts.ERC20HandlerAddress, resourceId, erc20Address)
	// Helpful for debugging
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(client, contracts.BridgeAddress, utils.ProposalExecuted)

	// Process initial message
	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.Erc20AssertBalance(t, client, amount, erc20Address, recipient)

	// Capture nonces
	nonceAPre, err := writerA.conn.conn.PendingNonceAt(context.Background(), writerA.conn.kp.CommonAddress())
	if err != nil {
		t.Fatal(err)
	}
	nonceBPre, err := writerA.conn.conn.PendingNonceAt(context.Background(), writerB.conn.kp.CommonAddress())
	if err != nil {
		t.Fatal(err)
	}

	// Try processing the same message again
	if ok := writerA.ResolveMessage(m); !ok {
		t.Fatal("Alice failed to resolve the message")
	}
	if ok := writerB.ResolveMessage(m); !ok {
		t.Fatal("Bob failed to resolve the message")
	}

	// Ensure the votes are recorded
	if !writerA.hasVoted(m.Source, m.DepositNonce) {
		t.Fatal("Relayer vote not found on chain")
	}
	if !writerB.hasVoted(m.Source, m.DepositNonce) {
		t.Fatal("Relayer vote not found on chain")
	}

	// Capture new nonces
	nonceAPost, err := writerA.conn.conn.PendingNonceAt(context.Background(), writerA.conn.kp.CommonAddress())
	if err != nil {
		t.Fatal(err)
	}
	nonceBPost, err := writerA.conn.conn.PendingNonceAt(context.Background(), writerB.conn.kp.CommonAddress())
	if err != nil {
		t.Fatal(err)
	}

	// Enssure neither relayer has submitted a tx
	if nonceAPost != nonceAPre {
		t.Error("Writer A nonce has changed.")
	}
	if nonceBPost != nonceBPre {
		t.Error("Writer B nonce has changed.")
	}

}
