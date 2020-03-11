// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"bytes"
	"encoding/hex"
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

// A list to show all valid test keys, for testing and documentation purposes
var KeyList = []string{AliceKey, BobKey, CharlieKey, DaveKey, EveKey}

// The Constants for Substrate Keys
const AliceSeed = "e5be9a5092b81bca64be81d212e7f2f9eba183bb7a90954f7b76361f6edb5c0a"
const BobSeed = "398f0c28f98885e046333d4a41c19cee4c37368a9832c6502f6cfd182e2aef89"
const CharlieSeed = "bc1ede780f784bb6991a585e4f6e61522c14e1cae6ad0895fb57b9a205a8f938"
const DaveSeed = "868020ae0687dda7d57565093a69090211449845a7e11453612800b663307246"
const EveSeed = "786ad0e2df456fe43dd1f91ebca22e235bc162e0bb8d53c633e8c85b2af68b7a"

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

func getSeed(key []byte) []byte {
	var seed = make([]byte, 0)
	var err error
	if bytes.Equal(key, []byte(AliceKey)) {
		seed, err = hex.DecodeString(AliceSeed)
	}
	if bytes.Equal(key, []byte(BobKey)) {
		seed, err = hex.DecodeString(BobSeed)
	}
	if bytes.Equal(key, []byte(CharlieKey)) {
		seed, err = hex.DecodeString(CharlieSeed)
	}
	if bytes.Equal(key, []byte(DaveKey)) {
		seed, err = hex.DecodeString(DaveSeed)
	}
	if bytes.Equal(key, []byte(EveKey)) {
		seed, err = hex.DecodeString(EveSeed)
	}

	if err != nil {
		panic(err)
	}

	if len(seed) == 0 {
		panic("Seed Not Found!")
	}

	return seed
}

// createKeypair creates keypairs based on the private key seed inputted for the specfied chain
func createKeypair(key []byte, chain string) crypto.Keypair {
	switch chain {
	case ETHChain:
		secpPrivateKey := errorWrap(secp256k1.NewPrivateKey(padWithZeros(key, secp256k1.PrivateKeyLength))).(*secp256k1.PrivateKey)
		return errorWrap(secp256k1.NewKeypairFromPrivate(secpPrivateKey)).(*secp256k1.Keypair)
	case CTFGChain:
		seed := getSeed(key)
		return errorWrap(sr25519.NewKeypairFromSeed(seed)).(*sr25519.Keypair)
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
}
