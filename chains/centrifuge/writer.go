package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/core"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
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

func (p *Writer) ResolveMessage(msg msg.Message) error {
	panic("not implemented")
}
