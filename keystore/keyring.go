package keystore

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/ed25519"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/crypto/sr25519"
)

// The Constant "keys". These are the name that the keys are based on. This can be expanded, but
// any additions but be added to TestKeyRing and to insecureKeyFromAddress
const AliceKey = "Alice"
const BobKey = "Bob"
const CharlieKey = "Charlie"
const DaveKey = "Dave"
const EveKey = "Eve"

var TestKeyRing *TestKeyStore
var TestKeyStoreMap map[string]*Keystore

// Init function to create a keyRing that can be accessed anywhere without having to recreate the data
func init() {
	TestKeyRing = &TestKeyStore{
		Alice:   createNamedKeyStore([]byte(AliceKey)),
		Bob:     createNamedKeyStore([]byte(BobKey)),
		Charlie: createNamedKeyStore([]byte(CharlieKey)),
		Dave:    createNamedKeyStore([]byte(DaveKey)),
		Eve:     createNamedKeyStore([]byte(EveKey)),
	}

	TestKeyStoreMap = map[string]*Keystore{
		AliceKey:   NewTestKeystore(AliceKey),
		BobKey:     NewTestKeystore(BobKey),
		CharlieKey: NewTestKeystore(CharlieKey),
		DaveKey:    NewTestKeystore(DaveKey),
		EveKey:     NewTestKeystore(EveKey),
	}
}

// NamedKeyStore contains the data needed for test keys
type NamedKeyStore struct {
	PrivateKey  []byte         // The named PrivateKey (eg Alice, Bob)
	SecpKeypair crypto.Keypair //The required keypair for Secp
	SrKeypair   crypto.Keypair //The required keypair for Sr
	EdKeypair   crypto.Keypair //The required keypair for Ed
}

// padWithZeros adds on extra 0 bytes to make a byte array of a specified length
func padWithZeros(key []byte, targetLength int) []byte {
	res := make([]byte, targetLength-len(key))
	return append(res, key...)
}

// errorWrap is a helper function that panics on errors, to make the code cleaner
func errorWrap(in interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return in
}

// createdNamedKeyStore creates private keys for all 3 protocols based off a given private key,
// then makes keypair based on those new private keys
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

// TestKeyStore is a struct that holds a Keystore of all the test keys
type TestKeyStore struct {
	Alice   *NamedKeyStore
	Bob     *NamedKeyStore
	Charlie *NamedKeyStore
	Dave    *NamedKeyStore
	Eve     *NamedKeyStore
}

//NewTestKeystore creates an insecure keystores for testing purposes
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
