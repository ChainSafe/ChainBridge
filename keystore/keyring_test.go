package keystore

import (
	"bytes"
	"testing"
)

func TestByteLength(t *testing.T) {
	res := makeProperKeyLength([]byte("Alice"), 32)
	if len(res) != 32 {
		t.Errorf("Length is incorrect.")
	}
}

func TestNameKeystore(t *testing.T) {
	tpk := []byte("test")
	ks := createNameKeyStore(tpk)

	tpk32 := makeProperKeyLength(tpk, 32)
	tpk64 := makeProperKeyLength(tpk, 64)

	if !bytes.Equal(ks.SecpKeypair.Private().Encode(), tpk32) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", ks.SecpKeypair.Private(), tpk32)
	}

	if !bytes.Equal(ks.EdKeypair.Private().Encode(), tpk64) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", ks.EdKeypair.Private(), tpk64)
	}
}

func TestTestKeystore(t *testing.T) {
	tps := createTestKeyStore()
	keys := []*NameKeystore{tps.Alice, tps.Bob, tps.Charlie, tps.Dave, tps.Eve}
	for _, key := range keys {
		tpk32 := makeProperKeyLength(key.PrivateKey, 32)
		tpk64 := makeProperKeyLength(key.PrivateKey, 64)

		if !bytes.Equal(key.SecpKeypair.Private().Encode(), tpk32) {
			t.Fatalf("unexpected key. got: %s expected: %s\n", key.SecpKeypair.Private(), tpk32)
		}

		if !bytes.Equal(key.EdKeypair.Private().Encode(), tpk64) {
			t.Fatalf("unexpected key. got: %s expected: %s\n", key.EdKeypair.Private(), tpk64)
		}
	}
}
