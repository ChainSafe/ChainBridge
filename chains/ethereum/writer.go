package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	//"github.com/ChainSafe/ChainBridgeV2/common"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	//ethcommon "github.com/ethereum/go-ethereum/common"
	//ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	cfg  Config
	conn *Connection
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

// ResolveMessage handles any given message based on type
// Note: We are currently panicking here, we should develop a better method for handling failures (possibly a queue?)
func (w *Writer) ResolveMessage(m msg.Message) {
	log15.Trace("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	// TODO: make this configurable
	gasLimit := big.NewInt(6721975)
	gasPrice := big.NewInt(20000000000)

	if m.Type == msg.DepositAssetType {
		log15.Info("Handling DepositAsset message", "to", w.conn.cfg.receiver, "msgdata", m.Data)

		method := "store"
		hash := [32]byte{}
		copy(hash[:], m.Data)

		opts, err := w.conn.newTransactOpts(big.NewInt(0), gasLimit, gasPrice)
		if err != nil {
			log15.Error("Failed to build transaction opts", "err", err)
			return
		}

		_, err = w.conn.Transact(opts, method, hash)
		if err != nil {
			log15.Error("Failed to submit transaction", "err", err)
		}
	} else if m.Type == msg.CreateDepositProposalType {
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

		_, err = w.conn.Transact(opts, method, hash, depositId, originChain)
		if err != nil {
			log15.Error("Failed to submit transaction", "err", err)
		}
	} else if m.Type == msg.VoteDepositProposalType {
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

		_, err = w.conn.Transact(opts, method, depositId, originChain, vote)
		if err != nil {
			log15.Error("Failed to submit transaction", "err", err)
		}
	} else if m.Type == msg.ExecuteDepositType {

	} else {
		panic("not implemented")
	}
}

func (w *Writer) Stop() error {
	return nil
}