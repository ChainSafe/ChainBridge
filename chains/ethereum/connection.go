package ethereum

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
)

var _ core.Connection = &Connection{}

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func InitializeChain(id msg.ChainId, home, away []byte) *core.Chain {
	c := core.NewChain(id, home, away)
	c.SetConnection(NewConnection())
	c.SetListener(NewListener(c.Connection()))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}

func (c *Connection) Connect() error {
	return nil
}

func (c *Connection) SubmitTx(data []byte) error {
	panic("not implemented")
	return nil
}
