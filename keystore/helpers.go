// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"errors"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
)

// PrivateKeyToKeypair returns a public, private keypair given a private key
func PrivateKeyToKeypair(priv crypto.PrivateKey) (kp crypto.Keypair, err error) {
	if key, ok := priv.(*secp256k1.PrivateKey); ok {
		kp, err = secp256k1.NewKeypairFromPrivate(key)
	} else {
		return nil, errors.New("cannot decode key: invalid key type")
	}

	return kp, err
}

// DecodePrivateKey turns input bytes into a private key based on the specified key type
func DecodePrivateKey(in []byte, keytype crypto.KeyType) (priv crypto.PrivateKey, err error) {
	if keytype == crypto.Secp256k1Type {
		priv, err = secp256k1.NewPrivateKey(in)
	} else {
		return nil, errors.New("cannot decode key: invalid key type")
	}

	return priv, err
}
