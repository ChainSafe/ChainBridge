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

// ResolveMessage handles any given message based on type
// Note: We are currently panicking here, we should develop a better method for handling failures (possibly a queue?)
func (w *Writer) ResolveMessage(m msg.Message) {
	var tx *ethtypes.Transaction

	if m.Type == msg.AssetTransferType {
		log15.Info("sending asset transfer...", "to", w.conn.emitter, "msgdata", m.Data)
		currBlock, err := w.conn.LatestBlock()
		if err != nil {
			panic(err)
		}

		address := ethcommon.HexToAddress(w.conn.kp.Public().Address())

		nonce, err := w.conn.NonceAt(address, currBlock.Number())
		if err != nil {
			panic(err)
		}

		id := common.FunctionId("store(bytes32)")
		calldata := append(id, m.Data...)

		tx = ethtypes.NewTransaction(
			nonce,
			w.conn.receiver,
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
		panic(err)
	}

	err = w.conn.SubmitTx(data)
	if err != nil {
		panic(err)

	}
}

func (w *Writer) Stop() error {
	return nil
}
