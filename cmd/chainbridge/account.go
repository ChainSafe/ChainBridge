// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/chainbridge-utils/crypto"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/crypto/sr25519"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	log "github.com/ChainSafe/log15"
	gokeystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/urfave/cli/v2"
)

//dataHandler is a struct which wraps any extra data our CMD functions need that cannot be passed through parameters
type dataHandler struct {
	datadir string
}

// wrapHandler takes in a Cmd function (all declared below) and wraps
// it in the correct signature for the Cli Commands
func wrapHandler(hdl func(*cli.Context, *dataHandler) error) cli.ActionFunc {

	return func(ctx *cli.Context) error {
		err := startLogger(ctx)
		if err != nil {
			return err
		}

		datadir, err := getDataDir(ctx)
		if err != nil {
			return fmt.Errorf("failed to access the datadir: %w", err)
		}

		return hdl(ctx, &dataHandler{datadir: datadir})
	}
}

// handleGenerateCmd generates a keystore for the accounts
func handleGenerateCmd(ctx *cli.Context, dHandler *dataHandler) error {

	log.Info("Generating keypair...")

	// check if --ed25519 or --sr25519 is set
	keytype := crypto.Secp256k1Type
	if flagtype := ctx.Bool(config.Sr25519Flag.Name); flagtype {
		keytype = crypto.Sr25519Type
	} else if flagtype := ctx.Bool(config.Secp256k1Flag.Name); flagtype {
		keytype = crypto.Secp256k1Type
	}

	// check if --password is set
	var password []byte = nil
	if pwdflag := ctx.String(config.PasswordFlag.Name); pwdflag != "" {
		password = []byte(pwdflag)
	}

	_, err := generateKeypair(keytype, dHandler.datadir, password, ctx.String(config.SubkeyNetworkFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}
	return nil
}

// handleImportCmd imports external keystores into the bridge
func handleImportCmd(ctx *cli.Context, dHandler *dataHandler) error {
	log.Info("Importing key...")
	var err error

	// check if --ed25519 or --sr25519 is set
	keytype := crypto.Secp256k1Type
	if flagtype := ctx.Bool(config.Sr25519Flag.Name); flagtype {
		keytype = crypto.Sr25519Type
	} else if flagtype := ctx.Bool(config.Secp256k1Flag.Name); flagtype {
		keytype = crypto.Secp256k1Type
	}

	if ctx.Bool(config.EthereumImportFlag.Name) {
		if keyimport := ctx.Args().First(); keyimport != "" {
			// check if --password is set
			var password []byte = nil
			if pwdflag := ctx.String(config.PasswordFlag.Name); pwdflag != "" {
				password = []byte(pwdflag)
			}
			_, err = importEthKey(keyimport, dHandler.datadir, password, nil)
		} else {
			return fmt.Errorf("Must provide a key to import.")
		}
	} else if privkeyflag := ctx.String(config.PrivateKeyFlag.Name); privkeyflag != "" {
		// check if --password is set
		var password []byte = nil
		if pwdflag := ctx.String(config.PasswordFlag.Name); pwdflag != "" {
			password = []byte(pwdflag)
		}

		_, err = importPrivKey(ctx, keytype, dHandler.datadir, privkeyflag, password)
	} else {
		if keyimport := ctx.Args().First(); keyimport != "" {
			_, err = importKey(keyimport, dHandler.datadir)
		} else {
			return fmt.Errorf("Must provide a key to import.")
		}
	}

	if err != nil {
		return fmt.Errorf("failed to import key: %w", err)
	}

	return nil
}

// handleListCmd lists all accounts currently in the bridge
func handleListCmd(ctx *cli.Context, dHandler *dataHandler) error {

	_, err := listKeys(dHandler.datadir)
	if err != nil {
		return fmt.Errorf("failed to list keys: %w", err)
	}

	return nil
}

// getDataDir obtains the path to the keystore and returns it as a string
func getDataDir(ctx *cli.Context) (string, error) {
	// key directory is datadir/keystore/
	if dir := ctx.String(config.KeystorePathFlag.Name); dir != "" {
		datadir, err := filepath.Abs(dir)
		if err != nil {
			return "", err
		}
		log.Trace(fmt.Sprintf("Using keystore dir: %s", datadir))
		return datadir, nil
	}
	return "", fmt.Errorf("datadir flag not supplied")
}

//importPrivKey imports a private key into a keypair
func importPrivKey(ctx *cli.Context, keytype, datadir, key string, password []byte) (string, error) {
	if password == nil {
		password = keystore.GetPassword("Enter password to encrypt keystore file:")
	}
	keystorepath, err := keystoreDir(datadir)

	if keytype == "" {
		log.Info("Using default key type", "type", keytype)
		keytype = crypto.Secp256k1Type
	}

	var kp crypto.Keypair

	if keytype == crypto.Sr25519Type {
		// generate sr25519 keys
		network := ctx.String(config.SubkeyNetworkFlag.Name)
		kp, err = sr25519.NewKeypairFromSeed(key, network)
		if err != nil {
			return "", fmt.Errorf("could not generate sr25519 keypair from given string: %w", err)
		}
	} else if keytype == crypto.Secp256k1Type {
		// Hex must not have leading 0x
		if key[0:2] == "0x" {
			kp, err = secp256k1.NewKeypairFromString(key[2:])
		} else {
			kp, err = secp256k1.NewKeypairFromString(key)
		}

		if err != nil {
			return "", fmt.Errorf("could not generate secp256k1 keypair from given string: %w", err)
		}
	} else {
		return "", fmt.Errorf("invalid key type: %s", keytype)
	}

	fp, err := filepath.Abs(keystorepath + "/" + kp.Address() + ".key")
	if err != nil {
		return "", fmt.Errorf("invalid filepath: %w", err)
	}

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", fmt.Errorf("Unable to Open File: %w", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Error("import private key: could not close keystore file")
		}
	}()

	err = keystore.EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		return "", fmt.Errorf("could not write key to file: %w", err)
	}

	log.Info("private key imported", "address", kp.Address(), "file", fp)
	return fp, nil

}

