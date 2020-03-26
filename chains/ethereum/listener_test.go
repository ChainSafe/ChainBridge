// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

var listenerTestConfig = &Config{
	id:       msg.EthereumId,
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
}

func TestListener_start_stop(t *testing.T) {
	conn := newLocalConnection(t, listenerTestConfig)
	defer conn.Close()

	listener := NewListener(conn, listenerTestConfig)

	err := listener.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = listener.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

// TODO TESTS
// registerEventHandler
// watchEvent <- Must trigger an event
