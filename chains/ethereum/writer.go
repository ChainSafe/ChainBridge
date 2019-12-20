package ethereum

import (
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/common"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	conn *Connection
}

func NewWriter(conn *Connection) *Writer {
	return &Writer{
		conn: conn,
	}
}

func (w *Writer) Start() error {
	log15.Info("Starting ethereum writer...")
	log15.Warn("Writer.Start() not fully implemented")
	return nil
}

func (w *Writer) ResolveMessage(m msg.Message) error {
	// TODO: make sure message is bound for this chain
	// if m.Destination != w.conn.Id()

	var tx *ethtypes.Transaction

	if m.Type == msg.AssetTransferType {
		log15.Info("sending asset transfer...", "to", w.conn.away, "msgdata", m.Data)
		currBlock, err := w.conn.LatestBlock()
		if err != nil {
			return err
		}

		address := ethcommon.HexToAddress(w.conn.kp.Public().Address())

		nonce, err := w.conn.NonceAt(address, currBlock.Number())
		if err != nil {
			return err
		}

		id := common.FunctionId("store(bytes32)")
		calldata := append(id, m.Data...)

		tx = ethtypes.NewTransaction(
			nonce,
			w.conn.away,
			big.NewInt(0),  // TODO: value?
			1000000,        // TODO: gasLimit
			big.NewInt(10), // TODO: gasPrice
			calldata,
		)
	} else {
		panic("not implemented")
	}

	data, err := tx.MarshalJSON()
	if err != nil {
		return err
	}

	err = w.conn.SubmitTx(data)
	if err != nil {
		return err
	}
	return nil
}

func (w *Writer) Stop() error {
	panic("not implemented")
}
