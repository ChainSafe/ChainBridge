package keystore

import (
	"bytes"
	"sync"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
)

type Keystore struct {
	// map of public keys to keypairs
	keys map[crypto.PublicKey]crypto.Keypair
	lock sync.RWMutex
}

func NewKeystore() *Keystore {
	return &Keystore{
		keys: make(map[crypto.PublicKey]crypto.Keypair),
	}
}

func (ks *Keystore) Insert(kp crypto.Keypair) {
	ks.lock.Lock()
	defer ks.lock.Unlock()
	pub := kp.Public()
	ks.keys[pub] = kp
}

func (ks *Keystore) Get(pub crypto.PublicKey) crypto.Keypair {
	ks.lock.RLock()
	defer ks.lock.RUnlock()
	return ks.keys[pub]
}

//
//func (ks *Keystore) Ed25519PublicKeys() []crypto.PublicKey {
//	edkeys := []crypto.PublicKey{}
//	for _, key := range ks.keys {
//		if _, ok := key.(*ed25519.Keypair); ok {
//			edkeys = append(edkeys, key.Public())
//		}
//	}
//	return edkeys
//}
//
//func (ks *Keystore) Ed25519Keypairs() []crypto.Keypair {
//	edkeys := []crypto.Keypair{}
//	for _, key := range ks.keys {
//		if _, ok := key.(*ed25519.Keypair); ok {
//			edkeys = append(edkeys, key)
//		}
//	}
//	return edkeys
//}
//
//func (ks *Keystore) Sr25519PublicKeys() []crypto.PublicKey {
//	srkeys := []crypto.PublicKey{}
//	for _, key := range ks.keys {
//		if _, ok := key.(*sr25519.Keypair); ok {
//			srkeys = append(srkeys, key.Public())
//		}
//	}
//	return srkeys
//}
//
//func (ks *Keystore) Sr25519Keypairs() []crypto.Keypair {
//	edkeys := []crypto.Keypair{}
//	for _, key := range ks.keys {
//		if _, ok := key.(*sr25519.Keypair); ok {
//			edkeys = append(edkeys, key)
//		}
//	}
//	return edkeys
//}

func (ks *Keystore) GetKeypair(pub crypto.PublicKey) crypto.Keypair {
	for _, key := range ks.keys {
		if bytes.Equal(key.Public().Encode(), pub.Encode()) {
			return key
		}
	}
	return nil
}
