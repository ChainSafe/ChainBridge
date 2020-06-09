// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"math/big"
	"time"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 100

// Time between retrying a failed tx
const TxRetryInterval = time.Second * 2

// Maximum number of tx retries before exiting
const TxRetryLimit = 10

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var ErrFatalTx = errors.New("submission of transaction failed")
var ErrFatalQuery = errors.New("query of chain state failed")

// containsVote is used to search the vote lists of a proposal
func containsVote(votes []common.Address, voter common.Address) bool {
	for _, v := range votes {
		if v == voter {
			return true
		}
	}
	return false
}

// constructErc20ProposalData returns the bytes to construct a proposal suitable for Erc20
func constructErc20ProposalData(amount []byte, resourceId msg.ResourceId, recipient []byte) []byte {
	var data []byte
	data = append(data, resourceId[:]...)                   // resourceId (bytes32)
	data = append(data, common.LeftPadBytes(amount, 32)...) // amount (uint256)

	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	data = append(data, common.LeftPadBytes(recipientLen, 32)...) // length of recipient (uint256)
	data = append(data, recipient...)                             // recipient ([]byte)
	return data
}

// constructErc721ProposalData returns the bytes to construct a proposal suitable for Erc721
func constructErc721ProposalData(tokenId []byte, resourceId msg.ResourceId, recipient []byte, metadata []byte) []byte {
	var data []byte
	data = append(data, resourceId[:]...)                    // resourceId (bytes32)
	data = append(data, common.LeftPadBytes(tokenId, 32)...) // tokenId ([]byte)

	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	data = append(data, common.LeftPadBytes(recipientLen, 32)...) // length of recipient
	data = append(data, recipient...)                             // recipient ([]byte)

	metadataLen := big.NewInt(int64(len(metadata))).Bytes()
	data = append(data, common.LeftPadBytes(metadataLen, 32)...) // length of metadata (uint256)
	data = append(data, metadata...)                             // metadata ([]byte)
	return data
}

// constructGenericProposalData returns the bytes to construct a generic proposal
func constructGenericProposalData(resourceId msg.ResourceId, metadata []byte) []byte {
	var data []byte

	metadataLen := big.NewInt(int64(len(metadata)))
	data = append(resourceId[:], math.PaddedBigBytes(metadataLen, 32)...) // length of metadata (uint256)
	data = append(data, metadata...)                                      // metadata ([]byte)
	return data
}

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId msg.ChainId, nonce msg.Nonce) bool {
	prop, err := w.bridgeContract.GetProposal(w.callOpts, uint8(srcId), uint64(nonce))
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId msg.ChainId, nonce msg.Nonce) bool {
	prop, err := w.bridgeContract.GetProposal(w.callOpts, uint8(srcId), uint64(nonce))
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
}

// hasVoted checks if this relayer has already voted
func (w *writer) hasVoted(srcId msg.ChainId, nonce msg.Nonce) bool {
	prop, err := w.bridgeContract.GetProposal(w.callOpts, uint8(srcId), uint64(nonce))
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}

	return containsVote(prop.YesVotes, w.conn.kp.CommonAddress()) || containsVote(prop.NoVotes, w.conn.kp.CommonAddress())
}

// createErc20Proposal creates an Erc20 proposal.
// Returns true if the proposal is successfully created or is complete
func (w *writer) createErc20Proposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal", "src", m.Source, "nonce", m.DepositNonce)

	data := constructErc20ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte))
	hash := utils.Hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.latestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, data, latestBlock)

	w.voteProposal(m, hash)

	return true
}

// createErc721Proposal creates an Erc721 proposal.
// Returns true if the proposal is succesfully created or is complete
func (w *writer) createErc721Proposal(m msg.Message) bool {
	w.log.Info("Creating erc721 proposal", "src", m.Source, "nonce", m.DepositNonce)

	data := constructErc721ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte), m.Payload[2].([]byte))
	hash := utils.Hash(append(w.cfg.erc721HandlerContract.Bytes(), data...))

	// Capture latest block so we know where to watch from
	latestBlock, err := w.conn.latestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, data, latestBlock)

	w.voteProposal(m, hash)

	return true
}

