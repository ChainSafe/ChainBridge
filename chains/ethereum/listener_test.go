// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"testing"

	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
)

var listenerTestConfig = &Config{
	id:       msg.ChainId(0),
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
}

func TestListener_start_stop(t *testing.T) {
	conn := newLocalConnection(t, listenerTestConfig)
	defer conn.Close()

	listener := NewListener(conn, listenerTestConfig, TestLogger)

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
