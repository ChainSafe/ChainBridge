// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	conn *Connection
}

func NewWriter(conn *Connection) *Writer {
	return &Writer{
		conn: conn,
	}
}

func (w *Writer) Start() error {
	return nil
}

func (w *Writer) ResolveMessage(m msg.Message) bool {
	switch m.Type {
	case msg.DepositAssetType:
		prop, err := createProposalFromAssetTx(m, w.conn.meta)
		if err != nil {
			log15.Error("Failed to construct proposal from message", "err", err)
			return false
		}
		if w.proposalExists(prop) {
			log15.Debug("Voting for an existing proposal", "hash", prop.hash)
			err = w.conn.SubmitTx(Vote, prop.hash, true)
		} else {
			log15.Trace("Creating a new proposal", "hash", prop.hash.Hex())
			err = w.conn.SubmitTx(CreateProposal, prop.hash, prop.call)
		}
		if err != nil {
			log15.Error("Failed to execute extrinsic", "err", err)
			return false
		}

	default:
		log15.Error("Unrecognized message type", "type", m.Type)
	}
	return true

}

func (w *Writer) proposalExists(prop *proposal) bool {
	var expected types.Call
	err := w.conn.queryStorage("Bridge", "Proposals", prop.hash[:], nil, &expected)
	if err != nil {
		log15.Error("Failed to check proposals existence", "err", err)
		return false
	}
	// TODO: We want to know if it exists on chain, but this actually tell us they aren't equal.
	// If prop is a valid proposal we might want to freak out.
	if types.Eq(expected, prop.call) {
		return true
	} else {
		return false
	}
}

func (w *Writer) Stop() error {
	return nil
}