//importEthKey takes an ethereum keystore and converts it to our keystore format
func importEthKey(filename, datadir string, password, newPassword []byte) (string, error) {
	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return "", fmt.Errorf("could not get keystore directory: %w", err)
	}

	importdata, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", fmt.Errorf("could not read import file: %w", err)
	}

	if password == nil {
		password = keystore.GetPassword("Enter password to decrypt keystore file:")
	}

	key, err := gokeystore.DecryptKey(importdata, string(password))
	if err != nil {
		return "", fmt.Errorf("Unable to decrypt file: %w", err)
	}

	kp := secp256k1.NewKeypair(*key.PrivateKey)

	fp, err := filepath.Abs(keystorepath + "/" + kp.Address() + ".key")
	if err != nil {
		return "", fmt.Errorf("invalid filepath: %w", err)
	}

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Error("generate keypair: could not close keystore file")
		}
	}()

	if newPassword == nil {
		newPassword = keystore.GetPassword("Enter password to encrypt new keystore file:")
	}

	err = keystore.EncryptAndWriteToFile(file, kp, newPassword)
	if err != nil {
		return "", fmt.Errorf("could not write key to file: %w", err)
	}

	log.Info("ETH key imported", "address", kp.Address(), "file", fp)
	return fp, nil

}

// importKey imports a key specified by its filename to datadir/keystore/
// it saves it under the filename "[publickey].key"
// it returns the absolute path of the imported key file
func importKey(filename, datadir string) (string, error) {
	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return "", fmt.Errorf("could not get keystore directory: %w", err)
	}

	importdata, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", fmt.Errorf("could not read import file: %w", err)
	}

	ksjson := new(keystore.EncryptedKeystore)
	err = json.Unmarshal(importdata, ksjson)
	if err != nil {
		return "", fmt.Errorf("could not read file contents: %w", err)
	}

	keystorefile, err := filepath.Abs(keystorepath + "/" + ksjson.Address[2:] + ".key")
	if err != nil {
		return "", fmt.Errorf("could not create keystore file path: %w", err)
	}

	err = ioutil.WriteFile(keystorefile, importdata, 0600)
	if err != nil {
		return "", fmt.Errorf("could not write to keystore directory: %w", err)
	}

	log.Info("successfully imported key", "address", ksjson.Address, "file", keystorefile)
	return keystorefile, nil
}

