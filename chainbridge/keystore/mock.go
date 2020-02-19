package keystore

import (
	"encoding/hex"
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
)

const TestEthereumPrivateKey = "39a9ea0dce63086c64a80ce045b796bebed2006554e3992d92601515c7b19807"

func NewTestKeystore() *Keystore {
	return &Keystore{
		insecure: true,
	}
}

// insecureKeypairFromAddress is used for resolving address in an insecure keystore.
// Instead of providing an address a chain reference can be used to fetch a default keypair (eg. "ethereum").
func (ks *Keystore) insecureKeypairFromAddress(keyType string) (crypto.Keypair, error) {
	if keyType == "ethereum" {
		privBytes, err := hex.DecodeString(TestEthereumPrivateKey)
		if err != nil {
			panic(err)
		}
		priv, err := secp256k1.NewPrivateKey(privBytes)
		if err != nil {
			panic(err)
		}
		return secp256k1.NewKeypair(*priv.Key()), nil
	} else {
		fmt.Println(keyType)
		panic("unrecognized key type")
	}
}
