package keystore

import (
	"bytes"
	"crypto/rand"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
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

func TestEncryptAndDecryptPrivateKey(t *testing.T) {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	priv, err := secp256k1.NewPrivateKey(buf)
	if err != nil {
		t.Fatal(err)
	}

	password := []byte("noot")

	data, err := EncryptPrivateKey(priv, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := DecryptPrivateKey(data, password, "secp256k1")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(priv, res) {
		t.Fatalf("Fail: got %v expected %v", res, priv)
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
	priv := kp.Private()

	err = EncryptAndWriteToFile(file, priv, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ReadFromFileAndDecrypt(fp, password)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(priv.Encode(), res.Encode()) {
		t.Fatalf("Fail: got %v expected %v", res, priv)
	}
}
