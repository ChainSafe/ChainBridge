// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	message "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"gotest.tools/assert"
)

func assertProposalState(conn *Connection, key *proposalKey, votes *voteState, hasValue bool) error {
	log15.Trace("Fetching votes", "key", key)
	var voteRes voteState
	keyBz, err := types.EncodeToBytes(key)
	if err != nil {
		return nil
	}
	ok, err := conn.queryStorage("Bridge", "Votes", keyBz, nil, &voteRes)
	if err != nil {
		return fmt.Errorf("failed to query votes: %s", err)
	}
	if hasValue {
		if !reflect.DeepEqual(&voteRes, votes) {
			return fmt.Errorf("Vote state incorrect.\n\tExpected: %#v\n\tGot: %#v", votes, &voteRes)
		}
	}

	if !ok && hasValue {
		return fmt.Errorf("expected vote to exists but is None")
	}

	return nil
}

func TestWriter_ResolveMessage_DepositAsset(t *testing.T) {
	ac, bc := createAliceAndBobConnections(t)

	alice := NewWriter(ac, TestLogger)
	bob := NewWriter(bc, TestLogger)

	// Assert Bob's starting balances
	var startingBalance types.U128
	getFreeBalance(bob.conn, &startingBalance)

	// Construct the message to initiate a vote
	amount := types.MustHexDecodeString("1000000000")
	data := append(bob.conn.key.PublicKey, amount...)
	m := message.Message{
		Source:       0,
		Destination:  1,
		Type:         message.DepositAssetType,
		DepositNonce: 0,
		To:           bob.conn.key.PublicKey,
		Metadata:     data,
	}

	// Create a assetTxProposal to help us check results
	meta := alice.conn.getMetadata()
	prop, err := createAssetTxProposal(m, &meta)
	if err != nil {
		t.Fatal(err)
	}

	// First, ensure the assetTxProposal doesn't already exist
	assert.NilError(t, assertProposalState(alice.conn, prop.getKey(), nil, false))

	// Submit the message for processing
	ok := alice.ResolveMessage(m)
	if !ok {
		t.Fatal("Alice failed to resolve the message")
	}

	// Now check if the assetTxProposal exists on chain
	singleVoteState := &voteState{
		VotesFor: []types.AccountID{types.NewAccountID(alice.conn.key.PublicKey)},
	}
	assert.NilError(t, assertProposalState(alice.conn, prop.getKey(), singleVoteState, true))

	// Submit a second vote from Bob this time
	ok = bob.ResolveMessage(m)
	if !ok {
		t.Fatalf("Bob failed to resolve the message")
	}

	// Check the vote was added
	finalVoteState := &voteState{
		VotesFor: []types.AccountID{
			types.NewAccountID(alice.conn.key.PublicKey),
			types.NewAccountID(bob.conn.key.PublicKey),
		},
	}
	assert.NilError(t, assertProposalState(alice.conn, prop.getKey(), finalVoteState, true))

	// Assert balance has updated
	// TODO: This doesn't account for gas and stakng rewards, will update once example chain is implemented
	var bBal types.U128
	getFreeBalance(bob.conn, &bBal)
	if bBal == startingBalance {
		t.Fatalf("Internal transaction failed to update Bobs balance")
	} else {
		t.Logf("Bob's new balance: %s (amount: %s)", bBal.String(), big.NewInt(0).Sub(bBal.Int, startingBalance.Int).String())
	}
}
