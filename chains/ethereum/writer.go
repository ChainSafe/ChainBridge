package ethereum

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

func (p *Writer) ResolveMessage(msg msg.Message) {
	panic("not implemented")
}
