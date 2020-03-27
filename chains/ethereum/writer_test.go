// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

var writerTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
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
