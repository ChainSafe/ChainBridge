package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	cfg              Config
	conn             *Connection
	receiverContract ReceiverContract // instance of bound receiver contract
}

func NewWriter(conn *Connection, cfg *Config) *Writer {
	return &Writer{
		cfg:  *cfg,
		conn: conn,
	}
}

func (w *Writer) Start() error {
	log15.Debug("Starting ethereum writer...")
	log15.Warn("Writer.Start() not fully implemented")
	return nil
}

func (w *Writer) SetReceiverContract(rc ReceiverContract) {
	w.receiverContract = rc
}

func (w *Writer) ResolveMessage(message msg.Message) {
	if message.Type == msg.CreateDepositProposalType {
		w.CreateDepositProposal(message)
	} else if message.Type == msg.VoteDepositProposalType {
		w.VoteDepositProposal(message)
	} else if message.Type == msg.ExecuteDepositType {
		w.ExecuteDeposit(message)
	}
}

func (w *Writer) CreateDepositProposal(m msg.Message) {
	log15.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	// TODO: make this configurable
	gasLimit := big.NewInt(6721975)
	gasPrice := big.NewInt(20000000000)

	log15.Info("Handling CreateDepositProposal message", "to", w.conn.cfg.receiver, "msgdata", m.Data)

	method := "createDepositProposal"
	hash, depositId, originChain, err := m.DecodeCreateDepositProposalData()
	if err != nil {
		log15.Error("Failed to handle CreateDepositProposal message", "error", err)
		return
	}

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return
	}

	_, err = w.Transact(opts, method, hash, depositId, originChain)
	if err != nil {
		log15.Error("Failed to submit transaction", "err", err)
	}
}

func (w *Writer) VoteDepositProposal(m msg.Message) {
	log15.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	// TODO: make this configurable
	gasLimit := big.NewInt(6721975)
	gasPrice := big.NewInt(20000000000)

	log15.Info("Handling VoteDepositProposal message", "to", w.conn.cfg.receiver, "msgdata", m.Data)

	method := "voteDepositProposal"
	depositId, originChain, vote, err := m.DecodeVoteDepositProposalData()
	if err != nil {
		log15.Error("Failed to handle VoteDepositProposal message", "error", err)
		return
	}

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return
	}

	_, err = w.Transact(opts, method, depositId, originChain, vote)
	if err != nil {
		log15.Error("Failed to submit transaction", "err", err)
	}
}

func (w *Writer) ExecuteDeposit(m msg.Message) {
	log15.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	// TODO: make this configurable
	gasLimit := big.NewInt(6721975)
	gasPrice := big.NewInt(20000000000)

	log15.Info("Handling ExecuteDeposit message", "to", w.conn.cfg.receiver, "msgdata", m.Data)

	method := "executeDeposit"
	depositId, originChain, address, data, err := m.DecodeExecuteDepositData()
	if err != nil {
		log15.Error("Failed to handle VoteDepositProposal message", "error", err)
		return
	}

	opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		log15.Error("Failed to build transaction opts", "err", err)
		return
	}

	_, err = w.Transact(opts, method, depositId, originChain, address, data)
	if err != nil {
		log15.Error("Failed to submit transaction", "err", err)
	}
}

func (w *Writer) Stop() error {
	return nil
}

// Transact submits a transaction to the receiver contract intsance.
func (w *Writer) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*ethtypes.Transaction, error) {
	return w.receiverContract.Transact(opts, method, params...)
}
