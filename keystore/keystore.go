package keystore

import (
	"fmt"
	"os"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
)

const KEYSTORE_ENV = "KEYSTORE_PASSWORD"

// Keystore is used to track the location of key files and provide easy retrieval
type Keystore struct {
	path     string // path to keys
	insecure bool   // if true keystore will use predetermined keys, for testing only (see Keystore.insecureKeypairFromAddress())
}

func NewKeystore(path string) *Keystore {
	return &Keystore{
		path: path,
	}
}

// KeypairFromAddress attempts to load the encrypted key file for the provided address,
// prompting the user for the password.
func (ks *Keystore) KeypairFromAddress(addr string, chain_type string) (crypto.Keypair, error) {
	if ks.insecure {
		return ks.insecureKeypairFromAddress(addr, chain_type)
	}
	path := fmt.Sprintf("%s/%s.key", ks.path, addr)
	// Make sure key exists before prompting password
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("key file not found: %s", path)
	}

	var pswd []byte
	if os.Getenv("") != "" {
		pswd = []byte(os.Getenv(KEYSTORE_ENV))
	} else {
		pswd = GetPassword(fmt.Sprintf("Enter password for key %s:", path))
	}

	priv, err := ReadFromFileAndDecrypt(path, pswd)
	if err != nil {
		return nil, err
	}

	return PrivateKeyToKeypair(priv)
}
