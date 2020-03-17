// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

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
const EthChain = "ethereum"
const SubChain = "substrate"

var TestKeyRing *TestKeyRingHolder

//var TestKeyStoreMap map[string]*Keystore

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
		EthereumKeys:   createKeyRing(EthChain),
		CentrifugeKeys: createKeyRing(SubChain),
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
		AliceKey:   createKeypair(AliceKey, chain),
		BobKey:     createKeypair(BobKey, chain),
		CharlieKey: createKeypair(CharlieKey, chain),
		DaveKey:    createKeypair(DaveKey, chain),
		EveKey:     createKeypair(EveKey, chain),
	}

	return ring

}

// createKeypair creates keypairs based on the private key seed inputted for the specfied chain
func createKeypair(key, chain string) crypto.Keypair {
	switch chain {
	case EthChain:
		bz := padWithZeros([]byte(key), secp256k1.PrivateKeyLength)
		return errorWrap(secp256k1.NewKeypairFromPrivateKey(bz)).(*secp256k1.Keypair)
	case SubChain:
		return errorWrap(sr25519.NewKeypairFromSeed("//" + key)).(*sr25519.Keypair)
	}
	return nil

}

// insecureKeypairFromAddress is used for resolving address in an insecure keystore.
// Instead of providing an address a chain reference can be used to fetch a default keypair (eg. "ethereum").
func insecureKeypairFromAddress(keyType string, chainType string) (crypto.Keypair, error) {
	if chainType == EthChain {
		return TestKeyRing.EthereumKeys[keyType], nil
	} else if chainType == SubChain {
		return TestKeyRing.CentrifugeKeys[keyType], nil
	} else {
		fmt.Println(chainType)
		panic("unrecognized key type")
	}
}
