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

// TODO: make this configurable
var gasLimit = big.NewInt(6721975)
var gasPrice = big.NewInt(20000000000)

func (w *Writer) depositAsset(m msg.Message) bool {
	log15.Info("Handling DepositAsset message", "to", w.conn.cfg.receiver)

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	//TODO: Should this be metadata?
	_, err = w.receiverContract.Transact(opts, StoreMethod, keccakHash(m.Metadata))
	if err != nil {
		log15.Error("Failed to submit depositASset transaction", "err", err)
		return false
	}
	return true
}

func (w *Writer) createDepositProposal(m msg.Message) bool {
	log15.Trace("Handling CreateDepositProposal message", "to", w.conn.cfg.receiver)
	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	types := []string{"bytes"}
	values := []interface{}{m.Metadata}
	hash := solsha3.SoliditySHA3(types, values)

	var sizedHash [32]byte
	copy(sizedHash[:], hash)
	_, err = w.receiverContract.Transact(
		opts,
		CreateDepositProposalMethod,
		sizedHash,
		u32toBigInt(m.DepositId),
		m.Source.Big(),
	)

	if err != nil {
		log15.Error("Failed to submit createDepositProposal transaction", "err", err)
		return false
	}
	log15.Info("Succesfully created deposit!", "chain", m.Source, "deposit_id", m.DepositId)
	return true
}

func (w *Writer) voteDepositProposal(m msg.Message) bool {
	log15.Trace("Handling VoteDepositProposal message", "to", w.conn.cfg.receiver)

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	vote := uint8(0)

	_, err = w.receiverContract.Transact(
		opts,
		VoteDepositProposalMethod,
		m.Source.Big(),
		u32toBigInt(m.DepositId),
		vote,
	)
	if err != nil {
		log15.Error("Failed to submit vote!", "chain", m.Source, "deposit_id", m.DepositId, "err", err)
		return false
	}
	log15.Info("Succesfully voted!", "chain", m.Source, "deposit_id", m.DepositId, "Vote", vote)
	return true
}

func (w *Writer) executeDeposit(m msg.Message) bool {
	log15.Info("Handling ExecuteDeposit message", "to", w.conn.cfg.receiver)

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	_, err = w.receiverContract.Transact(
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
