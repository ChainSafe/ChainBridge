// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
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
	var prop *proposal
	var err error
	meta := w.conn.getMetadata()

	switch m.Type {
	case msg.FungibleTransfer:
		prop, err = createFungibleProposal(m, &meta)
		if err != nil {
			w.log.Error("Failed to construct fungible token from message", "err", err)
			return false
		}

	case msg.GenericTransfer:
		prop, err = createGenericProposal(m, &meta)
		if err != nil {
			w.log.Error("Failed to construct generic transfer from message", "err", err)
			return false
		}

	default:
		w.log.Error("Unrecognized message type", "type", m.Type)
		return false
	}

	log15.Trace("Acknowledging proposal on chain", "nonce", prop.DepositNonce)
	err = w.conn.SubmitTx(AcknowledgeProposal, prop.DepositNonce, prop.Call)
	if err != nil {
		w.log.Error("Failed to execute extrinsic", "err", err)
		return false
	}

	return true
}

func (w *Writer) Stop() error {
	return nil
}
