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
	tkp, err := insecureKeypairFromAddress(AliceKey, EthChain)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(tkp.Private().Encode(), TestKeyRing.EthereumKeys[AliceKey].Private().Encode()) {
		t.Fatalf("Key is not being returned correctly")
	}

	tkp, err = insecureKeypairFromAddress(AliceKey, SubChain)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(tkp.Private().Encode(), TestKeyRing.CentrifugeKeys[AliceKey].Private().Encode()) {
		t.Fatalf("Key is not being returned correctly")
	}

}
