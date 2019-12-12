package centrifuge

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
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
	panic("not implemented")
}

func (w *Writer) ResolveMessage(msg msg.Message) error {
	panic("not implemented")
}

func (w *Writer) Stop() error {
	panic("not implemented")
}
