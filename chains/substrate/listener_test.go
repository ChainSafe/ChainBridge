// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"reflect"
	"testing"
	"time"

	msg "github.com/ChainSafe/ChainBridge/message"
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

func whitelistChain(c *Connection, id uint32) error {
	destId := types.U32(id)
	meta := c.getMetadata()
	call, err := types.NewCall(
		&meta,
		WhitelistChain.String(),
		destId,
	)
	if err != nil {
		return err
	}

	err = c.SubmitTx(Sudo, call)
	if err != nil {
		return err
	}
	return nil
}

func constructMessage(t *testing.T, hash string, depositNonce uint32, destId uint32) (msg.Message, types.Hash) {
	assetHashBytes := types.Bytes(types.MustHexDecodeString(hash))
	assetHash, err := types.NewHashFromHexString(hash)
	if err != nil {
		t.Fatal(err)
	}

	return msg.Message{
		Source:       1,
		Destination:  0,
		Type:         msg.GenericTransfer,
		DepositNonce: depositNonce,
		Metadata:     []interface{}{assetHashBytes},
	}, assetHash
}

func Test_AssetTx(t *testing.T) {
	ac := createAliceConnection(t)
	r := &mockRouter{msgs: make(chan msg.Message)}
	alice := NewListener(ac, "Alice", 1, 0, TestLogger)
	alice.SetRouter(r)
	err := alice.Start()
	if err != nil {
		t.Fatal(err)
	}

	// First we have to whitelist the destination chain with sudo
	var destId uint32 = 0
	err = whitelistChain(ac, destId)
	if err != nil {
		t.Fatal(err)
	}

	// Construct our expected message
	hashStr := "0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2"
	expected, hash := constructMessage(t, hashStr, 0, 1)

	// Initiate transfer
	recipient := types.Bytes([]byte{0xAB, 0xCD})
	err = ac.SubmitTx(ExampleTransferHash, hash, recipient, destId)
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
	hashStr = "0x5d57194171ec6cd04032c42dde05d73f1c9366c16b64071a551c657c1db4c8f9"
	expected, hash = constructMessage(t, hashStr, 1, 1)

	// Initiate transfer
	err = ac.SubmitTx(ExampleTransferHash, hash, recipient, destId)
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
