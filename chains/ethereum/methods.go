// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
)

const StoreMethod = "store"
const CreateDepositProposalMethod = "createDepositProposal"
const VoteDepositProposalMethod = "voteDepositProposal"
const ExecuteDepositMethod = "executeDepositProposal"

func (w *Writer) depositAsset(m msg.Message) bool {

	w.cfg.errorLog.Info("Handling DepositAsset message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.cfg.errorLog.Error("Failed to build transaction opts", "err", err)
		return false
	}

	//TODO: Should this be metadata?
	_, err = w.bridgeContract.BridgeRaw.Transact(opts, StoreMethod, hash(m.Metadata))

	if err != nil {
		w.cfg.errorLog.Error("Failed to submit depositASset transaction", "err", err)
		return false
	}
	return true
}

func (w *Writer) createDepositProposal(m msg.Message) bool {
	w.cfg.errorLog.Info("Handling CreateDepositProposal message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.cfg.errorLog.Error("Failed to build transaction opts", "err", err)
		return false
	}
	log15.Info("opts", "from", opts.From.String())

	hash := hash(m.Metadata)
	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		CreateDepositProposalMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositNonce),
		hash,
	)

	if err != nil {
		w.cfg.errorLog.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	w.cfg.errorLog.Info("Succesfully created deposit!", "chain", m.Source, "deposit_id", m.DepositNonce)
	return true
}

func (w *Writer) voteDepositProposal(m msg.Message) bool {
	w.cfg.errorLog.Info("Handling VoteDepositProposal message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.cfg.errorLog.Error("Failed to build transaction opts", "err", err)
		return false
	}

	vote := uint8(1)
	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		VoteDepositProposalMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositNonce),
		vote,
	)

	if err != nil {
		w.cfg.errorLog.Error("Failed to submit vote!", "chain", m.Source, "deposit_id", m.DepositNonce, "err", err)
		return false
	}
	w.cfg.errorLog.Info("Succesfully voted!", "chain", m.Source, "deposit_id", m.DepositNonce, "Vote", vote)
	return true
}

func (w *Writer) executeDeposit(m msg.Message) bool {
	w.cfg.errorLog.Info("Handling ExecuteDeposit message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		w.cfg.errorLog.Error("Failed to build transaction opts", "err", err)
		return false
	}

	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		ExecuteDepositMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositNonce),
		byteSliceTo32Bytes(m.To),
		m.Metadata,
	)

	if err != nil {
		w.cfg.errorLog.Error("Failed to submit executeDeposit transaction", "err", err)
		return false
	}
	return true
}
