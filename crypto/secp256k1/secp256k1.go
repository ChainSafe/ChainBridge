// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package secp256k1

import (
	"crypto/ecdsa"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"

	secp256k1 "github.com/ethereum/go-ethereum/crypto"
)

var _ crypto.Keypair = &Keypair{}

const PrivateKeyLength = 32

type Keypair struct {
	public  *ecdsa.PublicKey
	private *ecdsa.PrivateKey
}

func NewKeypairFromPrivateKey(priv []byte) (*Keypair, error) {
	pk, err := secp256k1.ToECDSA(priv)
	if err != nil {
		return nil, err
	}

	return &Keypair{
		public:  pk.Public().(*ecdsa.PublicKey),
		private: pk,
	}, nil
}

// NewKeypairFromPrivateKey parses a string for a hex private key. Must be at least
// PrivateKeyLength long.
func NewKeypairFromString(priv string) (*Keypair, error) {
	pk, err := secp256k1.HexToECDSA(priv)
	if err != nil {
		return nil, err
	}

	return &Keypair{
		public:  pk.Public().(*ecdsa.PublicKey),
		private: pk,
	}, nil
}

func NewKeypair(pk ecdsa.PrivateKey) *Keypair {
	pub := pk.Public()

	return &Keypair{
		public:  pub.(*ecdsa.PublicKey),
		private: &pk,
	}
}

func GenerateKeypair() (*Keypair, error) {
	priv, err := secp256k1.GenerateKey()
	if err != nil {
		return nil, err
	}

	return NewKeypair(*priv), nil
}

func (kp *Keypair) Encode() []byte {
	return secp256k1.FromECDSA(kp.private)
}

func (kp *Keypair) Decode(in []byte) error {
	key, err := secp256k1.ToECDSA(in)
	if err != nil {
		return err
	}

	kp.public = key.Public().(*ecdsa.PublicKey)
	kp.private = key

	return nil
}

func (kp *Keypair) Address() string {
	return secp256k1.PubkeyToAddress(*kp.public).String()
}

func (kp *Keypair) PublicKey() string {
	return hexutil.Encode(secp256k1.CompressPubkey(kp.public))
}

func (kp *Keypair) PrivateKey() *ecdsa.PrivateKey {
	return kp.private
}
