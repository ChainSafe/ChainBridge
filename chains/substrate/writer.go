// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"errors"
	"fmt"
	"time"

	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var _ chains.Writer = &writer{}

var AcknowledgeProposal utils.Method = utils.BridgePalletName + ".acknowledge_proposal"
var TerminatedError = errors.New("terminated")

type writer struct {
	conn   *Connection
	log    log15.Logger
	sysErr chan<- error
}

func NewWriter(conn *Connection, log log15.Logger, sysErr chan<- error) *writer {
	return &writer{
		conn:   conn,
		log:    log,
		sysErr: sysErr,
	}
}

func (w *writer) start() error {
	return nil
}

func (w *writer) ResolveMessage(m msg.Message) bool {
	var prop *proposal
	var err error

	// Construct the proposal
	switch m.Type {
	case msg.FungibleTransfer:
		prop, err = w.createFungibleProposal(m)
	case msg.NonFungibleTransfer:
		prop, err = w.createNonFungibleProposal(m)
	case msg.GenericTransfer:
		prop, err = w.createGenericProposal(m)
	default:
		w.sysErr <- fmt.Errorf("unrecognized message type received (chain=%d, name=%s)", m.Destination, w.conn.name)
		return false
	}

	if err != nil {
		w.sysErr <- fmt.Errorf("failed to construct proposal (chain=%d, name=%s) Error: %s", m.Destination, w.conn.name, err)
		return false
	}

	for i := 0; i < BlockRetryLimit; i++ {
		// Ensure we only submit a vote if the proposal hasn't completed
		active, err := w.proposalNotCompleted(prop)
		if err != nil {
			w.log.Error("Failed to assert proposal state", "err", err)
			time.Sleep(BlockRetryInterval)
			continue
		}

		// If active submit call, otherwise skip it. Retry on failure.
		if active {
			w.log.Trace("Acknowledging proposal on chain", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", fmt.Sprintf("%x", prop.resourceId), "method", prop.method)
			err = w.conn.SubmitTx(AcknowledgeProposal, prop.depositNonce, prop.sourceId, prop.resourceId, prop.call)
			if err != nil && err.Error() == TerminatedError.Error() {
				return false
			} else if err != nil {
				w.log.Error("Failed to execute extrinsic", "err", err)
				time.Sleep(BlockRetryInterval)
				continue
			}
			return true
		} else {
			w.log.Debug("Ignoring previously completed proposal", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", prop.resourceId)
			return true
		}
	}
	return true
}

func (w *writer) resolveResourceId(id [32]byte) (string, error) {
	var res []byte
	exists, err := w.conn.queryStorage(utils.BridgeStoragePrefix, "Resources", id[:], nil, &res)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("resource %x not found on chain", id)
	}

	return string(res), nil
}

func (w *writer) proposalNotCompleted(prop *proposal) (bool, error) {
	var voteRes voteState
	srcId, err := types.EncodeToBytes(prop.sourceId)
	if err != nil {
		return false, err
	}
	propBz, err := prop.encode()
	if err != nil {
		return false, err
	}
	exists, err := w.conn.queryStorage(utils.BridgeStoragePrefix, "Votes", srcId, propBz, &voteRes)
	if err != nil {
		return false, err
	}

	// Ensure proposal either doesn't yet exist or is still active
	if !exists || voteRes.Status.IsActive {
		return true, nil
	}
	return false, nil
}
