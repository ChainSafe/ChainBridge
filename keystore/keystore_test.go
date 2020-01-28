package keystore

import (
	"bytes"
	"testing"
)

func TestMockKeystore(t *testing.T) {
	ks := NewTestKeystore("Alice")
	tks := createTestKeyStore()
	kp, _ := ks.KeypairFromAddress("Alice", "ethereum")
	if !bytes.Equal(kp.Private().Encode(), tks.Alice.SecpKeypair.Private().Encode()) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", kp.Private(), TestEthereumPrivateKey)
	}
}
