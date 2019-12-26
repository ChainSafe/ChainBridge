package keystore

import (
	"fmt"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
)

type Keystore struct {
	// map of public keys to keypairs
	//keys map[crypto.PublicKey]crypto.Keypair
	//lock sync.RWMutex
	path     string
	insecure bool
}

func NewKeystore(path string) *Keystore {
	return &Keystore{
		path: path,
	}
}

// KeypairFromAddress attempts to load the encrypted key file for the provided address,
// prompting the user for the password.
func (ks *Keystore) KeypairFromAddress(addr string) (crypto.Keypair, error) {
	if ks.insecure {
		return ks.insecureKeypairFromAddress(addr)
	}
	path := fmt.Sprintf("%s/%s.key", ks.path, addr)
	//if _, err := os.Stat(path); os.IsNotExist(err) {
	//	return nil, fmt.Errorf("key file not found: %s", path)
	//}
	pswd := common.GetPassword(fmt.Sprintf("Enter password for key %s:", path))
	priv, err := ReadFromFileAndDecrypt(path, pswd)
	if err != nil {
		return nil, err
	}

	return PrivateKeyToKeypair(priv)
}
