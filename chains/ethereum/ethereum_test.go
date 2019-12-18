package ethereum

import (
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

func TestInitializeChain(t *testing.T) {
	kp, err := secp256k1.GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}

	cfg := &core.ChainConfig{
		Id:       msg.EthereumId,
		Endpoint: TestEthereumEndpoint,
		Home:     "",
		Away:     "",
		From:     string(kp.Public().Encode()),
	}

	chain, err := InitializeChain(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = chain.Start()
	if err != nil {
		t.Fatal(err)
	}

	id := chain.Id()
	if id != msg.EthereumId {
		t.Fatalf("Fail: got %x expected %x", id, msg.EthereumId)
	}

	err = chain.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
