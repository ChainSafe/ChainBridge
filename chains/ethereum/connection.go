package ethereum

import (
	"context"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"

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
	signer ethtypes.Signer
	kp     crypto.Keypair
}

type ConnectionConfig struct {
	Ctx      context.Context
	Endpoint string
	Keypair  crypto.Keypair
	Signer   ethtypes.Signer
}

func NewConnection(cfg *ConnectionConfig) *Connection {
	return &Connection{
		ctx:      cfg.Ctx,
		endpoint: cfg.Endpoint,
		kp:       cfg.Keypair,
		signer:   cfg.Signer,
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

	signedTx, err := ethtypes.SignTx(tx, c.signer, c.kp.Private().(*secp256k1.PrivateKey).Key())
	if err != nil {
		return err
	}

	return c.conn.SendTransaction(c.ctx, signedTx)
}
