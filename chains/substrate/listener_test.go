// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"reflect"
	"testing"
	"time"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/utils/substrate"
	subtest "github.com/ChainSafe/ChainBridge/utils/substrate/testing"
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

//func whitelistChain(t *testing.T, c *Connection, id msg.ChainId) {
//	destId := types.U8(id)
//	meta := c.getMetadata()
//	call, err := types.NewCall(
//		&meta,
//		string(utils.WhitelistChain),
//		destId,
//	)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	err = c.SubmitTx(utils.Sudo, call)
//	if err != nil {
//		t.Fatal(err)
//	}
//}

func Test_GenericTransferEvent(t *testing.T) {
	r := &mockRouter{msgs: make(chan msg.Message)}

	ac := createAliceConnection(t)
	alice := NewListener(ac, "Alice", 1, 0, AliceTestLogger)
	alice.SetRouter(r)
	err := alice.Start()
	if err != nil {
		t.Fatal(err)
	}

	// First we have to whitelist the destination chain with sudo
	var destId msg.ChainId = 0
	subtest.WhitelistChain(t, ac, destId)

	// Construct our expected message
	var rId msg.ResourceId
	err = alice.conn.getConst("Example", "HashId", &rId)
	if err != nil {
		t.Fatal(err)
	}
	hashBz := types.MustHexDecodeString("0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2")
	hash := types.NewHash(hashBz)
	expected := msg.NewGenericTransfer(1, 0, 1, rId, hash[:])

	// Initiate transfer
	err = ac.SubmitTx(utils.ExampleTransferHash, hash, destId)
	if err != nil {
		t.Fatal(err)
	}

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
	err = ac.SubmitTx(utils.ExampleTransferHash, hash, destId)
	if err != nil {
		t.Fatal(err)
	}

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
