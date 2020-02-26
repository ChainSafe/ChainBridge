// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	centrifuge "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

var testConfig = &Config{
	endpoint: TestEndpoint,
	receiver: TestCentrifugeContractAddress,
	keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
	from:     keystore.AliceKey,
}

var randomHash = []byte{0x12, 0x34}
var sourceChain = msg.ChainId(1)
var depositId uint32 = 0

func createTestReceiverContract(t *testing.T, conn *Connection) ReceiverContract {
	addressBytes := TestReceiverContractAddress.Bytes()

	address := [20]byte{}
	copy(address[:], addressBytes)

	contract, err := receiver.NewReceiver(address, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	instance := &receiver.ReceiverRaw{
		Contract: contract,
	}

	return instance
}

func createTestCentrifugeContract(t *testing.T, conn *Connection) ReceiverContract {
	addressBytes := TestCentrifugeContractAddress.Bytes()

	address := [20]byte{}
	copy(address[:], addressBytes)

	contract, err := centrifuge.NewBridgeAsset(address, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	instance := &centrifuge.BridgeAssetRaw{
		Contract: contract,
	}

	return instance
}

func TestResolveMessage(t *testing.T) {
	m := msg.Message{
		Type:     msg.DepositAssetType,
		Metadata: randomHash,
	}

	conn := NewConnection(testConfig)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	centrifugeContract := createTestCentrifugeContract(t, conn)
	w := NewWriter(conn, testConfig)
	w.SetReceiverContract(centrifugeContract)
	ok := w.ResolveMessage(m)
	if !ok {
		t.Fatal("Transaction failed")
	}
}

func TestWriteToReceiverContract(t *testing.T) {
	cfg := &Config{
		endpoint: TestEndpoint,
		keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	contract := createTestReceiverContract(t, conn)
	auth := createTestAuth(t, conn)

	depositId := big.NewInt(421)
	originChain := big.NewInt(1)

	data := []byte("nootwashere")
	hash := ethcrypto.Keccak256Hash(data)

	_, err = contract.Transact(auth, "createDepositProposal", hash, depositId, originChain)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriter_createDepositProposal(t *testing.T) {
	m := msg.Message{
		Type:      msg.CreateDepositProposalType,
		DepositId: depositId,
		Source:    sourceChain,
	}

	conn := NewConnection(testConfig)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	rcvr := createTestReceiverContract(t, conn)
	w := NewWriter(conn, testConfig)
	w.SetReceiverContract(rcvr)
	ok := w.ResolveMessage(m)
	if !ok {
		t.Fatal("Transaction failed")
	}
}

func TestWriter_voteDepositProposal(t *testing.T) {
	m := msg.Message{
		Type:      msg.VoteDepositProposalType,
		Source:    sourceChain,
		DepositId: 0,
		Metadata:  []byte{1},
	}

	conn := NewConnection(testConfig)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	rcvr := createTestReceiverContract(t, conn)
	w := NewWriter(conn, testConfig)
	w.SetReceiverContract(rcvr)
	ok := w.ResolveMessage(m)
	if !ok {
		t.Fatal("Transaction failed")
	}
}

func TestWriter_executeDeposit(t *testing.T) {
	m := msg.Message{
		Source:    sourceChain,
		Type:      msg.ExecuteDepositType,
		To:        TestAddress.Bytes(),
		DepositId: depositId,
	}

	conn := NewConnection(testConfig)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	rcvr := createTestReceiverContract(t, conn)
	w := NewWriter(conn, testConfig)
	w.SetReceiverContract(rcvr)
	ok := w.ResolveMessage(m)
	if !ok {
		t.Fatal("Transaction failed")
	}
}
