// secp256k1 wraps the go-ethereum secp256k1 library.
package crypto

import (
	keys "github.com/ChainSafe/gossamer/crypto"
	secp "github.com/ethereum/go-ethereum/crypto"
)

type SecpKeypair struct {
	public  *SecpPublicKey
	private *SecpPrivateKey
}

type SecpPublicKey secp.PublicKey
type SecpPrivateKey secp.PrivateKey

func NewSecpKeypair(pk SecpPrivateKey) *SecpKeypair {}

// Keypair methods
func (kp *SecpKeypair) Sign(msg []byte) ([]byte, error) {
	return kp.private.Sign(msg)
}

func (kp *SecpKeypair) Public() keys.PublicKey {
	return kp.public
}
func (kp *SecpKeypair) Private() keys.PrivateKey {
	return kp.private
}

// PublicKey methods
func (k *SecpPublicKey) Verify(msg, sig []byte) bool {
	return secp.VerifySignature(k, msg, sig)
}

func (k *SecpPublicKey) Encode() []byte {}
func (k *SecpPublicKey) Decode([]byte) error {}
func (k *SecpPublicKey) Address() common.Address {}
func (k *SecpPublicKey) Hex() string {}

// PrivateKey methods
func (pk *SecpPrivateKey) Sign(msg []byte) ([]byte, error) {
	return secp.Sign(msg, pk.private)
}

func (pk *SecpPrivateKey) Public() (PublicKey, error) {
	kp := NewSecpKeypair(pk)
	return kp.private, nil
}

func (pk *SecpPrivateKey) Encode() []byte {}
func (pk *SecpPrivateKey) Decode([]byte) error {}