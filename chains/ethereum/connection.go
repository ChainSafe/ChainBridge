package ethereum

import (
	"context"
	"math/big"

	"ChainBridgeV2/core"
	//"ChainBridgeV2/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// buildQuery constructs a query for the contract by hashing sig to get the event topic
func (c *Connection) buildQuery(contract common.Address, sig EventSig) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		// BlockHash: nil,
		FromBlock: nil,
		// ToBlock:   nil,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

// subscribeToEvent registers an rpc subscription for the event with the signature sig for contract at address
func (c *Connection) subscribeToEvent(addr common.Address, sig EventSig) (chan ethtypes.Log, ethereum.Subscription, error) {
	query := c.buildQuery(addr, sig)
	ch := make(chan ethtypes.Log)
	sub, err := c.conn.SubscribeFilterLogs(c.ctx, query, ch)
	if err != nil {
		close(ch)
		return nil, nil, err
	}
	return ch, sub, nil
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
