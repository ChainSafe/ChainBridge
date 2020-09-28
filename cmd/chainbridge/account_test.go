// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/crypto"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

var testKeystoreDir = "./test_datadir/"
var testPassword = []byte("1234")

// gethKeystore is a struct representation of the geth keystore file
// used for testing so that we don't need to call geth to test imported ethereum keys
type gethKeystore struct {
	Address string `json:"address"`
	Crypto  struct {
		Cipher       string `json:"cipher"`
		Cipherparams struct {
			Iv string `json:"iv"`
		}
		Ciphertext string `json:"ciphertext"`
		Kdf        string `json:"kdf"`
		Kdfparams  struct {
			Dklen int    `json:"dklen"`
			N     int    `json:"n"`
			P     int    `json:"p"`
			R     int    `json:"r"`
			Salt  string `json:"salt"`
		}
		Mac string `json:"mac"`
	}
	Id      string `json:"id"`
	Version int    `json:"version"`
}

func createTestGethKeystore() gethKeystore {
	gks := gethKeystore{}
	gks.Address = "395a40da74fe28fb3154135ea64007230a68ffb1"

	gks.Crypto.Cipher = "aes-128-ctr"
	gks.Crypto.Ciphertext = "54d1981684fa0a9e2e0b0d3f748884b745363a721381f1be9d64c1689cf762ee"
	gks.Crypto.Cipherparams.Iv = "c05a1f9374d4d8e227a90208d41462b9"
	gks.Crypto.Kdf = "scrypt"
	gks.Crypto.Kdfparams.Dklen = 32
	gks.Crypto.Kdfparams.N = 262144
	gks.Crypto.Kdfparams.P = 1
	gks.Crypto.Kdfparams.R = 8
	gks.Crypto.Kdfparams.Salt = "f4a8c387cc8939b656c92feb705b110725e4ed17f6af9acce367989b526e03f7"
	gks.Crypto.Mac = "dfaf15c18c25b3364e1f96b89207a3b76ed24fac9beea7cbf4937fa9da9d8b1c"

	gks.Id = "deb2f093-c5d8-46d6-9c71-40e09c483667"
	gks.Version = 3

	return gks
}

// newTestContext creates a cli context for a test given a set of flags and values
func newTestContext(description string, flags []string, values []interface{}) (*cli.Context, error) {
	set := flag.NewFlagSet(description, 0)
	for i := range values {
		switch v := values[i].(type) {
		case bool:
			set.Bool(flags[i], v, "")
		case string:
			set.String(flags[i], v, "")
		case uint:
			set.Uint(flags[i], v, "")
		default:
			return nil, fmt.Errorf("unexpected cli value type: %T", values[i])
		}
	}

	ctx := cli.NewContext(nil, set, nil)

	for i := range values {
		switch v := values[i].(type) {
		case bool:
			if v {
				err := ctx.Set(flags[i], "true")
				if err != nil {
					return nil, fmt.Errorf("failed to set cli flag: %T", flags[i])
				}
			} else {
				err := ctx.Set(flags[i], "false")
				if err != nil {
					return nil, fmt.Errorf("failed to set cli flag: %T", flags[i])
				}
			}
		case string:
			err := ctx.Set(flags[i], v)
			if err != nil {
				return nil, fmt.Errorf("failed to set cli flag: %T", flags[i])
			}
		case uint:
			err := ctx.Set(flags[i], strconv.Itoa(int(v)))
			if err != nil {
				return nil, fmt.Errorf("failed to set cli flag: %T", flags[i])
			}
		default:
			return nil, fmt.Errorf("unexpected cli value type: %T", values[i])
		}
	}

	return ctx, nil
}

