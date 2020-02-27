package ethereum

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
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
		log15.Error("Failed to submit transaction", "err", err)
		return false
	}
	return true
}

func (w *Writer) createDepositProposal(m msg.Message) bool {
	log15.Info("Handling CreateDepositProposal message", "to", w.conn.cfg.receiver)

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	_, err = w.receiverContract.Transact(
		opts,
		CreateDepositProposalMethod,
		keccakHash(m.Metadata),
		u32toBigInt(m.DepositId),
		m.Source.Big())

	if err != nil {
		log15.Error("Failed to submit transaction", "err", err)
		return false
	}
	return true
}

func (w *Writer) voteDepositProposal(m msg.Message) bool {
	log15.Info("Handling VoteDepositProposal message", "to", w.conn.cfg.receiver)

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return false
	}

	_, err = w.receiverContract.Transact(
		opts,
		VoteDepositProposalMethod,
		u32toBigInt(m.DepositId),
		m.Source.Big(),
		uint8(1),
	)
	if err != nil {
		log15.Error("Failed to submit transaction", "err", err)
		return false
	}
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
		log15.Error("Failed to submit transaction", "err", err)
		return false
	}
	return true
}
