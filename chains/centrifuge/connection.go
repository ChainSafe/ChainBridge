package centrifuge

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
)

var _ core.Connection = &Connection{}

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func InitializeChain(id msg.ChainId, endpoint string, home, away []byte) *core.Chain {
	c := core.NewChain(id, endpoint, home, away)
	c.SetConnection(NewConnection())
	c.SetListener(NewListener(c.Connection()))
	c.SetPusher(NewPusher(c.Connection()))
	return c
}

func (c *Connection) Connect() error {
	return nil
}

func (c *Connection) SubmitTx(data []byte) error {
	panic("not implemented")
	return nil
}
