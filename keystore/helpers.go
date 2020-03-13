// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

// PrivateKeyToKeypair returns a public, private keypair given a private key
//func PrivateKeyToKeypair(priv crypto.PrivateKey) (kp crypto.Keypair, err error) {
//	if key, ok := priv.(*secp256k1.PrivateKey); ok {
//		kp, err = secp256k1.NewKeypairFromPrivate(key)
//	} else {
//		return nil, errors.New("cannot decode key: invalid key type")
//	}
//
//	return kp, err
//}
