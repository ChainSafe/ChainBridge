// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"bytes"
	"crypto/rand"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridge/crypto/sr25519"
)

func TestEncryptAndDecrypt(t *testing.T) {
	password := []byte("noot")
	msg := []byte("helloworld")

	ciphertext, err := Encrypt(msg, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := Decrypt(ciphertext, password)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(msg, res) {
		t.Fatalf("Fail to decrypt: got %x expected %x", res, msg)
	}
}

func TestEncryptAndDecryptKeypair(t *testing.T) {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	kp, err := secp256k1.NewKeypairFromPrivateKey(buf)
	if err != nil {
		t.Fatal(err)
	}

	password := []byte("noot")

	data, err := EncryptKeypair(kp, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := DecryptKeypair(kp.PublicKey(), data, password, "secp256k1")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(kp, res) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}

func createTestFile(t *testing.T) (*os.File, string) {
	filename := "./test_key"

	fp, err := filepath.Abs(filename)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create(fp)
	if err != nil {
		t.Fatal(err)
	}

	return file, fp
}

func TestEncryptAndDecryptFromFile_Secp256k1(t *testing.T) {
	password := []byte("noot")
	file, fp := createTestFile(t)
	defer os.Remove(fp)

	kp, err := secp256k1.GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	err = EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ReadFromFileAndDecrypt(fp, password, "secp256k1")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(kp.Encode(), res.Encode()) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}

func TestEncryptAndDecryptFromFile_Sr25519(t *testing.T) {
	password := []byte("ansermino")
	file, fp := createTestFile(t)
	defer os.Remove(fp)

	kp, err := sr25519.NewKeypairFromSeed("//seed", "substrate")
	if err != nil {
		t.Fatal(err)
	}

	err = EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ReadFromFileAndDecrypt(fp, password, "sr25519")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(kp.Encode(), res.Encode()) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}

func TestDecryptIncorrectType(t *testing.T) {
	password := []byte("ansermino")
	file, fp := createTestFile(t)
	defer os.Remove(fp)

	kp, err := sr25519.NewKeypairFromSeed("//seed", "substrate")
	if err != nil {
		t.Fatal(err)
	}

	err = EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ReadFromFileAndDecrypt(fp, password, "secp256k1")
	if err == nil {
		t.Fatal("Expected mismatch error, got none.")
	}
}
