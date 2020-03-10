package substrate

import (
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/core"
)

func TestListener(t *testing.T) {
	alice := NewConnection(TestEndpoint)
	alice.key = Alice
	err := alice.Connect()
	if err != nil {
		t.Fatal(err)
	}
	
	cfg := core.ChainConfig{
		Id:            0,
		Endpoint:      "",
		From:          "",
		Keystore:      nil,
		Subscriptions: []string{ValidatorAdded, ValidatorRemoved, VoteFor, VoteAgainst, ProposalSucceeded, ProposalFailed},
		Opts:          nil,
	}
	
	l := NewListener(alice, &cfg)

	err = l.Start()
	if err != nil {
		t.Fatal(err)
	}

	for {}
}