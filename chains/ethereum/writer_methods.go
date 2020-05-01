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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var NonceRetryInterval = time.Second * 2

const TxRetryLimit = 10

func constructErc20ProposalData(amount []byte, resourceId msg.ResourceId, recipient []byte) []byte {
	var data []byte
	data = append(data, resourceId[:]...)                   // resourceId
	data = append(data, common.LeftPadBytes(amount, 32)...) // amount

	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	data = append(data, common.LeftPadBytes(recipientLen, 32)...) // Length of recipient
	data = append(data, recipient...)                             // recipient
	return data
}

func constructErc721ProposalData(tokenId []byte, resourceId msg.ResourceId, recipient []byte, metadata []byte) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(tokenId, 32)...)
	data = append(data, resourceId[:]...)

	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	data = append(data, common.LeftPadBytes(recipientLen, 32)...)
	data = append(data, recipient...)

	metadataLen := big.NewInt(int64(len(metadata))).Bytes()
	data = append(data, common.LeftPadBytes(metadataLen, 32)...)
	data = append(data, metadata...)
	return data
}

func constructGenericProposalData(resourceId msg.ResourceId, metadata []byte) []byte {
	var data []byte
	data = append(resourceId[:], math.PaddedBigBytes(big.NewInt(int64(len(metadata))), 32)...)
	data = append(data, metadata...)
	return data
}

// proposalIsComplete returns true if the proposal state is either Passed(2) or Transferred(3)
func (w *writer) proposalIsComplete(destId msg.ChainId, nonce msg.Nonce) bool {
	prop, err := w.bridgeContract.GetProposal(&bind.CallOpts{}, uint8(destId), nonce.Big())
	if err != nil {
		w.log.Error("Failed to check deposit proposal", "err", err)
		return false
	}
	return prop.Status >= PassedStatus // Passed (2) or Transferred (3)
}

func (w *writer) createErc20Proposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal")

	data := constructErc20ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte))
	hash := utils.Hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		return true
	}
	// watch for execution event
	go w.watchThenExecute(m, w.cfg.erc20HandlerContract, data)

	w.voteProposal(m, hash)

	return true
}

func (w *writer) createErc721Proposal(m msg.Message) bool {
	w.log.Info("Creating erc721 proposal")

	data := constructErc721ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte), m.Payload[2].([]byte))
	hash := utils.Hash(append(w.cfg.erc721HandlerContract.Bytes(), data...))

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		return true
	}
	// watch for execution event
	go w.watchThenExecute(m, w.cfg.erc721HandlerContract, data)

	w.voteProposal(m, hash)

	return true
}

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

	// watch for execution event
	go w.watchThenExecute(m, w.cfg.genericHandlerContract, data)

	w.voteProposal(m, dataHash)

	return true
}

func (w *writer) watchThenExecute(m msg.Message, handler common.Address, data []byte) {
	w.log.Trace("Watching for finalization event", "depositNonce", m.DepositNonce)

	query := buildQuery(w.cfg.bridgeContract, utils.ProposalFinalized, w.cfg.startBlock, nil)
	eventSubscription, err := w.conn.subscribeToEvent(query)
	if err != nil {
		w.log.Error("Failed to subscribe to finalization event", "err", err)
	}

	for {
		select {
		case evt := <-eventSubscription.ch:
			sourceId := evt.Topics[1].Big().Uint64()
			destId := evt.Topics[2].Big().Uint64()
			depositNonce := evt.Topics[3].Big().Uint64()

			if m.Source == msg.ChainId(sourceId) &&
				m.Destination == msg.ChainId(destId) &&
				uint64(m.DepositNonce) == depositNonce {
				eventSubscription.sub.Unsubscribe()
				w.executeProposal(m, handler, data)
				return
			} else {
				w.log.Trace("Ignoring finalization event", "source", sourceId, "dest", destId, "nonce", depositNonce)
			}
		case err := <-eventSubscription.sub.Err():
			if err != nil {
				w.log.Error("finalization subscription error", "err", err)
				return
			}
		}
	}
}

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
			w.log.Info("Nonce too low, will retry")
			time.Sleep(NonceRetryInterval)
		} else {
			w.log.Info("Unexpected tx error", "err", err)
			time.Sleep(NonceRetryInterval)
		}
	}
}
