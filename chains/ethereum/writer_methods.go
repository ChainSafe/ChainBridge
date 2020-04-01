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

func (w *Writer) createErc20DepositProposal(m msg.Message) bool {
	w.log.Info("Creating VoteDepositProposal transaction", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	var data []byte
	data = append(data, common.LeftPadBytes(m.Metadata[0].([]byte), 32)...)           // tokenId
	data = append(data, common.LeftPadBytes(m.Metadata[1].([]byte), 32)...)           // recipient
	data = append(data, common.LeftPadBytes(m.Metadata[2].(*big.Int).Bytes(), 32)...) // amount
	hash := hash(data)

	// watch for execution event
	go w.watchAndExecute(m, w.cfg.erc20HandlerContract, data)

	log15.Trace("Submitting CreateDepositPropsoal transaction")
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

	w.log.Info("Successfully submitted deposit proposal!", "source", m.Source, "depositNonce", m.DepositNonce)

	return true
}

func (w *Writer) watchAndExecute(m msg.Message, handler common.Address, data []byte) {
	w.log.Trace("Watching for finalization event", "depositNonce", m.DepositNonce)
	// TODO: Skip existing blocks
	query := buildQuery(w.cfg.contract, DepositProposalFinalized, big.NewInt(0))
	eventSubscription, err := w.conn.subscribeToEvent(query)
	if err != nil {
		log15.Error("Failed to subscribe to finalization event", "err", err)
	}

	for {
		select {
		case evt := <-eventSubscription.ch:
			sourceId := evt.Topics[1].Big().Uint64()
			destId := evt.Topics[2].Big().Uint64()
			depositNone := evt.Topics[3].Big().Uint64()

			if m.Source == msg.ChainId(sourceId) &&
				m.Destination == msg.ChainId(destId) &&
				m.DepositNonce == uint32(depositNone) {
				eventSubscription.sub.Unsubscribe()
				w.executeProposal(m, handler, data)
				return
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
	w.log.Info("Executing proposal", "to", w.conn.cfg.contract)

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
