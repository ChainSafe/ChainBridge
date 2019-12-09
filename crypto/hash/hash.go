package hash

import (
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

// Hash used to store a blake2b hash
type Hash [32]byte

// Blake2bHash returns the 256-bit blake2b hash of the input data
func Blake2bHash(in []byte) (Hash, error) {
	h, err := blake2b.New256(nil)
	if err != nil {
		return [32]byte{}, err
	}

	var res []byte
	_, err = h.Write(in)
	if err != nil {
		return [32]byte{}, err
	}

	res = h.Sum(nil)
	var buf = [32]byte{}
	copy(buf[:], res)
	return buf, err
}

// Keccak256 returns the keccak256 hash of the input data
func Keccak256(in []byte) Hash {
	h := sha3.NewLegacyKeccak256()
	hash := h.Sum(in)
	var buf = [32]byte{}
	copy(buf[:], hash)
	return buf
}
