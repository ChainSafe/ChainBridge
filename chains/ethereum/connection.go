package ethereum

import (
	"context"
	"math/big"

	"ChainBridgeV2/core"
	//"ChainBridgeV2/types"

	//eth "github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ core.Connection = &Connection{}

type Connection struct {
	ctx      context.Context
	endpoint string
	conn     *ethclient.Client
	// rpcConn  *rpc.Client

	// TODO: keystore
}

func NewConnection(ctx context.Context, endpoint string) *Connection {
	return &Connection{
		ctx:      ctx,
		endpoint: endpoint,
	}
}

func (c *Connection) Connect() error {
	rpcClient, err := rpc.DialHTTP(c.endpoint)
	if err != nil {
		return err
	}

	c.conn = ethclient.NewClient(rpcClient)
	return nil
}

func (c *Connection) Close() {
	c.conn.Close()
}

func (c *Connection) NetworkId() (*big.Int, error) {
	return c.conn.NetworkID(c.ctx)
}

func (c *Connection) SubmitTx(data []byte) error {
	tx := &ethtypes.Transaction{}
	err := tx.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	// TODO: need to set up keystore before a tx can be sent and accepted
	// signedTx, err := ethtypes.SignTx(tx, signer, priv)
	// if err != nil {
	// 	return err
	// }

	return c.conn.SendTransaction(c.ctx, tx)
}
