// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var writerTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
}

func writer_deployContracts(t *testing.T) {
	port := "8545"
	numRelayers := 2
	relayerThreshold := big.NewInt(1)
	pk := hexutil.Encode(AliceKp.Encode())[2:]
	DeployedContracts, err := DeployContracts(pk, port, numRelayers, relayerThreshold, uint8(0))
	if err != nil {
		t.Fatal(err)
	}

	writerTestConfig.contract = DeployedContracts.BridgeAddress
}

func TestWriter_start_stop(t *testing.T) {
	conn := newLocalConnection(t, writerTestConfig)
	defer conn.Close()

	writer := NewWriter(conn, writerTestConfig)

	err := writer.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeccak(t *testing.T) {
	args := []string{
		"test",
		"testing",
		"chainsafe",
	}
	expected := []string{
		"0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
		"5f16f4c7f149ac4f9510d9cf8cf384038ad348b3bcdc01915f95de12df9d1b02",
		"699c776c7e6ce8e6d96d979b60e41135a13a2303ae1610c8d546f31f0c6dc730",
	}

	for i, str := range args {
		expectedHash := [32]byte{}
		copy(expectedHash[:], []byte(expected[i]))

		res := keccakHash([]byte(str))

		t.Error(res)
		t.Error(expectedHash)

		if res != expectedHash {
			t.Fatalf("Input: %s, Expected: %s, Output: %s", str, expectedHash, string(res[:]))
		}
	}
}

// TODO TESTS
// Test ResolveMessage
