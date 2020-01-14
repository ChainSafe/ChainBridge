package keystore

import (
	"bytes"
	"testing"

	"github.com/status-im/keycard-go/hexutils"
)

func TestMockKeystore(t *testing.T) {
	ks := NewTestKeystore()
	kp, _ := ks.KeypairFromAddress("ethereum")
	if !bytes.Equal(kp.Private().Encode(), hexutils.HexToBytes(TestEthereumPrivateKey)) {
		t.Fatalf("unexpected key. got: %s expected: %s\n", kp.Private(), TestEthereumPrivateKey)
	}
}
