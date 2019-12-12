package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
)

var _ core.Writer = &Writer{}

type Writer struct {
	conn core.Connection
}

func NewWriter(conn core.Connection) *Writer {
	return &Writer{
		conn: conn,
	}
}

func (w *Writer) Start() error {
	log15.Info("Starting ethereum writer...")
	log15.Warn("Writer.Start() not fully implemented")
	return nil
}

func (w *Writer) ResolveMessage(msg msg.Message) error {
	panic("not implemented")
}

func (w *Writer) Stop() error {
	panic("not implemented")
}
