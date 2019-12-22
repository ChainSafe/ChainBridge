package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
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
	panic("not implemented")
}

func (w *Writer) ResolveMessage(msg msg.Message) {
	panic("not implemented")
}

func (w *Writer) Stop() error {
	panic("not implemented")
}
