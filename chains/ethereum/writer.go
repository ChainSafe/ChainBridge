package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	conn chains.Connection
}

func NewWriter(conn chains.Connection) *Writer {
	return &Writer{
		conn: conn,
	}
}

func (w *Writer) Start() error {
	log15.Info("Starting ethereum writer...")
	log15.Warn("Writer.Start() not fully implemented")
	return nil
}

func (w *Writer) ResolveMessage(msg msg.Message) {
	panic("not implemented")
}

func (w *Writer) Stop() error {
	panic("not implemented")
}
