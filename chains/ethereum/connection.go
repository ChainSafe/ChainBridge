package ethereum

import (
	"context"
	"math/big"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/crypto"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	eth "github.com/ethereum/go-ethereum"

	"github.com/ChainSafe/log15"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ chains.Connection = &Connection{}

type Connection struct {
	ctx      context.Context
	endpoint string
	home     ethcommon.Address
	away     ethcommon.Address
	conn     *ethclient.Client
	signer   ethtypes.Signer
	kp       crypto.Keypair
}

type ConnectionConfig struct {
	Ctx      context.Context
	Endpoint string
	Home     string
	Away     string
	Keypair  crypto.Keypair
	Signer   ethtypes.Signer
}

func NewConnection(cfg *ConnectionConfig) *Connection {
	return &Connection{
		ctx:      cfg.Ctx,
		endpoint: cfg.Endpoint,
		home:     ethcommon.HexToAddress(cfg.Home),
		away:     ethcommon.HexToAddress(cfg.Away),
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
func (c *Connection) subscribeToEvent(query eth.FilterQuery) (*Subscription, error) {
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

// SubmitTx submits a transaction to the chain
// It assumes the input data is an ethtypes.Transaction, marshalled as JSON
func (c *Connection) SubmitTx(data []byte) error {
	tx := &ethtypes.Transaction{}
	err := tx.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	log15.Debug("Submitting new tx", "to", tx.To(), "nonce", tx.Nonce(), "value", tx.Value(),
		"gasLimit", tx.Gas(), "gasPrice", tx.GasPrice(), "calldata", tx.Data())

	signedTx, err := ethtypes.SignTx(tx, c.signer, c.kp.Private().(*secp256k1.PrivateKey).Key())
	if err != nil {
		log15.Trace("Signing tx failed", "err", err)
		return err
	}

	return c.conn.SendTransaction(c.ctx, signedTx)
}

// NonceAt returns the nonce of the given account and the given block
func (c *Connection) NonceAt(account [20]byte, blockNum *big.Int) (uint64, error) {
	return c.conn.NonceAt(c.ctx, ethcommon.Address(account), blockNum)
}

// LatestBlock returns the latest block
func (c *Connection) LatestBlock() (*ethtypes.Block, error) {
	return c.conn.BlockByNumber(c.ctx, nil)
}
