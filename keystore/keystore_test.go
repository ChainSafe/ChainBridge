package keystore

import (
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
)

func TestKeystore(t *testing.T) {
	ks := NewKeystore()

	kp, err := secp256k1.GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	addr := kp.Public()
	ks.Insert(kp)
	kp2 := ks.Get(addr)

	if !reflect.DeepEqual(kp, kp2) {
		t.Fatalf("Fail: got %v expected %v", kp2, kp)
	}
}
