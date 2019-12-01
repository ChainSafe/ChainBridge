package centrifuge

import (
	"ChainBridgeV2/core"
)

var _ core.Connection = &Connection{}

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect() error {
	return nil
}

func (c *Connection) SubmitTx(data []byte) error {
	panic("not implemented")
	return nil
}
