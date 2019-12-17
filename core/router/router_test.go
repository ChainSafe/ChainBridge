package router

import (
	"reflect"
	"testing"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

func TestRouter(t *testing.T) {
	router := NewRouter()

	ethCh := make(chan msg.Message)
	err := router.Listen(msg.EthereumId, ethCh)
	if err != nil {
		t.Fatal(err)
	}

	ctfgCh := make(chan msg.Message)
	err = router.Listen(msg.CentrifugeId, ctfgCh)
	if err != nil {
		t.Fatal(err)
	}

	msgEthToCtfg := msg.Message{
		Source:      msg.EthereumId,
		Destination: msg.CentrifugeId,
	}

	msgCtfgToEth := msg.Message{
		Source:      msg.CentrifugeId,
		Destination: msg.EthereumId,
	}

	err = router.Send(msgCtfgToEth)
	if err != nil {
		t.Fatal(err)
	}
	err = router.Send(msgEthToCtfg)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(<- ethCh, msgCtfgToEth) {
		t.Error("Unexpected message")
	}

	if !reflect.DeepEqual(<- ctfgCh, msgEthToCtfg) {
		t.Error("Unexpected message")
	}
}