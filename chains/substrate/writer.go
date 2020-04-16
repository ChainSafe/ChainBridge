// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"

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

	// Ensure we only submit a vote if the proposal hasn't completed
	active, err := w.proposalNotCompleted(prop)
	if err != nil {
		w.log.Error("Failed to assert proposal state", "err", err)
		return false
	}
	if active {
		w.log.Trace("Acknowledging proposal on chain", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", fmt.Sprintf("%x", prop.resourceId), "method", prop.method)
		err = w.conn.SubmitTx(AcknowledgeProposal, prop.depositNonce, prop.sourceId, prop.resourceId, prop.call)
		if err != nil {
			w.log.Error("Failed to execute extrinsic", "err", err)
			return false
		}
	} else {
		w.log.Debug("Ignoring previously completed proposal", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", prop.resourceId)
	}

	return true
}

func (w *Writer) resolveResourceId(id [32]byte) (string, error) {
	var res []byte
	exists, err := w.conn.queryStorage("Bridge", "Resources", id[:], nil, &res)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("resource %x not found on chain", id)
	}

	return string(res), nil
}

func (w *Writer) proposalNotCompleted(prop *proposal) (bool, error) {
	var voteRes voteState
	srcId, err := types.EncodeToBytes(prop.sourceId)
	if err != nil {
		return false, err
	}
	propBz, err := prop.encode()
	if err != nil {
		return false, err
	}
	exists, err := w.conn.queryStorage("Bridge", "Votes", srcId, propBz, &voteRes)
	if err != nil {
		return false, err
	}

	// Ensure proposal either doesn't yet exist or is still active
	if !exists || voteRes.Status.IsActive {
		return true, nil
	}
	return false, nil
}

func (w *Writer) Stop() error {
	return nil
}
