// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

func constructErc20ProposalData(amount []byte, resourceId msg.ResourceId, recipient []byte) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(amount, 32)...) // amount
	data = append(data, resourceId[:]...)                   // resourceId

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
	prop, err := w.bridgeContract.GetDepositProposal(&bind.CallOpts{}, uint8(destId), nonce.Big())
	if err != nil {
		log.Error("Failed to check deposit proposal", "err", err)
		return false
	}
	return prop.Status >= PassedStatus // Passed (2) or Transferred (3)
}

func (w *writer) createErc20Proposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal")

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	data := constructErc20ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte))
	hash := utils.Hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		nonce.lock.Unlock()
		return true
	}
	// watch for execution event
	go w.watchAndExecute(m, w.cfg.erc20HandlerContract, data)

	w.log.Debug("Submitting CreateDepositProposal for ERC20", "source", m.Source, "depositNonce", m.DepositNonce)

	_, err = w.bridgeContract.VoteDepositProposal(
		opts,
		uint8(m.Source),
		m.DepositNonce.Big(),
		hash,
	)

	if err != nil {
		w.log.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	nonce.lock.Unlock()

	return true
}

func (w *writer) createErc721Proposal(m msg.Message) bool {
	w.log.Info("Creating erc721 proposal")

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	data := constructErc721ProposalData(m.Payload[0].([]byte), m.ResourceId, m.Payload[1].([]byte), m.Payload[2].([]byte))
	hash := utils.Hash(append(w.cfg.erc721HandlerContract.Bytes(), data...))

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		nonce.lock.Unlock()
		return true
	}
	// watch for execution event
	go w.watchAndExecute(m, w.cfg.erc721HandlerContract, data)

	w.log.Debug("Submitting CreateDepositProposal for ERC721", "source", m.Source, "depositNonce", m.DepositNonce)

	_, err = w.bridgeContract.VoteDepositProposal(
		opts,
		uint8(m.Source),
		m.DepositNonce.Big(),
		hash,
	)

	if err != nil {
		w.log.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	nonce.lock.Unlock()

	return true
}

func (w *writer) createGenericDepositProposal(m msg.Message) bool {
	w.log.Info("Creating generic proposal", "handler", w.cfg.genericHandlerContract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	metadata := m.Payload[0].([]byte)
	data := constructGenericProposalData(m.ResourceId, metadata)
	toHash := append(w.cfg.genericHandlerContract.Bytes(), data...)
	dataHash := utils.Hash(toHash)

	if w.proposalIsComplete(m.Destination, m.DepositNonce) {
		w.log.Debug("Proposal complete, not voting")
		nonce.lock.Unlock()
		return true
	}

	// watch for execution event
	go w.watchAndExecute(m, w.cfg.genericHandlerContract, data)

	w.log.Trace("Submitting CreateDepositProposal transaction", "source", m.Source, "depositNonce", m.DepositNonce)
	_, err = w.bridgeContract.VoteDepositProposal(
		opts,
		uint8(m.Source),
		m.DepositNonce.Big(),
		dataHash,
	)

	if err != nil {
		w.log.Error("Failed to submit voteDepsitProposal transaction", "err", err)
		return false
	}
	nonce.lock.Unlock()

	return true
}

func (w *writer) watchAndExecute(m msg.Message, handler common.Address, data []byte) {
	w.log.Trace("Watching for finalization event", "depositNonce", m.DepositNonce)
	// TODO: Skip existing blocks
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

func (w *writer) executeProposal(m msg.Message, handler common.Address, data []byte) {
	w.log.Info("Executing proposal", "handler", handler, "data", fmt.Sprintf("%x", data))

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return
	}
	defer nonce.lock.Unlock()

	_, err = w.bridgeContract.ExecuteDepositProposal(
		opts,
		uint8(m.Source),
		m.DepositNonce.Big(),
		handler,
		data,
	)

	if err != nil {
		w.log.Warn("Failed to execute proposal, may already be complete", "err", err)
	}
}
