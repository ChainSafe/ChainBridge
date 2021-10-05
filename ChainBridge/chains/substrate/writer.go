// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/ChainSafe/chainbridge-utils/core"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var _ core.Writer = &writer{}

var AcknowledgeProposal utils.Method = utils.BridgePalletName + ".acknowledge_proposal"
var TerminatedError = errors.New("terminated")

type writer struct {
	conn       *Connection
	log        log15.Logger
	sysErr     chan<- error
	metrics    *metrics.ChainMetrics
	extendCall bool // Extend extrinsic calls to substrate with ResourceID.Used for backward compatibility with example pallet.
}

func NewWriter(conn *Connection, log log15.Logger, sysErr chan<- error, m *metrics.ChainMetrics, extendCall bool) *writer {
	return &writer{
		conn:       conn,
		log:        log,
		sysErr:     sysErr,
		metrics:    m,
		extendCall: extendCall,
	}
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
		w.sysErr <- fmt.Errorf("failed to construct proposal (chain=%d, name=%s) Error: %w", m.Destination, w.conn.name, err)
		return false
	}

	for i := 0; i < BlockRetryLimit; i++ {
		// Ensure we only submit a vote if the proposal hasn't completed
		valid, reason, err := w.proposalValid(prop)
		if err != nil {
			w.log.Error("Failed to assert proposal state", "err", err)
			time.Sleep(BlockRetryInterval)
			continue
		}

		// If active submit call, otherwise skip it. Retry on failure.
		if valid {
			w.log.Info("Acknowledging proposal on chain", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", fmt.Sprintf("%x", prop.resourceId), "method", prop.method)

			err = w.conn.SubmitTx(AcknowledgeProposal, prop.depositNonce, prop.sourceId, prop.resourceId, prop.call)
			if err != nil && err.Error() == TerminatedError.Error() {
				return false
			} else if err != nil {
				w.log.Error("Failed to execute extrinsic", "err", err)
				time.Sleep(BlockRetryInterval)
				continue
			}
			if w.metrics != nil {
				w.metrics.VotesSubmitted.Inc()
			}
			return true
		} else {
			w.log.Info("Ignoring proposal", "reason", reason, "nonce", prop.depositNonce, "source", prop.sourceId, "resource", prop.resourceId)
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

// proposalValid asserts the state of a proposal. If the proposal is active and this relayer
// has not voted, it will return true. Otherwise, it will return false with a reason string.
func (w *writer) proposalValid(prop *proposal) (bool, string, error) {
	var voteRes voteState
	srcId, err := types.EncodeToBytes(prop.sourceId)
	if err != nil {
		return false, "", err
	}
	propBz, err := prop.encode()
	if err != nil {
		return false, "", err
	}
	exists, err := w.conn.queryStorage(utils.BridgeStoragePrefix, "Votes", srcId, propBz, &voteRes)
	if err != nil {
		return false, "", err
	}

	if !exists {
		return true, "", nil
	} else if voteRes.Status.IsActive {
		if containsVote(voteRes.VotesFor, types.NewAccountID(w.conn.key.PublicKey)) ||
			containsVote(voteRes.VotesAgainst, types.NewAccountID(w.conn.key.PublicKey)) {
			return false, "already voted", nil
		} else {
			return true, "", nil
		}
	} else {
		return false, "proposal complete", nil
	}
}

func containsVote(votes []types.AccountID, voter types.AccountID) bool {
	for _, v := range votes {
		if bytes.Equal(v[:], voter[:]) {
			return true
		}
	}
	return false
}
