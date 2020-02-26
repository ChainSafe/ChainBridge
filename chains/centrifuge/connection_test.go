// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package centrifuge

import (
	"testing"
)

var TestPolkadotEndpoint = "wss://poc3-rpc.polkadot.io"

func TestConnect(t *testing.T) {
	t.Skip()
	// conn := NewConnection(TestPolkadotEndpoint)
	// err := conn.Connect()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// conn.Close()
}
