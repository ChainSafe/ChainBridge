// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

const VoteDepositProposalMethod = "voteDepositProposal"
const ExecuteDepositMethod = "executeDepositProposal"

func constructErc20ProposalData(amount, tokenId, recipient []byte) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(amount, 32)...) // amount

	tokenIdLen := big.NewInt(int64(len(tokenId))).Bytes()
	data = append(data, common.LeftPadBytes(tokenIdLen, 32)...) // len(tokenId)
	data = append(data, tokenId...)                             // tokenId (chainId)

	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	data = append(data, common.LeftPadBytes(recipientLen, 32)...) // Length of recipient
	data = append(data, recipient...)                             // recipient
	return data
}

func (w *Writer) createErc20DepositProposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal")

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	data := constructErc20ProposalData(m.Metadata[0].([]byte), m.Metadata[1].([]byte), m.Metadata[2].([]byte))
	hash := hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))
	// watch for execution event
	go w.watchAndExecute(m, w.cfg.erc20HandlerContract, data)

	log15.Trace("Submitting CreateDepositPropsoal transaction", "source", m.Source, "depositNonce", m.DepositNonce)
	_, err = w.bridgeContract.Transact(
		opts,
		VoteDepositProposalMethod,
		big.NewInt(int64(m.Source)),
		u32toBigInt(m.DepositNonce),
		hash,
	)

	if err != nil {
		w.log.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	nonce.lock.Unlock()

	return true
}

func (w *Writer) createGenericDepositProposal(m msg.Message) bool {
	w.log.Info("Creating generic proposal", "handler", w.cfg.genericHandlerContract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	h := m.Metadata[0].([]byte)
	dataHash := hash(append(w.cfg.genericHandlerContract.Bytes(), h...))

	// watch for execution event
	go w.watchAndExecute(m, w.cfg.genericHandlerContract, h)

	log15.Trace("Submitting CreateDepositProposal transaction", "source", m.Source, "depositNonce", m.DepositNonce)
	_, err = w.bridgeContract.Transact(
		opts,
		VoteDepositProposalMethod,
		big.NewInt(int64(m.Source)),
		u32toBigInt(m.DepositNonce),
		dataHash,
	)

	if err != nil {
		w.log.Error("Failed to submit voteDepsitProposal transaction", "err", err)
		return false
	}
	nonce.lock.Unlock()

	return true
}

func (w *Writer) watchAndExecute(m msg.Message, handler common.Address, data []byte) {
	w.log.Trace("Watching for finalization event", "depositNonce", m.DepositNonce)
	// TODO: Skip existing blocks
	query := buildQuery(w.cfg.bridgeContract, DepositProposalFinalized, w.cfg.startBlock)
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
				m.DepositNonce == uint32(depositNonce) {
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

func (w *Writer) executeProposal(m msg.Message, handler common.Address, data []byte) {
	w.log.Info("Executing proposal", "to", w.conn.cfg.bridgeContract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return
	}

	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		ExecuteDepositMethod,
		big.NewInt(int64(m.Source)),
		u32toBigInt(m.DepositNonce),
		handler,
		data,
	)

	if err != nil {
		w.log.Error("Failed to submit executeDeposit transaction", "err", err)
	}
}
