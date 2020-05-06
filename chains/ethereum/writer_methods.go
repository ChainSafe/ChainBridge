// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// Number of blocks to wait for an finalization event
var ExecuteBlockWatchLimit = 50

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var NonceRetryInterval = time.Second * 2

const TxRetryLimit = 10

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
	data = append(data, common.LeftPadBytes(tokenId, 32)...) // tokenId ([]byte)
	data = append(data, resourceId[:]...)                    // resourceId (bytes32)

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

// proposalIsComplete returns true if the proposal state is either Passed(2) or Transferred(3)
func (w *writer) proposalIsComplete(destId msg.ChainId, nonce msg.Nonce) bool {
	prop, err := w.bridgeContract.GetProposal(&bind.CallOpts{From: w.conn.kp.CommonAddress()}, uint8(destId), nonce.Big())
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status >= PassedStatus // Passed (2) or Transferred (3)
}

// createErc20Proposal creates an Erc20 proposal
// returns true if the proposal is succesfully created or is complete
func (w *writer) createErc20Proposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal")

	data := constructErc20ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte))
	hash := utils.Hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		return true
	}

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.latestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, w.cfg.erc20HandlerContract, data, latestBlock)

	w.voteProposal(m, hash)

	return true
}

// createErc721Proposal creates an Erc721 proposal
// returns true if the proposal is succesfully created or is complete
func (w *writer) createErc721Proposal(m msg.Message) bool {
	w.log.Info("Creating erc721 proposal")

	data := constructErc721ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte), m.Payload[2].([]byte))
	hash := utils.Hash(append(w.cfg.erc721HandlerContract.Bytes(), data...))

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		return true
	}

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.latestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, w.cfg.erc721HandlerContract, data, latestBlock)

	w.voteProposal(m, hash)

	return true
}

// createGenericDepositProposal creates a generic proposal
// returns true if the proposal is complete or is succesfully created
func (w *writer) createGenericDepositProposal(m msg.Message) bool {
	w.log.Info("Creating generic proposal", "handler", w.cfg.genericHandlerContract)

	metadata := m.Payload[0].([]byte)
	data := constructGenericProposalData(m.ResourceId, metadata)
	toHash := append(w.cfg.genericHandlerContract.Bytes(), data...)
	dataHash := utils.Hash(toHash)

	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		return true
	}

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.latestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, w.cfg.genericHandlerContract, data, latestBlock)

	w.voteProposal(m, dataHash)

	return true
}

// watchThenExecute watches for the latest block and executes once the matching finalized event is found
func (w *writer) watchThenExecute(m msg.Message, handler common.Address, data []byte, latestBlock *big.Int) {
	w.log.Trace("Watching for finalization event", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)

	// watching for the latest block, querying and matching the finalized event will be retried up to ExecuteBlockWatchLimit times
	for i := 0; i < ExecuteBlockWatchLimit; i++ {

		// watches for the lastest block
		err := w.conn.waitForBlock(latestBlock)
		if err != nil {
			w.log.Error("Waiting for block failed", "err", err)
			return
		}

		// querying for logs
		query := buildQuery(w.cfg.bridgeContract, utils.ProposalFinalized, latestBlock, latestBlock)
		evts, err := w.conn.conn.FilterLogs(w.conn.ctx, query)
		if err != nil {
			log.Error("Failed to fetch logs", "err", err)
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
				w.executeProposal(m, handler, data)
				return
			} else {
				w.log.Trace("Ignoring finalization event", "source", sourceId, "dest", destId, "nonce", depositNonce)
			}
		}

		latestBlock = latestBlock.Add(latestBlock, big.NewInt(1))
	}
	log.Warn("Block watch limit exceeded, skipping execution", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
}

// voteProposal submits a vote proposal
// a vote proposal will try to be submitted up to the TxRetryLimit times
func (w *writer) voteProposal(m msg.Message, hash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		err := w.lockAndUpdateNonce()
		if err != nil {
			w.log.Error("Failed to update nonce", "err", err)
			continue
		}

		w.log.Debug("Submitting VoteProposal", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
		_, err = w.bridgeContract.VoteProposal(
			w.opts,
			uint8(m.Source),
			m.DepositNonce.Big(),
			hash,
		)
		w.unlockNonce()

		if err == nil {
			return
		} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
			w.log.Info("Nonce too low, will retry")
			time.Sleep(NonceRetryInterval)
		} else {
			w.log.Info("Unexpected tx error", "err", err)
			time.Sleep(NonceRetryInterval)
		}
	}
}

// executeProposal executes the proposal
func (w *writer) executeProposal(m msg.Message, handler common.Address, data []byte) {
	for i := 0; i < TxRetryLimit; i++ {
		err := w.lockAndUpdateNonce()
		if err != nil {
			w.log.Error("Failed to update nonce", "err", err)
			return
		}

		w.log.Info("Executing proposal", "handler", handler, "data", fmt.Sprintf("%x", data))
		_, err = w.bridgeContract.ExecuteProposal(
			w.opts,
			uint8(m.Source),
			m.DepositNonce.Big(),
			handler,
			data,
		)
		w.unlockNonce()

		if err == nil {
			return
		} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
			w.log.Debug("Nonce too low, will retry")
			time.Sleep(NonceRetryInterval)
		} else {
			w.log.Warn("Execution failed, proposal may already be complete", "err", err)
			time.Sleep(NonceRetryInterval)
		}
	}
}
