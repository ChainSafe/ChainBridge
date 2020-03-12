package substrate

import (
	"testing"

	message "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func TestWriter_ResolveMessage_DepositAsset(t *testing.T) {
	conn := NewConnection(TestEndpoint, &TestKeyringPairAlice)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	w := NewWriter(conn)

	bob := TestKeyringPairBob.PublicKey
	amount := []byte{0x10, 0x10, 0x00, 0x00}
	data := append(bob, amount...)

	m := message.Message{
		Source:      0,
		Destination: 1,
		Type:        message.DepositAssetType,
		DepositId:   0,
		To:          bob,
		Metadata:    data,
	}

	prop, err := createProposalFromAssetTx(m, conn.meta)
	if err != nil {
		t.Fatal(err)
	}

	// First, ensure the proposal doesn't already exist
	var result types.Call
	err = conn.queryStorage("Bridge", "Proposals", prop.hash[:], nil, &result)
	if err != nil {
		t.Fatal(err)
	}
	if types.Eq(&result, &prop.call) {
		t.Fatal("Proposal already exists on chain")
	}

	// Submit the message for processing
	ok := w.ResolveMessage(m)
	if !ok {
		t.Fatal("Not okay :'(")
	}

	// Now check if the proposal exists on chain
	err = conn.queryStorage("Bridge", "Proposals", prop.hash[:], nil, &result)
	if err != nil {
		t.Fatal(err)
	}
	if !types.Eq(&result, &prop.call) {
		t.Fatalf("Did not get correct call back. \n\tGot: %#v\n\tExpected: %#v\n", result, prop.call)
	}
}