// createGenericDepositProposal creates a generic proposal
// returns true if the proposal is complete or is succesfully created
func (w *writer) createGenericDepositProposal(m msg.Message) bool {
	w.log.Info("Creating generic proposal", "src", m.Source, "nonce", m.DepositNonce)

	metadata := m.Payload[0].([]byte)
	data := constructGenericProposalData(m.ResourceId, metadata)
	toHash := append(w.cfg.genericHandlerContract.Bytes(), data...)
	dataHash := utils.Hash(toHash)

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.latestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, data, latestBlock)

	w.voteProposal(m, dataHash)

	return true
}

// watchThenExecute watches for the latest block and executes once the matching finalized event is found
func (w *writer) watchThenExecute(m msg.Message, data []byte, latestBlock *big.Int) {
	w.log.Info("Watching for finalization event", "src", m.Source, "nonce", m.DepositNonce)

	// watching for the latest block, querying and matching the finalized event will be retried up to ExecuteBlockWatchLimit times
	for i := 0; i < ExecuteBlockWatchLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			// watch for the lastest block, retry up to BlockRetryLimit times
			for waitRetrys := 0; waitRetrys < BlockRetryLimit; waitRetrys++ {
				err := w.conn.waitForBlock(latestBlock)
				if err != nil {
					w.log.Error("Waiting for block failed", "err", err)
					// Exit if retries exceeded
					if waitRetrys+1 == BlockRetryLimit {
						w.log.Error("Waiting for block retries exceeded, shutting down")
						w.sysErr <- ErrFatalQuery
						return
					}
				} else {
					break
				}
			}

			// query for logs
			query := buildQuery(w.cfg.bridgeContract, utils.ProposalFinalized, latestBlock, latestBlock)
			evts, err := w.conn.conn.FilterLogs(w.conn.ctx, query)
			if err != nil {
				w.log.Error("Failed to fetch logs", "err", err)
				return
			}

			// execute the proposal once we find the matching finalized event
			for _, evt := range evts {
				sourceId := evt.Topics[1].Big().Uint64()
				destId := evt.Topics[2].Big().Uint64()
				depositNonce := evt.Topics[3].Big().Uint64()

				if m.Source == msg.ChainId(sourceId) &&
					m.Destination == msg.ChainId(destId) &&
					m.DepositNonce.Big().Uint64() == depositNonce {
					w.executeProposal(m, data)
					return
				} else {
					w.log.Trace("Ignoring finalization event", "src", sourceId, "dest", destId, "nonce", depositNonce)
				}
			}
			w.log.Trace("No finalization event found in current block", "block", latestBlock, "src", m.Source, "nonce", m.DepositNonce)
			latestBlock = latestBlock.Add(latestBlock, big.NewInt(1))
		}
	}
	log.Warn("Block watch limit exceeded, skipping execution", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
}

// voteProposal submits a vote proposal
// a vote proposal will try to be submitted up to the TxRetryLimit times
func (w *writer) voteProposal(m msg.Message, hash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.lockAndUpdateNonce()
			if err != nil {
				w.log.Error("Failed to update nonce", "err", err)
				continue
			}

			tx, err := w.bridgeContract.VoteProposal(
				w.opts,
				uint8(m.Source),
				uint64(m.DepositNonce),
				m.ResourceId,
				hash,
			)
			w.unlockNonce()

			if err == nil {
				w.log.Info("Submitted proposal vote", "tx", tx.Hash(), "src", m.Source, "depositNonce", m.DepositNonce)
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Voting failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.proposalIsComplete(m.Source, m.DepositNonce) {
				w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Vote transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

// executeProposal executes the proposal
func (w *writer) executeProposal(m msg.Message, data []byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.lockAndUpdateNonce()
			if err != nil {
				w.log.Error("Failed to update nonce", "err", err)
				return
			}

			tx, err := w.bridgeContract.ExecuteProposal(
				w.opts,
				uint8(m.Source),
				uint64(m.DepositNonce),
				data,
			)
			w.unlockNonce()

			if err == nil {
				w.log.Info("Submitted proposal execution", "tx", tx.Hash(), "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Error("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Execution failed, proposal may already be complete", "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for execution, tx will fail if we aren't the first to execute,
			// but there is no need to retry
			if w.proposalIsFinalized(m.Source, m.DepositNonce) {
				w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Execute transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}
