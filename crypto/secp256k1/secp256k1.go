package secp256k1

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"ChainBridgeV2/crypto"

	"ChainBridgeV2/crypto/hash"

	secp256k1 "github.com/ethereum/go-ethereum/crypto"
)

const SignatureLength = 64

type Keypair struct {
	public  *PublicKey
	private *PrivateKey
}

type PublicKey struct {
	key ecdsa.PublicKey
}

type PrivateKey struct {
	key ecdsa.PrivateKey
}

func NewKeypair(pk ecdsa.PrivateKey) *Keypair {
	pub := pk.Public()

	return &Keypair{
		public:  &PublicKey{key: *pub.(*ecdsa.PublicKey)},
		private: &PrivateKey{key: pk},
	}
}

func GenerateKeypair() (*Keypair, error) {
	priv, err := secp256k1.GenerateKey()
	if err != nil {
		return nil, err
	}

	return NewKeypair(*priv), nil
}

func (kp *Keypair) Sign(msg []byte) ([]byte, error) {
	hash, err := hash.Blake2bHash(msg)
	if err != nil {
		return nil, err
	}
	return secp256k1.Sign(hash[:], &kp.private.key)
}

func (kp *Keypair) Public() crypto.PublicKey {
	return kp.public
}
func (kp *Keypair) Private() crypto.PrivateKey {
	return kp.private
}

func (k *PublicKey) Verify(msg, sig []byte) (bool, error) {
	if len(sig) != SignatureLength {
		return false, errors.New("invalid signature length")
	}

	hash, err := hash.Blake2bHash(msg)
	if err != nil {
		return false, errors.New("could not hash message")
	}
	return secp256k1.VerifySignature(k.Encode(), hash[:], sig), nil
}

func (k *PublicKey) Encode() []byte {
	return secp256k1.CompressPubkey(&k.key)
}

func (k *PublicKey) Decode(in []byte) error {
	pub, err := secp256k1.DecompressPubkey(in)
	if err != nil {
		return err
	}
	k.key = *pub
	return nil
}

func (k *PublicKey) Address() crypto.Address {
	return secp256k1.PubkeyToAddress(k.key).String()
}

func (k *PublicKey) Hex() string {
	enc := k.Encode()
	h := hex.EncodeToString(enc)
	return "0x" + h
}

func (pk *PrivateKey) Sign(msg []byte) ([]byte, error) {
	h, err := hash.Blake2bHash(msg)
	if err != nil {
		return nil, err
	}
	return secp256k1.Sign(h[:], &pk.key)
}

func (pk *PrivateKey) Public() (crypto.PublicKey, error) {
	kp := NewKeypair(pk.key)
	return kp.Public(), nil
}

func (pk *PrivateKey) Encode() []byte {
	return secp256k1.FromECDSA(&pk.key)
}

func (pk *PrivateKey) Decode(in []byte) error {
	key := secp256k1.ToECDSAUnsafe(in)
	pk.key = *key
	return nil
}
