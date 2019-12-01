// secp256k1 wraps the go-ethereum secp256k1 library.
package crypto

import (
	keys "github.com/ChainSafe/gossamer/crypto"
	secp "github.com/ethereum/go-ethereum/crypto"
)

// type Keypair interface {
// 	Sign(msg []byte) ([]byte, error)
// 	Public() PublicKey
// 	Private() PrivateKey
// }

// type PublicKey interface {
// 	Verify(msg, sig []byte) bool
// 	Encode() []byte
// 	Decode([]byte) error
// 	Address() common.Address
// 	Hex() string
// }

// type PrivateKey interface {
// 	Sign(msg []byte) ([]byte, error)
// 	Public() (PublicKey, error)
// 	Encode() []byte
// 	Decode([]byte) error
// }

type SecpKeypair struct {
	public  *SecpPublicKey
	private *SecpPrivateKey
}

type SecpPublicKey secp.PublicKey
type SecpPrivateKey secp.PrivateKey

func NewSecpKeypair(pk SecpPrivateKey) *SecpKeypair {
	pubkey := secp.PublicKey
}

// Keypair methods
func (kp *SecpKeypair) Sign(msg []byte) ([]byte, error) {
	return kp.private.Sign(msg)
}

func (kp *SecpKeypair) Public() keys.PublicKey {}
func (kp *SecpKeypair) Private() keys.PrivateKey {}

// PublicKey methods
func (k *SecpPublicKey) Verify(msg, sig []byte) bool {}
func (k *SecpPublicKey) Encode() []byte {}
func (k *SecpPublicKey) Decode([]byte) error {}
func (k *SecpPublicKey) Address() common.Address {}
func (k *SecpPublicKey) Hex() string {}

// PrivateKey methods
func (pk *SecpPrivateKey) Sign(msg []byte) ([]byte, error) {
	return secp.Sign(msg, pk.private)
}

func (pk *SecpPrivateKey) Public() (PublicKey, error) {}
func (pk *SecpPrivateKey) Sign(msg []byte) ([]byte, error) {}
func (pk *SecpPrivateKey) Public() (PublicKey, error) {}
func (pk *SecpPrivateKey) Encode() []byte {}
func (pk *SecpPrivateKey) Decode([]byte) error {}