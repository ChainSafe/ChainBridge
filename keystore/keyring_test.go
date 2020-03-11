// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"bytes"
	"testing"
)

func TestByteLength(t *testing.T) {
	res := padWithZeros([]byte("Alice"), 32)
	if len(res) != 32 {
		t.Errorf("Length is incorrect.")
	}
}
func TestInsecureAddresses(t *testing.T) {
	for _, key := range KeyList {
		tks := NewTestKeystore(key)
		tkp, err := tks.insecureKeypairFromAddress(key, ETHChain)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(tkp.Private().Encode(), TestKeyRing.EthereumKeys[key].Private().Encode()) {
			t.Fatalf("Key is not being returned correctly")
		}
		if !bytes.Equal(tkp.Public().Encode(), TestKeyRing.EthereumKeys[key].Public().Encode()) {
			t.Fatalf("Key is not being returned correctly")
		}

		tkp, err = tks.insecureKeypairFromAddress(key, CTFGChain)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(tkp.Private().Encode(), TestKeyRing.CentrifugeKeys[key].Private().Encode()) {
			t.Fatalf("Key is not being returned correctly")
		}
		if !bytes.Equal(tkp.Public().Encode(), TestKeyRing.CentrifugeKeys[key].Public().Encode()) {
			t.Fatalf("Key is not being returned correctly")
		}
	}
}
