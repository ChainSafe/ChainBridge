package ethereum

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
)

var _ core.Pusher = &Pusher{}

type Pusher struct {
	conn core.Connection
}

func NewPusher(conn core.Connection) *Pusher {
	return &Pusher{
		conn: conn,
	}
}

func (p *Pusher) ResolveMessage(msg msg.Message) error {
	panic("not implemented")
	return nil
}
