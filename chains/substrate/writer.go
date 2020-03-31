// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	conn *Connection
	log  log15.Logger
}

func NewWriter(conn *Connection, log log15.Logger) *Writer {
	return &Writer{
		conn: conn,
		log:  log,
	}
}

func (w *Writer) Start() error {
	return nil
}

func (w *Writer) ResolveMessage(m msg.Message) bool {
	switch m.Type {
	case msg.DepositAssetType:
		meta := w.conn.getMetadata()
		prop, err := createAssetTxProposal(m, &meta)
		if err != nil {
			w.log.Error("Failed to construct assetTxProposal from message", "err", err)
			return false
		}
		if w.proposalExists(prop) {
			w.log.Debug("Voting for an existing assetTxProposal", "nonce", prop.depositNonce)
			err = w.conn.SubmitTx(Approve, prop.depositNonce, prop.call)
		} else {
			w.log.Trace("Creating a new assetTxProposal", "nonce", prop.depositNonce)
			err = w.conn.SubmitTx(CreateProposal, prop.depositNonce, prop.call)
		}
		if err != nil {
			w.log.Error("Failed to execute extrinsic", "err", err)
			return false
		}

	default:
		w.log.Error("Unrecognized message type", "type", m.Type)
	}
	return true

}

func (w *Writer) proposalExists(prop *assetTxProposal) bool {
	key, err := types.EncodeToBytes(prop.getKey())
	if err != nil {
		w.log.Error("Failed to encode key", "nonce", prop.depositNonce, "err", err)
	}

	var expected voteState
	ok, err := w.conn.queryStorage("Bridge", "Votes", key, nil, &expected)
	if err != nil {
		w.log.Error("Failed to check proposals existence", "err", err)
		return false
	}
	return ok
}

func (w *Writer) Stop() error {
	return nil
}
