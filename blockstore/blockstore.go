package blockstore

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	msg "github.com/ChainSafe/ChainBridge/message"
)

// TryLoadLatestBlock will attempt to load the latest block for the chain/relayer pair, returning 0 if not found.
func TryLoadLatestBlock(store string, chain msg.ChainId, relayer string) (*big.Int, error) {
	fileName := getFileName(chain, relayer)

	if store == "" {
		path, err := getDefaultPath()
		if err != nil {
			return nil, err
		}
		store = path
	}
	path := store + "/" + fileName
	// If it exists, load and return
	exists, err := fileExists(path)
	if err != nil {
		return nil, err
	}
	if exists{
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		block, _ := big.NewInt(0).SetString(string(dat), 10)
		return block, nil
	}
	// Otherwise just return
	return big.NewInt(0), nil
}

// Save latest block will write the block number to disk for the chain/relayer pair.
func SaveLatestBlock(store string, chain msg.ChainId, relayer string, block *big.Int) error {
	fileName := getFileName(chain, relayer)

	if store == "" {
		path, err := getDefaultPath()
		if err != nil {
			return err
		}
		store = path
	}
	path := store + "/" + fileName
	// Create dir if it does not exist
	if _, err := os.Stat(store); os.IsNotExist(err) {
		_ = os.Mkdir(store, os.ModePerm)
	}

	// Write bytes to file
	data := []byte(block.String())
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getFileName(chain msg.ChainId, relayer string) string {
	return fmt.Sprintf("%s-%d.block", relayer, chain)
}

func getDefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/blockstore", nil
}

func fileExists(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}