// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	bridge "github.com/ChainSafe/ChainBridgeV2/contracts/Bridge"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var methodTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
	gasLimit: big.NewInt(6721975),
	gasPrice: big.NewInt(20000000000),
}

func method_deployContracts(t *testing.T) {
	port := "8545"
	numRelayers := 2
	relayerThreshold := big.NewInt(2)
	pk := hexutil.Encode(AliceKp.Encode())[2:]
	DeployedContracts, err := DeployContracts(pk, port, numRelayers, relayerThreshold, uint8(0))
	if err != nil {
		t.Fatal(err)
	}

	methodTestConfig.contract = DeployedContracts.BridgeAddress
}

func setupWriter(t *testing.T, config *Config) *Writer {
	conn := newLocalConnection(t, config)

	bridgeInstance, err := bridge.NewBridge(methodTestConfig.contract, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	raw := &bridge.BridgeRaw{
		Contract: bridgeInstance,
	}

	bridgeContract := BridgeContract{
		BridgeRaw:    raw,
		BridgeCaller: &bridgeInstance.BridgeCaller,
	}

	writer := NewWriter(conn, methodTestConfig)
	writer.SetBridgeContract(bridgeContract)

	err = writer.Start()
	if err != nil {
		t.Fatal(err)
	}
	return writer
}

func generateMessage() msg.Message {
	return msg.Message{
		Source:       msg.ChainId(0),
		Destination:  msg.ChainId(1),
		DepositNonce: uint32(2),
		To:           common.FromHex(keystore.TestKeyRing.EthereumKeys[keystore.BobKey].PublicKey()),
		Metadata:     []byte("metadata"),
	}
}

func TestWriter_createDepositProposal(t *testing.T) {
	method_deployContracts(t)
	w := setupWriter(t, methodTestConfig)
	m := generateMessage()

	res := w.createDepositProposal(m)
	if res != true {
		t.Fatal("Failed to create deposit proposal")
	}

	// Should fail, cannot make same proposal twice
	res = w.createDepositProposal(m)
	if res != false {
		t.Fatal("Failed to create deposit proposal")
	}
}

func TestWriter_voteDepositProposal(t *testing.T) {
	method_deployContracts(t)
	w := setupWriter(t, methodTestConfig)
	m := generateMessage()

	// Succeeds Vote: 1/2
	createRes := w.createDepositProposal(m)
	if createRes != true {
		t.Fatal("failed to create deposit proposal")
	}

	// Switch signer
	config2 := *methodTestConfig
	config2.from = keystore.BobKey
	w2 := setupWriter(t, &config2)

	// Succeeds Vote: 2/2
	voteRes := w2.voteDepositProposal(m)
	if voteRes != true {
		t.Fatal("Failed to vote")
	}

	// Switch signer
	config3 := *methodTestConfig
	config3.from = keystore.BobKey
	w3 := setupWriter(t, &config3)

	// Vote is finalized already
	voteRes = w3.voteDepositProposal(m)
	if voteRes != false {
		t.Fatal("Failed to vote")
	}
}

func TestWriter_voteDepositProposalFailed(t *testing.T) {
	method_deployContracts(t)
	w := setupWriter(t, methodTestConfig)
	m := generateMessage()

	createRes := w.createDepositProposal(m)
	if createRes != true {
		t.Fatal("failed to create deposit proposal")
	}

	// Proposal with nonce 5 doesn't exist, this should fail
	m.DepositNonce = uint32(5)
	voteRes := w.voteDepositProposal(m)
	if voteRes != false {
		t.Fatal("Vote was supposed to fail, but passed")
	}
}

// Helper function that allows us to create, vote, and execute a proposal
// This should be only be used when testing a handler
func createAndVote(t *testing.T, w *Writer, m msg.Message) {
	createRes := w.createDepositProposal(m)
	if createRes != true {
		t.Fatal("failed to create deposit proposal")
	}

	// Switch signer
	config2 := *methodTestConfig
	config2.from = keystore.BobKey
	w2 := setupWriter(t, &config2)

	voteRes := w2.voteDepositProposal(m)
	if voteRes != true {
		t.Fatal("Failed to vote")
	}

	// Requires that the handler in `m` is actually deployed
	// executeRes := w.executeDeposit(m)
	// if executeRes != true {
	// 	t.Fatal("Failed to execute the proposal")
	// }
}

func Test_createAndVote(t *testing.T) {
	method_deployContracts(t)
	w := setupWriter(t, methodTestConfig)
	m := generateMessage()
	createAndVote(t, w, m)
}
