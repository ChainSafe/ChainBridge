// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var methodTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
}

func method_deployContracts(t *testing.T) {
	port := "8545"
	numRelayers := 2
	relayerThreshold := big.NewInt(1)
	pk := hexutil.Encode(AliceKp.Encode())[2:]
	DeployedContracts, err := DeployContracts(pk, port, numRelayers, relayerThreshold, uint8(0))
	if err != nil {
		t.Fatal(err)
	}

	methodTestConfig.contract = DeployedContracts.BridgeAddress
}

func setupWriter(t *testing.T) *Writer {
	conn := newLocalConnection(t, methodTestConfig)
	defer conn.Close()

	writer := NewWriter(conn, methodTestConfig)

	err := writer.Start()
	if err != nil {
		t.Fatal(err)
	}
	return writer
}

func generateMessage() msg.Message {
	return msg.Message{
		Source:       msg.ChainId(0),
		Destination:  msg.ChainId(1),
		DepositNonce: uint32(1),
		To:           common.FromHex(keystore.TestKeyRing.EthereumKeys[keystore.BobKey].PublicKey()),
		Metadata:     []byte("metadata"),
	}
}

// func TestWriter_createDepositProposal(t *testing.T) {
// 	method_deployContracts(t)
// 	w := setupWriter(t)
// 	m := generateMessage()

// 	res := w.createDepositProposal(m)
// 	if res != true {
// 		t.Fatal("Failed to create deposit proposal")
// 	}
// }