// listKeys lists all the keys in the datadir/keystore/ directory and returns them as a list of filepaths
func listKeys(datadir string) ([]string, error) {
	keys, err := getKeyFiles(datadir)
	if err != nil {
		return nil, err
	}

	fmt.Printf("=== Found %d keys ===\n", len(keys))
	for i, key := range keys {
		fmt.Printf("[%d] %s\n", i, key)
	}

	return keys, nil
}

// getKeyFiles returns the filenames of all the keys in the datadir's keystore
func getKeyFiles(datadir string) ([]string, error) {
	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return nil, fmt.Errorf("could not get keystore directory: %w", err)
	}

	files, err := ioutil.ReadDir(keystorepath)
	if err != nil {
		return nil, fmt.Errorf("could not read keystore dir: %w", err)
	}

	keys := []string{}

	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if ext == ".key" {
			keys = append(keys, f.Name())
		}
	}

	return keys, nil
}

// generateKeypair create a new keypair with the corresponding type and saves it to datadir/keystore/[public key].key
// in json format encrypted using the specified password
// it returns the resulting filepath of the new key
func generateKeypair(keytype, datadir string, password []byte, subNetwork string) (string, error) {
	if password == nil {
		password = keystore.GetPassword("Enter password to encrypt keystore file:")
	}

	if keytype == "" {
		log.Info("Using default key type", "type", keytype)
		keytype = crypto.Secp256k1Type
	}

	var kp crypto.Keypair
	var err error

	if keytype == crypto.Sr25519Type {
		// generate sr25519 keys
		kp, err = sr25519.GenerateKeypair(subNetwork)
		if err != nil {
			return "", fmt.Errorf("could not generate sr25519 keypair: %w", err)
		}
	} else if keytype == crypto.Secp256k1Type {
		// generate secp256k1 keys
		kp, err = secp256k1.GenerateKeypair()
		if err != nil {
			return "", fmt.Errorf("could not generate secp256k1 keypair: %w", err)
		}
	} else {
		return "", fmt.Errorf("invalid key type: %s", keytype)
	}

	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return "", fmt.Errorf("could not get keystore directory: %w", err)
	}

	fp, err := filepath.Abs(keystorepath + "/" + kp.Address() + ".key")
	if err != nil {
		return "", fmt.Errorf("invalid filepath: %w", err)
	}

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Error("generate keypair: could not close keystore file")
		}
	}()

	err = keystore.EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		return "", fmt.Errorf("could not write key to file: %w", err)
	}

	log.Info("key generated", "address", kp.Address(), "type", keytype, "file", fp)
	return fp, nil
}

// keystoreDir returnns the absolute filepath of the keystore directory given a datadir
// by default, it is ./keys/
// otherwise, it is datadir/keys/
func keystoreDir(keyPath string) (keystorepath string, err error) {
	// datadir specified, return datadir/keys as absolute path
	if keyPath != "" {
		keystorepath, err = filepath.Abs(keyPath)
		if err != nil {
			return "", err
		}
	} else {
		// datadir not specified, use default
		keyPath = config.DefaultKeystorePath

		keystorepath, err = filepath.Abs(keyPath)
		if err != nil {
			return "", fmt.Errorf("could not create keystore file path: %w", err)
		}
	}

	// if datadir does not exist, create it
	if _, err = os.Stat(keyPath); os.IsNotExist(err) {
		err = os.Mkdir(keyPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// if datadir/keystore does not exist, create it
	if _, err = os.Stat(keystorepath); os.IsNotExist(err) {
		err = os.Mkdir(keystorepath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return keystorepath, nil
}
