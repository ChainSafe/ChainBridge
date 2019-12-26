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
	var tx *ethtypes.Transaction
	var calldata []byte

	if m.Type == msg.DepositAssetType {
		log15.Info("Handling Deposit Asset message", "to", w.conn.cfg.emitter, "msgdata", m.Data)
		id := common.FunctionId("store(bytes32)")
		calldata = append(id, m.Data...)
	} else {
		panic("not implemented")
	}

	currBlock, err := w.conn.LatestBlock()
	if err != nil {
		panic(err)
	}
	address := ethcommon.HexToAddress(w.conn.kp.Public().Address())

	nonce, err := w.conn.NonceAt(address, currBlock.Number())
	if err != nil {
		panic(err)
	}

	tx = ethtypes.NewTransaction(
		nonce,
		w.conn.cfg.receiver,
		big.NewInt(0),  // TODO: value?
		1000000,        // TODO: gasLimit
		big.NewInt(10), // TODO: gasPrice
		calldata,
	)

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
