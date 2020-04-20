// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/blockstore"
	msg "github.com/ChainSafe/ChainBridge/message"
	subutils "github.com/ChainSafe/ChainBridge/shared/substrate"
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

func newTestListener(t *testing.T) (*listener, *mockRouter, *subutils.Client) {
	testClient := subtest.CreateClient(t, AliceKey, TestEndpoint)
	r := &mockRouter{msgs: make(chan msg.Message)}

	ac := createAliceConnection(t)
	alice := NewListener(ac, "Alice", 1, 0, AliceTestLogger, &blockstore.EmptyStore{})
	alice.setRouter(r)
	err := alice.start()
	if err != nil {
		t.Fatal(err)
	}

	return alice, r, testClient
}

func verifyResultingMessage(t *testing.T, r *mockRouter, expected msg.Message) {
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

func Test_FungibleTransferEvent(t *testing.T) {
	_, r, testClient := newTestListener(t)

	// First we have to whitelist the destination chain with sudo
	var destId msg.ChainId = 0
	subtest.WhitelistChain(t, testClient, destId)

	// Construct our expected message
	var rId msg.ResourceId
	subtest.QueryConst(t, testClient, "Example", "NativeTokenId", &rId)
	amount := big.NewInt(1000000)
	recipient := BobKey.PublicKey
	expected := msg.NewFungibleTransfer(1, 0, 1, amount, rId, recipient)

	subtest.InitiateNativeTransfer(t, testClient, types.NewU32(uint32(amount.Int64())), recipient, 0)

	verifyResultingMessage(t, r, expected)
}

func Test_NonFungibleTransferEvent(t *testing.T) {
	_, r, testClient := newTestListener(t)

	// First we have to whitelist the destination chain with sudo
	var destId msg.ChainId = 0
	subtest.WhitelistChain(t, testClient, destId)

	// Construct our expected message
	var rId msg.ResourceId
	subtest.QueryConst(t, testClient, "Example", "NFTTokenId", &rId)
	recipient := BobKey.PublicKey
	tokenId := big.NewInt(99)
	metadata := big.NewInt(0x8080808).Bytes()
	expected := msg.NewNonFungibleTransfer(1, 0, 1, rId, tokenId.Bytes(), recipient, metadata)

	subtest.InitiateNonFungibleTransfer(t, testClient, types.NewU256(*tokenId), recipient, 0)

	verifyResultingMessage(t, r, expected)
}

func Test_GenericTransferEvent(t *testing.T) {
	_, r, testClient := newTestListener(t)

	// First we have to whitelist the destination chain with sudo
	var destId msg.ChainId = 0
	subtest.WhitelistChain(t, testClient, destId)

	// Construct our expected message
	var rId msg.ResourceId
	subtest.QueryConst(t, testClient, "Example", "HashId", &rId)
	hashBz := types.MustHexDecodeString("0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2")
	hash := types.NewHash(hashBz)
	expected := msg.NewGenericTransfer(1, 0, 1, rId, hash[:])

	subtest.InitiateHashTransfer(t, testClient, hash, destId)

	verifyResultingMessage(t, r, expected)

	// Repeat the process to assert nonce and hash change

	// Construct our expected message
	hashBz = types.MustHexDecodeString("0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2")
	hash = types.NewHash(hashBz)
	expected = msg.NewGenericTransfer(1, 0, 2, rId, hash[:])

	subtest.InitiateHashTransfer(t, testClient, hash, destId)

	verifyResultingMessage(t, r, expected)
}
