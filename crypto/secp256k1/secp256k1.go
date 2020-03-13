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

//const SignatureLength = 64
//const MessageLength = 32

type Keypair struct {
	public  *ecdsa.PublicKey
	private *ecdsa.PrivateKey
}

//type PublicKey struct {
//	key ecdsa.PublicKey
//}
//
//type PrivateKey struct {
//	key ecdsa.PrivateKey
//}

//func NewPrivateKey(in []byte) (*PrivateKey, error) {
//	if len(in) != PrivateKeyLength {
//		return nil, errors.New("input to create secp256k1 private key is not 32 bytes")
//	}
//	priv := new(PrivateKey)
//	err := priv.Decode(in)
//	return priv, err
//}
//
//func NewKeypairFromPrivate(priv *PrivateKey) (*Keypair, error) {
//	pub, err := priv.Public()
//	if err != nil {
//		return nil, err
//	}
//
//	return &Keypair{
//		public:  pub.(*PublicKey),
//		private: priv,
//	}, nil
//}

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

//func (kp *Keypair) Sign(msg []byte) ([]byte, error) {
//	if len(msg) != MessageLength {
//		return nil, errors.New("invalid message length: not 32 byte hash")
//	}
//
//	return secp256k1.Sign(msg, &kp.private.key)
//}

//func (kp *Keypair) Public() crypto.PublicKey {
//	return kp.public
//}
//
//func (kp *Keypair) Private() crypto.PrivateKey {
//	return kp.private
//}

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

//func (k *PublicKey) Verify(msg, sig []byte) (bool, error) {
//	if len(sig) != SignatureLength {
//		return false, errors.New("invalid signature length")
//	}
//
//	if len(msg) != MessageLength {
//		return false, errors.New("invalid message length: not 32 byte hash")
//	}
//
//	return secp256k1.VerifySignature(k.Encode(), msg, sig), nil
//}
//
//func (k *PublicKey) Encode() []byte {
//	return secp256k1.CompressPubkey(&k.key)
//}
//
//func (k *PublicKey) Decode(in []byte) error {
//	pub, err := secp256k1.DecompressPubkey(in)
//	if err != nil {
//		return err
//	}
//	k.key = *pub
//	return nil
//}
//
//func (k *PublicKey) Address() crypto.Address {
//	return secp256k1.PubkeyToAddress(k.key).String()
//}
//
//func (k *PublicKey) Hex() string {
//	enc := k.Encode()
//	h := hex.EncodeToString(enc)
//	return "0x" + h
//}
//
//func (k *PublicKey) Key() ecdsa.PublicKey {
//	return k.key
//}
//
//func (pk *PrivateKey) Sign(msg []byte) ([]byte, error) {
//	if len(msg) != MessageLength {
//		return nil, errors.New("invalid message length: not 32 byte hash")
//	}
//
//	return secp256k1.Sign(msg, &pk.key)
//}
//
//func (pk *PrivateKey) Public() (crypto.PublicKey, error) {
//	kp := NewKeypair(pk.key)
//	return kp.Public(), nil
//}
//
//func (pk *PrivateKey) Encode() []byte {
//	return secp256k1.FromECDSA(&pk.key)
//}
//
//func (pk *PrivateKey) Decode(in []byte) error {
//	key := secp256k1.ToECDSAUnsafe(in)
//	pk.key = *key
//	return nil
//}
//
//func (pk *PrivateKey) Key() *ecdsa.PrivateKey {
//	return &pk.key
//}
