package crypto

type KeyType = string

//const Ed25519Type KeyType = "ed25519"
//const Sr25519Type KeyType = "sr25519"
const Secp256k1Type KeyType = "secp256k1"

type Address = string

type Keypair interface {
	Sign(msg []byte) ([]byte, error)
	Public() PublicKey
	Private() PrivateKey
}

type PublicKey interface {
	Verify(msg, sig []byte) (bool, error)
	Encode() []byte
	Decode([]byte) error
	Address() Address
	Hex() string
}

type PrivateKey interface {
	Sign(msg []byte) ([]byte, error)
	Public() (PublicKey, error)
	Encode() []byte
	Decode([]byte) error
}
