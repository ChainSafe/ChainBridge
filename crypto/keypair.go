// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package crypto

type KeyType = string

const Sr25519Type KeyType = "sr25519"
const Secp256k1Type KeyType = "secp256k1"

type Keypair interface {
	// Sign(msg []byte) ([]byte, error)
	//Public() PublicKey
	//Private() PrivateKey
	Encode() []byte
	Decode([]byte) error
	Address() string
	PublicKey() string
	// TODO: Hopefully PublicKey() is sufficient replacement
	// Hex() string
}

//type PublicKey interface {
//	// Verify(msg, sig []byte) (bool, error)
//	// Decode([]byte) error
//
//	Hex() string
//}
//
//type PrivateKey interface {
//	// Sign(msg []byte) ([]byte, error)
//	// Public() (PublicKey, error)
//	//Encode() []byte
//	//Decode([]byte) error
//}

//var ss58Prefix = []byte("SS58PRE")

// PublicKeyToAddress returns an ss58 address given a PublicKey
// see: https://github.com/paritytech/substrate/wiki/External-Address-Format-(SS58)
// also see: https://github.com/paritytech/substrate/blob/master/primitives/core/src/crypto.rs#L275
//func PublicKeyToAddress(pub PublicKey) Address {
//	enc := pub.Encode()
//	hasher, err := blake2b.New(64, nil)
//	if err != nil {
//		return ""
//	}
//	_, err = hasher.Write(append(ss58Prefix, enc...))
//	if err != nil {
//		return ""
//	}
//	checksum := hasher.Sum(nil)
//	return base58.Encode(append(enc, checksum[:2]...))
//}
