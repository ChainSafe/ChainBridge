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

func createWritersAndClient(t *testing.T, contracts *utils.DeployedContracts) (*writer, *writer, chan error, chan error) {
	errA := make(chan error)
	bob := createTestWriter(t, bobTestConfig, contracts, errA)
	errB := make(chan error)
	charlie := createTestWriter(t, charlieTestConfig, contracts, errB)
	return bob, charlie, errA, errB
}

func createTestWriter(t *testing.T, cfg *Config, contracts *utils.DeployedContracts, errs chan<- error) *writer {
	conn := newLocalConnection(t, cfg)
	writer := NewWriter(conn, cfg, newTestLogger(cfg.name), make(chan int), errs)

	bridge, err := Bridge.NewBridge(contracts.BridgeAddress, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	writer.conn.cfg.bridgeContract = contracts.BridgeAddress
	writer.conn.cfg.erc20HandlerContract = contracts.ERC20HandlerAddress
	writer.conn.cfg.erc721HandlerContract = contracts.ERC721HandlerAddress
	writer.conn.cfg.genericHandlerContract = contracts.GenericHandlerAddress
	writer.cfg.bridgeContract = contracts.BridgeAddress
	writer.cfg.erc20HandlerContract = contracts.ERC20HandlerAddress
	writer.cfg.erc721HandlerContract = contracts.ERC721HandlerAddress
	writer.cfg.genericHandlerContract = contracts.GenericHandlerAddress
	writer.setContract(bridge)

	err = writer.start()
	if err != nil {
		t.Fatal(err)
	}
	return writer
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

func watchEvent(t *testing.T, client *utils.Client, bridge common.Address, subStr utils.EventSig) {
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
		log15.Error("Failed to subscribe to the streaming filter query", "err", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case evt := <-ch:
			fmt.Printf("%s: %#v\n", subStr, evt.Topics)

		case err := <-sub.Err():
			if err != nil {
				fmt.Println(err)
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

func TestCreateAndExecuteErc20DepositProposal(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
	writerA, writerB, errA, errB := createWritersAndClient(t, contracts)

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
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalExecuted)

	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.Erc20AssertBalance(t, client, amount, erc20Address, recipient)
}

func TestCreateAndExecuteErc721Proposal(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
	writerA, writerB, errA, errB := createWritersAndClient(t, contracts)
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
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalExecuted)

	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.Erc721AssertOwner(t, client, erc721Contract, tokenId, recipient)
}

func TestCreateAndExecuteGenericProposal(t *testing.T) {
	client := ethtest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id, AliceKp)
	writerA, writerB, errA, errB := createWritersAndClient(t, contracts)
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
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalCreated)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalVote)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalFinalized)
	go watchEvent(t, client, contracts.BridgeAddress, utils.ProposalExecuted)

	routeMessageAndWait(t, client, writerA, writerB, m, errA, errB)

	ethtest.AssertHashExistence(t, client, hash, assetStoreAddr)
}
