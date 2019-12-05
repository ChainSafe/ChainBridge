package centrifuge

import (
	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
)

var _ core.Connection = &Connection{}

type Connection struct {
	api *gsrpc.SubstrateAPI
	url string
}

func NewConnection(url string) *Connection {
	return &Connection{}
}

func (c *Connection) Connect() error {
	api, err := gsrpc.NewSubstrateAPI(c.url)
	if err != nil {
		return err
	}
	c.api = api
	return nil
}

func (c *Connection) SubmitTx(data []byte) error {
	panic("not implemented")
	return nil
}

func InitializeChain(id msg.ChainId, endpoint string, home, away []byte) *core.Chain {
	c := core.NewChain(id, endpoint, home, away)
	c.SetConnection(NewConnection())
	c.SetListener(NewListener(c.Connection()))
	c.SetWriter(NewWriter(c.Connection()))
	return c
}