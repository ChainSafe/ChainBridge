// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var TestKeyringPairAlice = signature.KeyringPair{
	URI:       "//Alice",
	PublicKey: []byte{0xd4, 0x35, 0x93, 0xc7, 0x15, 0xfd, 0xd3, 0x1c, 0x61, 0x14, 0x1a, 0xbd, 0x4, 0xa9, 0x9f, 0xd6, 0x82, 0x2c, 0x85, 0x58, 0x85, 0x4c, 0xcd, 0xe3, 0x9a, 0x56, 0x84, 0xe7, 0xa5, 0x6d, 0xa2, 0x7d}, //nolint:lll
	Address:   "5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY",
}

var TestKeyringPairBob = signature.KeyringPair{
	URI:       "//Bob",
	PublicKey: []byte{0x8e, 0xaf, 0x04, 0x15, 0x16, 0x87, 0x73, 0x63, 0x26, 0xc9, 0xfe, 0xa1, 0x7e, 0x25, 0xfc, 0x52, 0x87, 0x61, 0x36, 0x93, 0xc9, 0x12, 0x90, 0x9c, 0xb2, 0x26, 0xaa, 0x47, 0x94, 0xf2, 0x6a, 0x48}, //nolint:lll
	Address:   "5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty",
}

var TestEndpoint = "ws://127.0.0.1:9944"
var Alice = TestKeyringPairAlice
var Bob = TestKeyringPairBob

func TestConnect(t *testing.T) {
	// Create connection with Alice key
	alice := NewConnection(TestEndpoint)
	alice.key = Alice
	err := alice.Connect()

	if err != nil {
		t.Fatal(err)
	}
	defer alice.Close()

	// Create connection with Bob key
	bob := NewConnection(TestEndpoint)
	bob.key = Bob
	err = bob.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer bob.Close()

	// Add Bob as a validator. Currently genesis defautls to only Alice
	err = alice.SubmitTx(AddValidator, types.NewAccountID([]byte(Bob.Address)))
	if err != nil {
		t.Fatal(err)
	}

	proposal, err := types.NewCall(
		alice.meta,
		NativeTransfer.String(),
		types.NewAccountID(Bob.PublicKey),
		types.U32(500_000),
	)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Hash needs deposit ID
	hash, err := types.GetHash(proposal)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Proposal hash: %s\n", hash.Hex())

	err = alice.SubmitTx(CreateProposal, hash, proposal)
	if err != nil {
		t.Fatal(err)
	}

	err = bob.SubmitTx(Vote, hash, types.Bool(true))
	if err != nil {
		t.Fatal(err)
	}
}
