package centrifuge

import (
	"testing"
)

var TestPolkadotEndpoint = "wss://poc3-rpc.polkadot.io"

func TestConnect(t *testing.T) {
	conn := NewConnection(TestPolkadotEndpoint)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}
