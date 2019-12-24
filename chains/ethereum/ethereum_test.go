package ethereum

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

var testPassword = []byte("1234")

func TestInitializeChain(t *testing.T) {
	kp, err := secp256k1.GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	pub := kp.Public().Hex()
	fp, err := filepath.Abs(pub + ".key")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.OpenFile(fp, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()
	defer os.RemoveAll(fp)

	err = keystore.EncryptAndWriteToFile(file, kp.Private(), testPassword)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &core.ChainConfig{
		Id:       msg.EthereumId,
		Endpoint: TestEndpoint,
		Receiver: "",
		Emitter:  "",
		From:     fp,
		Password: string(testPassword),
	}

	chain, err := InitializeChain(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: not yet implemented
	// err = chain.Start()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	id := chain.Id()
	if id != msg.EthereumId {
		t.Fatalf("Fail: got %x expected %x", id, msg.EthereumId)
	}

	// err = chain.Stop()
	// if err != nil {
	// 	t.Fatal(err)
	// }
}
