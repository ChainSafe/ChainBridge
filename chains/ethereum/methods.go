// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const StoreMethod = "store"
const CreateDepositProposalMethod = "createDepositProposal"
const VoteDepositProposalMethod = "voteDepositProposal"
const ExecuteDepositMethod = "executeDeposit"

func (w *Writer) depositAsset(m msg.Message) bool {

	log15.Info("Handling DepositAsset message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	//TODO: Should this be metadata?
	_, err = w.bridgeContract.BridgeRaw.Transact(opts, StoreMethod, keccakHash(m.Metadata))

	if err != nil {
		log15.Error("Failed to submit depositASset transaction", "err", err)
		return false
	}
	return true
}

func (w *Writer) createDepositProposal(m msg.Message) bool {
	log15.Info("Handling CreateDepositProposal message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	types := []string{"bytes"}
	values := []interface{}{m.Metadata}
	hash := solsha3.SoliditySHA3(types, values)

	var sizedHash [32]byte
	copy(sizedHash[:], hash)

	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		CreateDepositProposalMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositId),
		&sizedHash,
	)

	if err != nil {
		log15.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	log15.Info("Succesfully created deposit!", "chain", m.Source, "deposit_id", m.DepositId)
	return true
}

func (w *Writer) voteDepositProposal(m msg.Message) bool {
	log15.Info("Handling VoteDepositProposal message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	vote := uint8(0)

	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		VoteDepositProposalMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositId),
	)

	if err != nil {
		log15.Error("Failed to submit vote!", "chain", m.Source, "deposit_id", m.DepositId, "err", err)
		return false
	}
	log15.Info("Succesfully voted!", "chain", m.Source, "deposit_id", m.DepositId, "Vote", vote)
	return true
}

func (w *Writer) executeDeposit(m msg.Message) bool {
	log15.Info("Handling ExecuteDeposit message", "to", w.conn.cfg.contract)

	opts, nonce, err := w.conn.newTransactOpts(big.NewInt(0), w.gasLimit, w.gasPrice)
	defer nonce.lock.Unlock()
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	_, err = w.bridgeContract.BridgeRaw.Transact(
		opts,
		ExecuteDepositMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositId),
		byteSliceTo32Bytes(m.To),
		m.Metadata,
	)

	if err != nil {
		log15.Error("Failed to submit executeDeposit transaction", "err", err)
		return false
	}
	return true
}
