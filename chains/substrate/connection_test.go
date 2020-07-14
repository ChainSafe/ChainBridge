// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func TestConnect_QueryStorage(t *testing.T) {
	// Create connection with Alice key
	errs := make(chan error)
	conn := NewConnection(TestEndpoint, "Alice", AliceKey, AliceTestLogger, make(chan int), errs)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Query storage
	var data types.AccountInfo
	_, err = conn.queryStorage("System", "Account", conn.key.PublicKey, nil, &data)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure no errors were propagated
	select {
	case err := <-errs:
		t.Fatal(err)
	default:
		return
	}
}

func TestConnect_CheckChainId(t *testing.T) {
	// Create connection with Alice key
	errs := make(chan error)
	conn := NewConnection(TestEndpoint, "Alice", AliceKey, AliceTestLogger, make(chan int), errs)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = conn.checkChainId(ThisChain)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure no errors were propagated
	select {
	case err := <-errs:
		t.Fatal(err)
	default:
		return
	}
}

func TestConnect_SubmitTx(t *testing.T) {
	// Create connection with Alice key
	errs := make(chan error)
	conn := NewConnection(TestEndpoint, "Alice", AliceKey, AliceTestLogger, make(chan int), errs)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Source: https://pkg.go.dev/github.com/centrifuge/go-substrate-rpc-client?tab=doc#example-package-MakeASimpleTransfer
	bob, err := types.NewAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	if err != nil {
		t.Fatal(err)
	}
	err = conn.SubmitTx("Balances.transfer", bob.AsAccountID, types.NewUCompactFromUInt(10))
	if err != nil {
		t.Fatal(err)
	}

	// Ensure no errors were propagated
	select {
	case err := <-errs:
		t.Fatal(err)
	default:
		return
	}
}
