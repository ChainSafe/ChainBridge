// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package crypto

type KeyType = string

const Sr25519Type KeyType = "sr25519"
const Secp256k1Type KeyType = "secp256k1"

type Keypair interface {
	Encode() []byte
	Decode([]byte) error
	Address() string
	PublicKey() string
}
