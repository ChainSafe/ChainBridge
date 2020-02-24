// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"bytes"
	"testing"
)

func TestMockKeystore(t *testing.T) {
	ks := NewTestKeystore(AliceKey)
	kp, _ := ks.KeypairFromAddress(AliceKey, "ethereum")
	if !bytes.Equal(kp.Private().Encode(), TestKeyRing.EthereumKeys[AliceKey].Private().Encode()) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", kp.Private(), TestKeyRing.EthereumKeys[AliceKey].Private().Encode())
	}
}
