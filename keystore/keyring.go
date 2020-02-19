package keystore

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/crypto/sr25519"
)

// The Constant "keys". These are the name that the keys are based on. This can be expanded, but
// any additions must be added to TestKeyRing and to insecureKeyFromAddress
const AliceKey = "Alice"
const BobKey = "Bob"
const CharlieKey = "Charlie"
const DaveKey = "Dave"
const EveKey = "Eve"

// The Chain type Constants
const ETHChain = "ethereum"
const CTFGChain = "centrifuge"

var TestKeyRing *TestKeyRingHolder
var TestKeyStoreMap map[string]*Keystore

// TestKeyStore is a struct that holds a Keystore of all the test keys
type TestKeyRingHolder struct {
	EthereumKeys   KeyRing
	CentrifugeKeys KeyRing
}

// KeyRing holds the keypair related to a specfic keypair type
type KeyRing map[string]crypto.Keypair

// Init function to create a keyRing that can be accessed anywhere without having to recreate the data
func init() {
	TestKeyRing = &TestKeyRingHolder{
		EthereumKeys:   createKeyRing(ETHChain),
		CentrifugeKeys: createKeyRing(CTFGChain),
	}

	TestKeyStoreMap = map[string]*Keystore{
		AliceKey:   NewTestKeystore(AliceKey),
		BobKey:     NewTestKeystore(BobKey),
		CharlieKey: NewTestKeystore(CharlieKey),
		DaveKey:    NewTestKeystore(DaveKey),
		EveKey:     NewTestKeystore(EveKey),
	}
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

// createKeyRing creates a KeyRing for the specfied chain/key type
func createKeyRing(chain string) KeyRing {
	ring := map[string]crypto.Keypair{
		AliceKey:   createKeypair([]byte(AliceKey), chain),
		BobKey:     createKeypair([]byte(BobKey), chain),
		CharlieKey: createKeypair([]byte(CharlieKey), chain),
		DaveKey:    createKeypair([]byte(DaveKey), chain),
		EveKey:     createKeypair([]byte(EveKey), chain),
	}

	return ring

}

// createKeypair creates keypairs based on the private key seed inputted for the specfied chain
func createKeypair(key []byte, chain string) crypto.Keypair {
	switch chain {
	case ETHChain:
		secpPrivateKey := errorWrap(secp256k1.NewPrivateKey(padWithZeros(key, secp256k1.PrivateKeyLength))).(*secp256k1.PrivateKey)
		return errorWrap(secp256k1.NewKeypairFromPrivate(secpPrivateKey)).(*secp256k1.Keypair)
	case CTFGChain:
		return errorWrap(sr25519.NewKeypairFromSeed(padWithZeros(key, sr25519.PrivateKeyLength))).(*sr25519.Keypair)
	}
	return nil

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
func (ks *Keystore) insecureKeypairFromAddress(key_type string, chain_type string) (crypto.Keypair, error) {
	if chain_type == ETHChain {
		return TestKeyRing.EthereumKeys[key_type], nil
	} else if chain_type == CTFGChain {
		return TestKeyRing.CentrifugeKeys[key_type], nil
	} else {
		fmt.Println(chain_type)
		panic("unrecognized key type")
	}
	return nil, nil
}
