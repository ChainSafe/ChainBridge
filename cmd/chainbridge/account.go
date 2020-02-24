package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/ed25519"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/crypto/sr25519"
	"github.com/ChainSafe/ChainBridgeV2/keystore"

	log "github.com/ChainSafe/log15"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// handleAccountsCmd manages the flags for the account subcommand
// first, if the generate flag is set, if so, it generates a new keypair
// then, if the import flag is set, if so, it imports a keypair
// finally, if the list flag is set, it lists all the keys in the keystore
func handleAccountsCmd(ctx *cli.Context) error {
	err := startLogger(ctx)
	if err != nil {
		return err
	}

	// key directory is datadir/keystore/
	datadir, err := getDataDir(ctx)
	if err != nil {
		return fmt.Errorf("failed to access the datadir: %s", err)
	}

	// import key
	if keyimport := ctx.String(ImportFlag.Name); keyimport != "" {
		log.Info("Importing key...")
		_, err = importKey(keyimport, datadir)
		if err != nil {
			return fmt.Errorf("failed to import key: %s", err)
		}
	}

	// list keys
	if keylist := ctx.Bool(ListFlag.Name); keylist {
		_, err = listKeys(datadir)
		if err != nil {
			return fmt.Errorf("failed to list keys: %s", err)
		}
	}

	return nil
}

func handleGenerateCmd(ctx *cli.Context) error {
	err := startLogger(ctx)
	if err != nil {
		return err
	}

	// key directory is datadir/keystore/
	datadir, err := getDataDir(ctx)
	if err != nil {
		return fmt.Errorf("failed to access the datadir: %s", err)
	}

	log.Info("Generating keypair...")

	// check if --ed25519 or --sr25519 is set
	keytype := crypto.Secp256k1Type
	if flagtype := ctx.Bool(Sr25519Flag.Name); flagtype {
		keytype = crypto.Sr25519Type
	} else if flagtype := ctx.Bool(Ed25519Flag.Name); flagtype {
		keytype = crypto.Ed25519Type
	} else if flagtype := ctx.Bool(Secp256k1Flag.Name); flagtype {
		keytype = crypto.Secp256k1Type
	}

	// check if --password is set
	var password []byte = nil
	if pwdflag := ctx.String(PasswordFlag.Name); pwdflag != "" {
		password = []byte(pwdflag)
	}

	var privKey string = ""
	if privkeyflag := ctx.String(PrivateKeyFlag.Name); privkeyflag != "" {
		// Hex must not have leading 0x
		if privkeyflag[0:2] == "0x" {
			privKey = privkeyflag[2:]
		} else {
			privKey = privkeyflag
		}
	}

	_, err = generateKeypair(keytype, datadir, password, privKey)
	if err != nil {
		return fmt.Errorf("failed to generate key: %s", err)
	}
	return nil
}

func getDataDir(ctx *cli.Context) (string, error) {
	// key directory is datadir/keystore/
	if dir := ctx.GlobalString(KeystorePathFlag.Name); dir != "" {
		datadir, err := filepath.Abs(dir)
		if err != nil {
			return "", err
		}
		log.Trace(fmt.Sprintf("Using keystore dir: %s", datadir))
		return datadir, nil
	}
	return "", fmt.Errorf("datadir flag not supplied")
}

// importKey imports a key specified by its filename to datadir/keystore/
// it saves it under the filename "[publickey].key"
// it returns the absolute path of the imported key file
func importKey(filename, datadir string) (string, error) {
	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return "", fmt.Errorf("could not get keystore directory: %s", err)
	}

	importdata, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", fmt.Errorf("could not read import file: %s", err)
	}

	ksjson := new(keystore.EncryptedKeystore)
	err = json.Unmarshal(importdata, ksjson)
	if err != nil {
		return "", fmt.Errorf("could not read file contents: %s", err)
	}

	keystorefile, err := filepath.Abs(keystorepath + "/" + ksjson.PublicKey[2:] + ".key")
	if err != nil {
		return "", fmt.Errorf("could not create keystore file path: %s", err)
	}

	err = ioutil.WriteFile(keystorefile, importdata, 0644)
	if err != nil {
		return "", fmt.Errorf("could not write to keystore directory: %s", err)
	}

	log.Info("successfully imported key", "public key", ksjson.PublicKey, "file", keystorefile)
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
		return nil, fmt.Errorf("could not get keystore directory: %s", err)
	}

	files, err := ioutil.ReadDir(keystorepath)
	if err != nil {
		return nil, fmt.Errorf("could not read keystore dir: %s", err)
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
func generateKeypair(keytype, datadir string, password []byte, privateKey string) (string, error) {
	if password == nil {
		password = getPassword("Enter password to encrypt keystore file:")
	}

	if keytype == "" {
		log.Info("Using default key type", "type", keytype)
		keytype = crypto.Secp256k1Type
	}

	var kp crypto.Keypair
	var err error

	if keytype == crypto.Sr25519Type {
		// generate sr25519 keys
		kp, err = sr25519.GenerateKeypair()
		if err != nil {
			return "", fmt.Errorf("could not generate sr25519 keypair: %s", err)
		}
	} else if keytype == crypto.Ed25519Type {
		// generate ed25519 keys
		kp, err = ed25519.GenerateKeypair()
		if err != nil {
			return "", fmt.Errorf("could not generate ed25519 keypair: %s", err)
		}
	} else if keytype == crypto.Secp256k1Type {
		// generate secp256k1 keys
		if privateKey != "" {
			kp, err = secp256k1.NewKeypairFromString(privateKey)
			if err != nil {
				return "", fmt.Errorf("could not generate secp256k1 keypair from given string: %s", err)
			}
		} else {
			kp, err = secp256k1.GenerateKeypair()
			if err != nil {
				return "", fmt.Errorf("could not generate secp256k1 keypair: %s", err)
			}
		}
	}

	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return "", fmt.Errorf("could not get keystore directory: %s", err)
	}

	pub := hex.EncodeToString(kp.Public().Encode())
	fp, err := filepath.Abs(keystorepath + "/" + pub + ".key")
	if err != nil {
		return "", fmt.Errorf("invalid filepath: %s", err)
	}

	file, err := os.OpenFile(fp, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Error("generate keypair: could not close keystore file")
		}
	}()

	err = keystore.EncryptAndWriteToFile(file, kp.Private(), password)
	if err != nil {
		return "", fmt.Errorf("could not write key to file: %s", err)
	}

	log.Info("key generated", "public key", pub, "type", keytype, "file", fp)
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
		keyPath = DefaultKeystorePath

		keystorepath, err = filepath.Abs(keyPath)
		if err != nil {
			return "", fmt.Errorf("could not create keystore file path: %s", err)
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

// prompt user to enter password for encrypted keystore
func getPassword(msg string) []byte {
	for {
		fmt.Println(msg)
		fmt.Print("> ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("invalid input: %s\n", err)
		} else {
			fmt.Printf("\n")
			return password
		}
	}
}
