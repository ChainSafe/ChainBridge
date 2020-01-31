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

func init() {
	TestKeyRing = &TestKeyStore{
		Alice:   createNamedKeyStore([]byte(AliceKey)),
		Bob:     createNamedKeyStore([]byte(BobKey)),
		Charlie: createNamedKeyStore([]byte(CharlieKey)),
		Dave:    createNamedKeyStore([]byte(DaveKey)),
		Eve:     createNamedKeyStore([]byte(EveKey)),
	}
}

// NamedKeyStore contains the data needed for test keys
type NamedKeyStore struct {
	PrivateKey  []byte
	SecpKeypair crypto.Keypair
	SrKeypair   crypto.Keypair
	EdKeypair   crypto.Keypair
}

func padWithZeros(key []byte, targetLength int) []byte {
	res := make([]byte, targetLength-len(key))
	return append(res, key...)
}

func errorWrap(in interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return in
}

func createNamedKeyStore(key []byte) *NamedKeyStore {
	secpPrivateKey := errorWrap(secp256k1.NewPrivateKey(padWithZeros(key, secp256k1.PrivateKeyLength))).(*secp256k1.PrivateKey)
	edPrivateKey := errorWrap(ed25519.NewPrivateKey(padWithZeros(key, ed25519.PrivateKeyLength))).(*ed25519.PrivateKey)

	return &NamedKeyStore{
		PrivateKey:  key,
		SecpKeypair: errorWrap(secp256k1.NewKeypairFromPrivate(secpPrivateKey)).(*secp256k1.Keypair),
		SrKeypair:   errorWrap(sr25519.NewKeypairFromSeed(padWithZeros(key, sr25519.PrivateKeyLength))).(*sr25519.Keypair),
		EdKeypair:   errorWrap(ed25519.NewKeypairFromPrivate(edPrivateKey)).(*ed25519.Keypair),
	}
}

type TestKeyStore struct {
	Alice   *NamedKeyStore
	Bob     *NamedKeyStore
	Charlie *NamedKeyStore
	Dave    *NamedKeyStore
	Eve     *NamedKeyStore
}

var TestKeyRing *TestKeyStore

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
