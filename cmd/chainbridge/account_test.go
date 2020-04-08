// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/ChainSafe/ChainBridge/crypto"
	"github.com/ChainSafe/ChainBridge/keystore"
)

var testKeystoreDir = "./test_datadir/"
var testPassword = []byte("1234")

func TestGenerateKey_NoType(t *testing.T) {
	keyfile, err := generateKeypair("", testKeystoreDir, testPassword)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(testKeystoreDir)

	contents, err := ioutil.ReadFile(keyfile)
	if err != nil {
		t.Fatal(err)
	}

	kscontents := new(keystore.EncryptedKeystore)
	err = json.Unmarshal(contents, kscontents)
	if err != nil {
		t.Fatal(err)
	}

	if kscontents.Type != "secp256k1" {
		t.Fatalf("Fail: got %s expected %s", kscontents.Type, "secp256k1")
	}
}

func TestImportKey_ShouldFail(t *testing.T) {
	_, err := importKey("./notakey.key", testKeystoreDir)
	if err == nil {
		t.Fatal("did not err")
	}
}

func TestImportKey(t *testing.T) {
	keypath := "../../"

	importkeyfile, err := generateKeypair("sr25519", keypath, testPassword)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(importkeyfile)

	keyfile, err := importKey(importkeyfile, testKeystoreDir)
	if err != nil {
		t.Fatal(err)
	}

	keys, err := listKeys(testKeystoreDir)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(testKeystoreDir)

	if len(keys) != 1 {
		t.Fatal("fail")
	}

	if strings.Compare(keys[0], filepath.Base(keyfile)) != 0 {
		t.Fatalf("Fail: got %s expected %s", keys[0], keyfile)
	}
}

func TestImportKey_withPk(t *testing.T) {
	keyfile, err := importPrivKey("", testKeystoreDir, "000000000000000000000000000000000000000000000000000000416c696365", testPassword)
	if err != nil {
		t.Fatal(err)
	}

	keys, err := listKeys(testKeystoreDir)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(testKeystoreDir)

	if len(keys) != 1 {
		t.Fatal("fail")
	}

	if strings.Compare(keys[0], filepath.Base(keyfile)) != 0 {
		t.Fatalf("Fail: got %s expected %s", keys[0], keyfile)
	}
}

func TestListKeys(t *testing.T) {
	expected := []string{}
	defer os.RemoveAll(testKeystoreDir)

	for i := 0; i < 5; i++ {
		var err error
		var keyfile string
		if i%2 == 0 {
			keyfile, err = generateKeypair(crypto.Sr25519Type, testKeystoreDir, testPassword)
			if err != nil {
				t.Fatal(err)
			}
		} else {
			keyfile, err = generateKeypair(crypto.Secp256k1Type, testKeystoreDir, testPassword)
			if err != nil {
				t.Fatal(err)
			}
		}

		expected = append(expected, keyfile)
	}

	keys, err := listKeys(testKeystoreDir)
	if err != nil {
		t.Fatal(err)
	}

	if len(keys) != len(expected) {
		t.Fatalf("Fail: expected %d keys in keystore, got %d", len(expected), len(keys))
	}

	sort.Slice(expected, func(i, j int) bool { return strings.Compare(expected[i], expected[j]) < 0 })

	for i, key := range keys {
		if strings.Compare(key, filepath.Base(expected[i])) != 0 {
			t.Fatalf("Fail: got %s expected %s", key, filepath.Base(expected[i]))
		}
	}
}
