package substrate

import (
	"encoding/binary"
	"reflect"
	"testing"
	"time"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const ListenerTimeout = time.Second * 30

type mockRouter struct {
	msgs chan msg.Message
}

func (r *mockRouter) Send(message msg.Message) error {
	r.msgs <- message
	return nil
}

func whitelistChain(c *Connection, id uint32) error {
	destId := types.U32(id)
	call, err := types.NewCall(
		c.meta,
		WhitelistChain.String(),
		destId,
	)
	if err != nil {
		return err
	}

	err = c.SubmitTx(Sudo, call)
	if err != nil {
		return err
	}
	return nil
}

func Test_AssetTx(t *testing.T) {
	ac := createAliceConnection(t)
	r := &mockRouter{msgs: make(chan msg.Message)}
	alice := NewListener(ac, "Alice", 1)
	alice.SetRouter(r)
	err := alice.Start()
	if err != nil {
		t.Fatal(err)
	}

	// First we have to whitelist the destination chain with sudo
	var destId uint32 = 0
	err = whitelistChain(ac, destId)
	if err != nil {
		t.Fatal(err)
	}

	assetHash, err := types.NewHashFromHexString("0x16078eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f2")
	if err != nil {
		t.Fatal(err)
	}
	recipient := types.Bytes([]byte{0xAB, 0xCD})

	meta := make([]byte, 8)
	binary.LittleEndian.PutUint32(meta, destId)

	expected := msg.Message{
		Source:       1,
		Destination:  0,
		Type:         msg.DepositAssetType,
		DepositNonce: 0,
		To:           recipient,
		Metadata:     meta,
	}

	err = ac.SubmitTx(ExampleTransferHash, assetHash, recipient, destId)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case m := <-r.msgs:
		if !reflect.DeepEqual(expected, m) {
			t.Fatalf("Unexpected message.\n\tExpected: %#v\n\tGot: %#v\n", expected, m)
		}
	case <-time.After(ListenerTimeout):
		t.Fatalf("test timed out")

	}
}
