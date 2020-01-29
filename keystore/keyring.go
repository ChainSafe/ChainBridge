package keystore

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/ed25519"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/crypto/sr25519"
)

const AliceKey = "Alice"
const BobKey = "Bob"
const CharlieKey = "Charlie"
const DaveKey = "Dave"
const EveKey = "Eve"

type NameKeystore struct {
	PrivateKey  []byte
	SecpKeypair crypto.Keypair
	SrKeypair   crypto.Keypair
	EdKeypair   crypto.Keypair
}

func makeProperKeyLength(key []byte, targetLength int) []byte {
	res := make([]byte, targetLength-len(key))
	return append(res, key...)
}

func errorWrap(in interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return in
}

func createNameKeyStore(key []byte) *NameKeystore {
	secpPrivateKey := errorWrap(secp256k1.NewPrivateKey(makeProperKeyLength(key, secp256k1.PrivateKeyLength))).(*secp256k1.PrivateKey)
	edPrivateKey := errorWrap(ed25519.NewPrivateKey(makeProperKeyLength(key, ed25519.PrivateKeyLength))).(*ed25519.PrivateKey)

	return &NameKeystore{
		PrivateKey:  key,
		SecpKeypair: errorWrap(secp256k1.NewKeypairFromPrivate(secpPrivateKey)).(*secp256k1.Keypair),
		SrKeypair:   errorWrap(sr25519.NewKeypairFromSeed(makeProperKeyLength(key, sr25519.PrivateKeyLength))).(*sr25519.Keypair),
		EdKeypair:   errorWrap(ed25519.NewKeypairFromPrivate(edPrivateKey)).(*ed25519.Keypair),
	}
}

type TestKeyStore struct {
	Alice   *NameKeystore
	Bob     *NameKeystore
	Charlie *NameKeystore
	Dave    *NameKeystore
	Eve     *NameKeystore
}

var TestKeyRing *TestKeyStore

func init() {
	TestKeyRing = &TestKeyStore{
		Alice:   createNameKeyStore([]byte(AliceKey)),
		Bob:     createNameKeyStore([]byte(BobKey)),
		Charlie: createNameKeyStore([]byte(CharlieKey)),
		Dave:    createNameKeyStore([]byte(DaveKey)),
		Eve:     createNameKeyStore([]byte(EveKey)),
	}
}

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
		switch keyType {
		case "Alice":
			return TestKeyRing.Alice.SecpKeypair, nil
		case "Bob":
			return TestKeyRing.Bob.SecpKeypair, nil
		case "Charlie":
			return TestKeyRing.Charlie.SecpKeypair, nil
		case "Dave":
			return TestKeyRing.Dave.SecpKeypair, nil
		case "Eve":
			return TestKeyRing.Eve.SecpKeypair, nil
		}
	} else {
		fmt.Println(chain_type)
		panic("unrecognized key type")
	}
	return nil, nil
}
