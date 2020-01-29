package keystore

import (
	"bytes"
	"testing"
)

func TestMockKeystore(t *testing.T) {
	ks := NewTestKeystore(AliceKey)
	kp, _ := ks.KeypairFromAddress(AliceKey, "ethereum")
	if !bytes.Equal(kp.Private().Encode(), TestKeyRing.Alice.SecpKeypair.Private().Encode()) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", kp.Private(), TestKeyRing.Alice.SecpKeypair.Private().Encode())
	}
}
