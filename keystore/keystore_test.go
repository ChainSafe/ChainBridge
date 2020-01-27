package keystore

import (
	"bytes"
	"testing"

	"github.com/status-im/keycard-go/hexutils"
)

func TestMockKeystore(t *testing.T) {
	ks := NewTestKeystore("Alice")
	kp, _ := ks.KeypairFromAddress("Alice", "ethereum")
	if !bytes.Equal(kp.Private().Encode(), hexutils.HexToBytes(TestEthereumPrivateKey)) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", kp.Private(), TestEthereumPrivateKey)
	}
}
