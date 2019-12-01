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

// Keypair methods
func (kp *SecpKeypair) Sign(msg []byte) ([]byte, error) {
	return kp.private.Sign(msg)
}

func (kp *SecpKeypair) Public() keys.PublicKey {}

// PublicKey methods

// PrivateKey methods
func (pk *SecpPrivateKey) Sign(msg []byte) ([]byte, error) {
	return secp.Sign(msg, pk.private)
}
