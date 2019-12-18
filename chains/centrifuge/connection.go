package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
)

var _ chains.Connection = &Connection{}

type Connection struct {
	api *gsrpc.SubstrateAPI
	url string
}

func NewConnection(url string) *Connection {
	return &Connection{url: url}
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
}

func (c *Connection) Close() {
	// TODO: Close API
}