func TestAccountCommands(t *testing.T) {
	testApp := cli.NewApp()
	testApp.Writer = ioutil.Discard
	testApp.Commands = []*cli.Command{
		&accountCommand,
	}

	keypath := "."
	srFile, err := generateKeypair("sr25519", keypath, testPassword, "substrate")
	if err != nil {
		t.Fatal(err)
	}
	secpFile, err := generateKeypair("secp256k1", keypath, testPassword, "")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(testKeystoreDir)
	defer os.RemoveAll(srFile)
	defer os.RemoveAll(secpFile)

	testcases := []struct {
		description string
		flags       []string
		values      []interface{}
		function    func(*cli.Context, *dataHandler) error
	}{
		{
			"Test chainbridge account generate --secp256k1 --password \"abc\"",
			[]string{"secp256k1", "password"},
			[]interface{}{true, "abc"},
			handleGenerateCmd,
		},
		{
			"Test chainbridge account generate --sr25519 --password \"abc\"",
			[]string{"sr25519", "password"},
			[]interface{}{true, "abc"},
			handleGenerateCmd,
		},
		{
			"Test chainbridge account import --secp256k1 --password \"abc\" --privateKey 000000000000000000000000000000000000000000000000000000416c696365",
			[]string{"secp256k1", "password", "privateKey"},
			[]interface{}{true, "abc", "000000000000000000000000000000000000000000000000000000416c696365"},
			handleImportCmd,
		},
		{
			"Test chainbridge account import --sr25519 --password \"abc\" --privateKey 000000000000000000000000000000000000000000000000000000416c696365",
			[]string{"sr25519", "password", "privateKey"},
			[]interface{}{true, "abc", "0xfa67e5b3c421ac1b9e3d59a74f09fa18f9ad41ec381ba95e5087cb164c03121b"},
			handleImportCmd,
		},
		{
			"Test chainbridge account list",
			[]string{},
			[]interface{}{},
			handleListCmd,
		},
	}

	for _, c := range testcases {
		c := c // bypass scopelint false positive
		t.Run(c.description, func(t *testing.T) {
			ctx, err := newTestContext(c.description, c.flags, c.values)
			require.Nil(t, err)
			keypath := testKeystoreDir
			dh := dataHandler{datadir: keypath}
			err = c.function(ctx, &dh)
			require.Nil(t, err)
		})
	}
}

func TestGetDatadir(t *testing.T) {
	description := "Testing Datadir"
	flags := []string{}
	values := []interface{}{}

	ctx, err := newTestContext(description, flags, values)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getDataDir(ctx)
	if err == nil {
		t.Fatal("No provided datadir should return error")
	}

	flags = []string{"keystore"}
	values = []interface{}{"../.."}
	ctx, err = newTestContext(description, flags, values)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getDataDir(ctx)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGenerateKey_NoType(t *testing.T) {
	keyfile, err := generateKeypair("", testKeystoreDir, testPassword, "")
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

	importkeyfile, err := generateKeypair("sr25519", keypath, testPassword, "")
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

func TestImportEthKey(t *testing.T) {
	gethJSON, err := json.Marshal(createTestGethKeystore())
	if err != nil {
		t.Fatal(err)
	}
	importkeyfile := "../../test.json"
	ioutil.WriteFile(importkeyfile, gethJSON, 0644)

	defer os.RemoveAll(importkeyfile)

	keyfile, err := importEthKey(importkeyfile, testKeystoreDir, []byte{}, []byte{})
	if err != nil {
		t.Fatal(err)
	}

	keys, err := listKeys(testKeystoreDir)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(testKeystoreDir)

	if len(keys) != 1 {
		t.Fatal("Wrong Number of Keys listed")
	}

	if strings.Compare(keys[0], filepath.Base(keyfile)) != 0 {
		t.Fatalf("Fail: got %s expected %s", keys[0], keyfile)
	}
}

func TestImportKey_withPk(t *testing.T) {
	keyfile, err := importPrivKey(cli.NewContext(app, nil, nil), "", testKeystoreDir, "000000000000000000000000000000000000000000000000000000416c696365", testPassword)
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
			keyfile, err = generateKeypair(crypto.Sr25519Type, testKeystoreDir, testPassword, "")
			if err != nil {
				t.Fatal(err)
			}
		} else {
			keyfile, err = generateKeypair(crypto.Secp256k1Type, testKeystoreDir, testPassword, "")
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
