// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"bytes"
	"reflect"
	"testing"
)

func TestByteLength(t *testing.T) {
	res := padWithZeros([]byte("Alice"), 32)

	exp := append([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, []byte("Alice")...)
	if !bytes.Equal(exp, res) {
		t.Fatalf("Fail. Got: %#v\n\tExpected: %#v\n", res, exp)
	}
}
func TestInsecureAddresses(t *testing.T) {
	tkp, err := insecureKeypairFromAddress(AliceKey, EthChain)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(tkp, TestKeyRing.EthereumKeys[AliceKey]) {
		t.Fatalf("Key is not being returned correctly")
	}

	tkp, err = insecureKeypairFromAddress(AliceKey, SubChain)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(tkp, TestKeyRing.SubstrateKeys[AliceKey]) {
		t.Fatalf("Key is not being returned correctly")
	}

}
