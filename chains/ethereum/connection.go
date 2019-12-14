package ethereum

import (
	"context"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// buildQuery constructs a query for the contract by hashing sig to get the event topic
func (c *Connection) buildQuery(contract common.Address, sig EventSig) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		// TODO: Might want current block
		FromBlock: nil,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

// subscribeToEvent registers an rpc subscription for the event with the signature sig for contract at address
func (c *Connection) subscribeToEvent(addr common.Address, sig EventSig) (*Subscription, error) {
	query := c.buildQuery(addr, sig)
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
