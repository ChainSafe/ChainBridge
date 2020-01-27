package keystore

import (
	"github.com/ChainSafe/ChainBridgeV2/crypto/ed25519"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/crypto/sr25519"
)

type NameKeystore struct {
	PrivateKey  []byte
	SecpKeypair *secp256k1.Keypair
	SrKeypair   *sr25519.Keypair
	EdKeypair   *ed25519.Keypair
}

func makeProperKeyLength(key []byte, targetLength int) []byte {
	res := make([]byte, targetLength-len(key))
	return append(res, key...)
}

func errorWrap(in interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return in
}

func createNameKeyStore(key []byte) *NameKeystore {
	secpPrivateKey := errorWrap(secp256k1.NewPrivateKey(makeProperKeyLength(key, secp256k1.PrivateKeyLength))).(*secp256k1.PrivateKey)
	edPrivateKey := errorWrap(ed25519.NewPrivateKey(makeProperKeyLength(key, ed25519.PrivateKeyLength))).(*ed25519.PrivateKey)

	return &NameKeystore{
		PrivateKey:  key,
		SecpKeypair: errorWrap(secp256k1.NewKeypairFromPrivate(secpPrivateKey)).(*secp256k1.Keypair),
		SrKeypair:   errorWrap(sr25519.NewKeypairFromSeed(makeProperKeyLength(key, sr25519.PrivateKeyLength))).(*sr25519.Keypair),
		EdKeypair:   errorWrap(ed25519.NewKeypairFromPrivate(edPrivateKey)).(*ed25519.Keypair),
	}
}

type TestKeyStore struct {
	Alice   *NameKeystore
	Bob     *NameKeystore
	Charlie *NameKeystore
	Dave    *NameKeystore
	Eve     *NameKeystore
}

func createTestKeyStore() *TestKeyStore {
	return &TestKeyStore{
		Alice:   createNameKeyStore([]byte("Alice")),
		Bob:     createNameKeyStore([]byte("Bob")),
		Charlie: createNameKeyStore([]byte("Charlie")),
		Dave:    createNameKeyStore([]byte("Dave")),
		Eve:     createNameKeyStore([]byte("Eve")),
	}
}
