package ethereum

import (
	"context"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	eth "github.com/ethereum/go-ethereum"

	"github.com/ChainSafe/log15"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ chains.Connection = &Connection{}

type Connection struct {
	ctx      context.Context
	endpoint string
	conn     *ethclient.Client
	signer   ethtypes.Signer
	kp       crypto.Keypair
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
	log15.Info("Connecting to ethereum...", "url", c.endpoint)
	rpcClient, err := rpc.DialWebsocket(c.ctx, c.endpoint, "/ws")
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

// subscribeToEvent registers an rpc subscription for the event with the signature sig for contract at address
func (c *Connection) subscribeToEvent(query eth.FilterQuery, sig EventSig) (*Subscription, error) {
	ch := make(chan ethtypes.Log)
	sub, err := c.conn.SubscribeFilterLogs(c.ctx, query, ch)
	if err != nil {
		close(ch)
		return nil, err
	}
	return &Subscription{
		ch:  ch,
		sub: sub,
	}, nil
}

func (c *Connection) SubmitTx(data []byte) error {
	log15.Debug("Submitting new tx", "data", data)
	tx := &ethtypes.Transaction{}
	err := tx.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	signedTx, err := ethtypes.SignTx(tx, c.signer, c.kp.Private().(*secp256k1.PrivateKey).Key())
	if err != nil {
		log15.Trace("Signing tx failed", "err", err)
		return err
	}

	return c.conn.SendTransaction(c.ctx, signedTx)
}
