// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var TestEndpoint = "ws://127.0.0.1:9944"

func TestConnect(t *testing.T) {
	conn := NewConnection(TestEndpoint)
	conn.key = signature.TestKeyringPairAlice
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = conn.SubmitTx(WhitelistChainMethod, types.Bytes([]byte{10}))
	if err != nil {
		t.Fatal(err)
	}
}
