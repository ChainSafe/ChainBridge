// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package router

import (
	"reflect"
	"testing"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type mockWriter struct {
	msgs []msg.Message
}

func (w *mockWriter) Start() error { return nil }
func (w *mockWriter) Stop() error  { return nil }

func (w *mockWriter) ResolveMessage(msg msg.Message) {
	w.msgs = append(w.msgs, msg)
}

func TestRouter(t *testing.T) {
	router := NewRouter()

	ethW := &mockWriter{msgs: *new([]msg.Message)}
	err := router.Listen(msg.EthereumId, ethW)
	if err != nil {
		t.Fatal(err)
	}

	ctfgW := &mockWriter{msgs: *new([]msg.Message)}
	err = router.Listen(msg.CentrifugeId, ctfgW)
	if err != nil {
		t.Fatal(err)
	}

	msgEthToCtfg := msg.Message{
		Source:      msg.EthereumId,
		Destination: msg.CentrifugeId,
	}

	msgCtfgToEth := msg.Message{
		Source:      msg.CentrifugeId,
		Destination: msg.EthereumId,
	}

	err = router.Send(msgCtfgToEth)
	if err != nil {
		t.Fatal(err)
	}
	err = router.Send(msgEthToCtfg)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(ethW.msgs[0], msgCtfgToEth) {
		t.Error("Unexpected message")
	}

	if !reflect.DeepEqual(ctfgW.msgs[0], msgEthToCtfg) {
		t.Error("Unexpected message")
	}
}
