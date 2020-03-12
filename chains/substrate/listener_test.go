package substrate

import (
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/core"
)

func TestListener(t *testing.T) {
	alice := NewConnection(TestEndpoint, &Alice)
	err := alice.Connect()
	if err != nil {
		t.Fatal(err)
	}

	cfg := core.ChainConfig{
		Id:       0,
		Endpoint: "",
		From:     "",
		Keystore: nil,
		Opts:     nil,
	}

	l := NewListener(alice, &cfg)

	err = l.Start()
	if err != nil {
		t.Fatal(err)
	}

	select {}
}
