// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"

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

	switch m.Type {
	case msg.FungibleTransfer:
		prop, err = w.createFungibleProposal(m)
		if err != nil {
			w.log.Error("Failed to construct fungible transfer from message", "err", err)
			return false
		}
	//case msg.NonFungibleTransfer:
	//	prop, err = w.createNonFungibleProposal(m)
	//	if err != nil {
	//		w.log.Error("Failed to construct nonfungible transfer from message", "err", err)
	//		return false
	//	}
	case msg.GenericTransfer:
		prop, err = w.createGenericProposal(m)
		if err != nil {
			w.log.Error("Failed to construct generic transfer from message", "err", err)
			return false
		}

	default:
		w.log.Error("Unrecognized message type", "type", m.Type)
		return false
	}

	w.log.Trace("Acknowledging proposal on chain", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", prop.resourceId)
	err = w.conn.SubmitTx(AcknowledgeProposal, prop.depositNonce, prop.sourceId, prop.resourceId, prop.call)
	if err != nil {
		w.log.Error("Failed to execute extrinsic", "err", err)
		return false
	}

	return true
}

func (w *Writer) resolveResourceId(id [32]byte) (string, error) {
	var res []byte
	exists, err := w.conn.queryStorage("Bridge", "Resources", id[:], nil, &res)
	if err != nil {
		w.log.Error("Failed to resolve resource ID", "err", err)
	}
	if !exists {
		w.log.Error("Resource not found on chain", "id", fmt.Sprintf("%x", id))
	}

	return string(res), nil
}

func (w *Writer) Stop() error {
	return nil
}
