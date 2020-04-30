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
	"strings"
	"testing"

	"github.com/ChainSafe/ChainBridge/crypto"
	"github.com/ChainSafe/ChainBridge/keystore"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

var testKeystoreDir = "./test_datadir/"
var testPassword = []byte("1234")

var GethKeystore struct {
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
			err := ctx.Set(flags[i], string(v))
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

	keypath := "."

	srFile, err := generateKeypair("sr25519", keypath, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	secpFile, err := generateKeypair("secp256k1", keypath, testPassword)
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
		// {
		// 	"Test chainbridge account import --secp256k1 --password \"abc\" " + secpFile,
		// 	[]string{"secp256k1", "password"},
		// 	[]interface{}{true, "abc"},
		// 	handleImportCmd,
		// },
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

func TestWrapHandler(t *testing.T) {

	wrapHandler(handleListCmd)

}

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
