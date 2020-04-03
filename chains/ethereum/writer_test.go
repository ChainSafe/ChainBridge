// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	msg "github.com/ChainSafe/ChainBridge/message"
)

func createTestWriter(t *testing.T, cfg *Config, contracts *DeployedContracts) (*Connection, *Writer) {
	conn := newLocalConnection(t, cfg)
	writer := NewWriter(conn, cfg, newTestLogger(cfg.name))

	bridge, err := createBridgeContract(contracts.BridgeAddress, conn)
	if err != nil {
		t.Fatal(err)
	}

	erc20Handler, err := createErc20HandlerContract(contracts.ERC20HandlerAddress, conn)
	if err != nil {
		t.Fatal(err)
	}

	writer.conn.cfg.contract = contracts.BridgeAddress
	writer.conn.cfg.erc20HandlerContract = contracts.ERC20HandlerAddress
	writer.conn.cfg.genericHandlerContract = contracts.CentrifugeHandlerAddress
	writer.cfg.contract = contracts.BridgeAddress
	writer.cfg.erc20HandlerContract = contracts.ERC20HandlerAddress
	writer.cfg.genericHandlerContract = contracts.CentrifugeHandlerAddress
	writer.SetContracts(bridge, erc20Handler)

	return conn, writer
}

func TestWriter_start_stop(t *testing.T) {
	conn := newLocalConnection(t, aliceTestConfig)
	defer conn.Close()

	writer := NewWriter(conn, aliceTestConfig, TestLogger)

	err := writer.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestHash(t *testing.T) {
	args := []string{
		"Hello World",
		"testing",
		"chainsafe",
	}
	expected := []string{
		"0x592fa743889fc7f92ac2a37bb1f5ba1daf2a5c84741ca0e0061d243a2e6707ba",
		"0x5f16f4c7f149ac4f9510d9cf8cf384038ad348b3bcdc01915f95de12df9d1b02",
		"0x699c776c7e6ce8e6d96d979b60e41135a13a2303ae1610c8d546f31f0c6dc730",
	}

	for i, str := range args {
		res := hash([]byte(str))

		if expected[i] != common.Hash(res).Hex() {
			t.Fatalf("Input: %s, Expected: %s, Output: %s", str, expected[i], common.Hash(res).String())
		}
	}
}

func watchEvent(conn *Connection, subStr EventSig) {
	fmt.Printf("Watching for event: %s\n", subStr)
	query := buildQuery(conn.cfg.contract, subStr, big.NewInt(0))
	eventSubscription, err := conn.subscribeToEvent(query)
	if err != nil {
		log15.Error("Failed to subscribe to finalization event", "err", err)
	}

	for {
		select {
		case evt := <-eventSubscription.ch:
			fmt.Printf("%s: %#v\n", subStr, evt.Topics)

		case err := <-eventSubscription.sub.Err():
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func verifyHash(t *testing.T, conn *Connection, opts *bind.CallOpts, hash [32]byte) {
	instance, err := centrifugeHandler.NewCentrifugeAssetHandler(conn.cfg.genericHandlerContract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := instance.CentrifugeAssetHandlerCaller.GetHash(opts, hash)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Hash doesn't exist")
	}

}

func TestCreateAndExecuteErc20DepositProposal(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id)

	aliceConn, alice := createTestWriter(t, aliceTestConfig, contracts)
	defer aliceConn.Close()

	bobConn, bob := createTestWriter(t, bobTestConfig, contracts)
	defer bobConn.Close()

	// We'll use alice to setup the erc20
	opts, nonce, err := aliceConn.newTransactOpts(big.NewInt(0), big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice))
	nonce.lock.Unlock() // We manual increment nonce in tests
	if err != nil {
		t.Fatal(err)
	}
	erc20Address := deployMintApproveErc20(t, aliceConn, opts)

	// Create initial transfer message
	tokenId := append(common.LeftPadBytes([]byte{}, 32), common.LeftPadBytes(erc20Address.Bytes(), 32)...)
	recipient := ethcrypto.PubkeyToAddress(bob.conn.kp.PrivateKey().PublicKey).Bytes()
	amount := big.NewInt(10)
	m := msg.NewFungibleTransfer(1, 0, 0, amount, tokenId, recipient)

	// Helpful for debugging
	go watchEvent(alice.conn, DepositProposalCreated)
	go watchEvent(alice.conn, DepositProposalVote)
	go watchEvent(alice.conn, DepositProposalFinalized)
	go watchEvent(alice.conn, DepositProposalExecuted)

	// Watch for executed event
	query := buildQuery(alice.cfg.contract, DepositProposalExecuted, big.NewInt(0))
	eventSubscription, err := alice.conn.subscribeToEvent(query)
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
		case evt := <-eventSubscription.ch:
			sourceId := evt.Topics[1].Big().Uint64()
			destId := evt.Topics[2].Big().Uint64()
			depositNone := evt.Topics[3].Big().Uint64()

			if m.Source == msg.ChainId(sourceId) &&
				m.Destination == msg.ChainId(destId) &&
				m.DepositNonce == uint32(depositNone) {
				return
			}

		case err := <-eventSubscription.sub.Err():
			if err != nil {
				t.Fatal(err)
			}

		case <-time.After(TestTimeout):
			t.Fatal("test timed out")
		}
	}

}

func TestCreateAndExecuteGenericProposal(t *testing.T) {
	contracts := deployTestContracts(t, aliceTestConfig.id)

	aliceConn, alice := createTestWriter(t, aliceTestConfig, contracts)
	defer aliceConn.Close()

	bobConn, bob := createTestWriter(t, bobTestConfig, contracts)
	defer bobConn.Close()

	// Create initial transfer message
	hash := common.HexToHash("0xf0a8748d2b102eb4e0e116047753b9beff0396d81b830693b19a1376ac4b14e8")
	m := msg.Message{
		Source:       1,
		Destination:  0,
		Type:         msg.GenericTransfer,
		DepositNonce: 0,
		Metadata: []interface{}{
			hash.Bytes(),
		},
	}

	// Helpful for debugging
	go watchEvent(alice.conn, DepositProposalCreated)
	go watchEvent(alice.conn, DepositProposalVote)
	go watchEvent(alice.conn, DepositProposalFinalized)
	go watchEvent(alice.conn, DepositProposalExecuted)

	// Watch for executed event
	query := buildQuery(alice.cfg.contract, DepositProposalExecuted, big.NewInt(0))
	eventSubscription, err := alice.conn.subscribeToEvent(query)
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

	func() {
		for {
			select {
			case evt := <-eventSubscription.ch:
				sourceId := evt.Topics[1].Big().Uint64()
				destId := evt.Topics[2].Big().Uint64()
				depositNonce := evt.Topics[3].Big().Uint64()

				if m.Source == msg.ChainId(sourceId) &&
					m.Destination == msg.ChainId(destId) &&
					m.DepositNonce == uint32(depositNonce) {
					return
				}

			case err := <-eventSubscription.sub.Err():
				if err != nil {
					t.Fatal(err)
				}

			case <-time.After(TestTimeout):
				t.Fatal("test timed out")
			}
		}
	}()

	verifyHash(t, aliceConn, &bind.CallOpts{}, hash)
}
