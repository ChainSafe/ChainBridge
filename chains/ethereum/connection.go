package ethereum

import (
	"context"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethparams "github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ core.Connection = &Connection{}

type Connection struct {
	ctx      context.Context
	endpoint string
	conn     *ethclient.Client
	// rpcConn  *rpc.Client
	kp crypto.Keypair
}

func NewConnection(ctx context.Context, endpoint string, kp crypto.Keypair) *Connection {
	return &Connection{
		ctx:      ctx,
		endpoint: endpoint,
		kp:       kp,
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

	signer := ethtypes.MakeSigner(ethparams.RinkebyChainConfig, ethparams.RinkebyChainConfig.IstanbulBlock)
	signedTx, err := ethtypes.SignTx(tx, signer, c.kp.Private().(*secp256k1.PrivateKey).Key())
	if err != nil {
		return err
	}

	return c.conn.SendTransaction(c.ctx, signedTx)
}
