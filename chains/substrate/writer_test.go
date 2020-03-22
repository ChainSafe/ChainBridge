// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	message "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"gotest.tools/assert"
)

func assertProposalState(conn *Connection, hash types.Hash, prop *proposal, votes VoteState) error {
	var callRes types.Call
	log15.Trace("Fetching proposal call", "hash", hash.Hex())
	err := conn.queryStorage("Bridge", "Proposals", hash[:], nil, &callRes)
	if err != nil {
		return fmt.Errorf("failed to query proposals: %s", err)
	}
	if !types.Eq(&callRes, &prop.call) {
		return fmt.Errorf("Proposal state incorrect.\n\tExpected: %#v\n\tGot: %#v", prop.call, callRes)
	}

	log15.Trace("Fetching votes", "hash", hash.Hex())
	var voteRes VoteState
	err = conn.queryStorage("Bridge", "Votes", hash[:], nil, &voteRes)
	if err != nil {
		return fmt.Errorf("failed to query votes: %s", err)
	}
	if !reflect.DeepEqual(&voteRes, &votes) {
		return fmt.Errorf("Vote state incorrect.\n\tExpected: %#v\n\tGot: %#v", votes, voteRes)
	}
	return nil
}

func TestWriter_ResolveMessage_DepositAsset(t *testing.T) {
	ac, bc := createAliceAndBobConnections(t)

	alice := NewWriter(ac)
	bob := NewWriter(bc)

	// Assert Bob's starting balances
	var startingBalance types.U128
	getFreeBalance(bob.conn, &startingBalance)

	// Construct the message to initiate a vote
	amount := types.MustHexDecodeString("1000000000")
	data := append(bob.conn.key.PublicKey, amount...)
	m := message.Message{
		Source:      0,
		Destination: 1,
		Type:        message.DepositAssetType,
		DepositId:   0,
		To:          bob.conn.key.PublicKey,
		Metadata:    data,
	}

	// Create a proposal to help us check results
	prop, err := createProposalFromAssetTx(m, alice.conn.meta)
	if err != nil {
		t.Fatal(err)
	}

	// First, ensure the proposal doesn't already exist
	initialVotes := VoteState{}
	emptyProp := &proposal{call: types.Call{Args: types.Args{}}}
	assert.NilError(t, assertProposalState(alice.conn, prop.hash, emptyProp, initialVotes))

	// Submit the message for processing
	ok := alice.ResolveMessage(m)
	if !ok {
		t.Fatal("Alice failed to resolve the message")
	}

	// Now check if the proposal exists on chain
	singleVoteState := VoteState{
		VotesFor: []types.AccountID{types.NewAccountID(alice.conn.key.PublicKey)},
		Hash:     prop.hash,
	}
	assert.NilError(t, assertProposalState(alice.conn, prop.hash, prop, singleVoteState))

	// Submit a second vote from Bob this time
	ok = bob.ResolveMessage(m)
	if !ok {
		t.Fatalf("Bob failed to resolve the message")
	}

	// Check the vote was added
	finalVoteState := VoteState{
		VotesFor: []types.AccountID{
			types.NewAccountID(alice.conn.key.PublicKey),
			types.NewAccountID(bob.conn.key.PublicKey),
		},
		Hash: prop.hash,
	}
	assert.NilError(t, assertProposalState(alice.conn, prop.hash, prop, finalVoteState))

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
