package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/types"
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

// Subscribe creates a subscription to all events
func (c *Connection) Subscribe() (*state.StorageSubscription, error) {
	meta, err := c.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	key, err := types.CreateStorageKey(meta, "System", "Events", nil, nil)
	if err != nil {
		return nil, err
	}

	sub, err := c.api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (c *Connection) Close() {
	// Do nothing for now
}
