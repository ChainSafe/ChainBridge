package ethereum

import (
	"context"
	"math/big"

	"ChainBridgeV2/core"
	msg "ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ core.Connection = &Connection{}

type Connection struct {
	ctx      context.Context
	endpoint string
	conn     *ethclient.Client

	// TODO: keystore
}

func NewConnection(ctx context.Context, endpoint string) *Connection {
	return &Connection{
		ctx:      ctx,
		endpoint: endpoint,
	}
}

func InitializeChain(id msg.ChainId, endpoint string, home, away []byte) *core.Chain {
	ctx := context.Background()
	c := core.NewChain(id, endpoint, home, away)

	c.SetConnection(NewConnection(ctx, endpoint))
	c.SetListener(NewListener(c.Connection()))
	c.SetPusher(NewPusher(c.Connection()))
	return c
}

func (c *Connection) Connect() error {
	rpcClient, err := rpc.Dial(c.endpoint)
	if err != nil {
		return err
	}

	c.conn = ethclient.NewClient(rpcClient)
	return nil
}

func (c *Connection) NetworkId() (*big.Int, error) {
	return c.conn.NetworkID(c.ctx)
}

func (c *Connection) SubmitTx(data []byte) error {
	panic("not implemented")
	return nil
}
