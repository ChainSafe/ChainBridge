package keystore

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
)

const TestEthereumPrivateKey = "39a9ea0dce63086c64a80ce045b796bebed2006554e3992d92601515c7b19807"

func NewTestKeystore(devmode string) *Keystore {
	return &Keystore{
		path:     devmode,
		insecure: true,
	}
}

// insecureKeypairFromAddress is used for resolving address in an insecure keystore.
// Instead of providing an address a chain reference can be used to fetch a default keypair (eg. "ethereum").
func (ks *Keystore) insecureKeypairFromAddress(keyType string, chain_type string) (crypto.Keypair, error) {
	if chain_type == "ethereum" {
		testKeyring := createTestKeyStore()
		switch keyType {
		case "Alice":
			return testKeyring.Alice.SecpKeypair, nil
		case "Bob":
			return testKeyring.Bob.SecpKeypair, nil
		case "Charlie":
			return testKeyring.Charlie.SecpKeypair, nil
		case "Dave":
			return testKeyring.Dave.SecpKeypair, nil
		case "Eve":
			return testKeyring.Eve.SecpKeypair, nil
		}
	} else {
		fmt.Println(chain_type)
		panic("unrecognized key type")
	}
	return nil, nil
}
