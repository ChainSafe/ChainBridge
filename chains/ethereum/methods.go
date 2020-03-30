// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
)

const StoreMethod = "store"
const CreateDepositProposalMethod = "createDepositProposal"
const VoteDepositProposalMethod = "voteDepositProposal"
const ExecuteDepositMethod = "executeDepositProposal"

func (w *Writer) depositAsset(m msg.Message) bool {

	w.log.Info("Handling DepositAsset message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	//TODO: Should this be metadata?
	_, err = w.bridgeContract.BridgeRaw.Transact(opts, StoreMethod, hash(m.Metadata))

	if err != nil {
		log15.Error("Failed to submit depositAsset transaction", "err", err)
		return false
	}
	return true
}

func (w *Writer) createDepositProposal(m msg.Message) bool {
	w.log.Info("Handling CreateDepositProposal message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}
	log15.Info("opts", "from", opts.From.String())

	hash := hash(m.Metadata)
	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		CreateDepositProposalMethod,
		big.NewInt(int64(m.Source)),
		u32toBigInt(m.DepositNonce),
		hash,
	)

	if err != nil {
		w.log.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	w.log.Info("Succesfully created deposit!", "chain", m.Source, "deposit_id", m.DepositNonce)
	return true
}

func (w *Writer) voteDepositProposal(m msg.Message) bool {
	w.log.Info("Handling VoteDepositProposal message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	vote := uint8(1)
	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		VoteDepositProposalMethod,
		big.NewInt(int64(m.Source)),
		u32toBigInt(m.DepositNonce),
		vote,
	)

	if err != nil {
		w.log.Error("Failed to submit vote!", "chain", m.Source, "deposit_id", m.DepositNonce, "err", err)
		return false
	}
	w.log.Info("Succesfully voted!", "chain", m.Source, "deposit_id", m.DepositNonce, "Vote", vote)
	return true
}

func (w *Writer) executeDeposit(m msg.Message) bool {
	w.log.Info("Handling ExecuteDeposit message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.log.Error("Failed to build transaction opts", "err", err)
		return false
	}

	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		ExecuteDepositMethod,
		big.NewInt(int64(m.Source)),
		u32toBigInt(m.DepositNonce),
		byteSliceTo32Bytes(m.To),
		m.Metadata,
	)

	if err != nil {
		w.log.Error("Failed to submit executeDeposit transaction", "err", err)
		return false
	}
	return true
}
