// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"fmt"
	"os"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
)

const EnvPassword = "KEYSTORE_PASSWORD"

// KeypairFromAddress attempts to load the encrypted key file for the provided address,
// prompting the user for the password.
func KeypairFromAddress(addr, chainType, path string, insecure bool) (crypto.Keypair, error) {
	if insecure {
		return insecureKeypairFromAddress(path, chainType)
	}
	path = fmt.Sprintf("%s/%s.key", path, addr)
	// Make sure key exists before prompting password
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("key file not found: %s", path)
	}

	var pswd []byte
	if pswd := os.Getenv(EnvPassword); pswd == "" {
		GetPassword(fmt.Sprintf("Enter password for key %s:", path))
	}

	kp, err := ReadFromFileAndDecrypt(path, pswd)
	if err != nil {
		return nil, err
	}

	return kp, nil
}
