// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"reflect"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/blockstore"
	msg "github.com/ChainSafe/ChainBridge/message"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const ListenerTimeout = time.Second * 30

type mockRouter struct {
	msgs chan msg.Message
}

func (r *mockRouter) Send(message msg.Message) error {
	r.msgs <- message
	return nil
}

func Test_GenericTransferEvent(t *testing.T) {
	testClient := subtest.CreateClient(t, AliceKey, TestEndpoint)
	r := &mockRouter{msgs: make(chan msg.Message)}

	ac := createAliceConnection(t)
	alice := NewListener(ac, "Alice", 1, 0, AliceTestLogger, &blockstore.EmptyStore{})
	alice.SetRouter(r)
	err := alice.Start()
	if err != nil {
		t.Fatal(err)
	}

	// First we have to whitelist the destination chain with sudo
	var destId msg.ChainId = 0
	subtest.WhitelistChain(t, testClient, destId)

	// Construct our expected message
	var rId msg.ResourceId
	subtest.QueryConst(t, testClient, "Example", "HashId", &rId)
	hashBz := types.MustHexDecodeString("0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2")
	hash := types.NewHash(hashBz)
	expected := msg.NewGenericTransfer(1, 0, 1, rId, hash[:])

	// Initiate transfer
	subtest.InitiateHashTransfer(t, testClient, hash, destId)

	// Verify message
	select {
	case m := <-r.msgs:
		if !reflect.DeepEqual(expected, m) {
			t.Fatalf("Unexpected message.\n\tExpected: %#v\n\tGot: %#v\n", expected, m)
		}
	case <-time.After(ListenerTimeout):
		t.Fatalf("test timed out")

	}

	// Repeat the process to assert nonce and hash change

	// Construct our expected message
	hashBz = types.MustHexDecodeString("0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2")
	hash = types.NewHash(hashBz)
	expected = msg.NewGenericTransfer(1, 0, 2, rId, hash[:])

	// Initiate transfer
	subtest.InitiateHashTransfer(t, testClient, hash, destId)

	// Verify message
	select {
	case m := <-r.msgs:
		if !reflect.DeepEqual(expected, m) {
			t.Fatalf("Unexpected message.\n\tExpected: %#v\n\tGot: %#v\n", expected, m)
		}
	case <-time.After(ListenerTimeout):
		t.Fatalf("test timed out")

	}
}
