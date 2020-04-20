// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	"github.com/ChainSafe/log15"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Nonce is a struct that wraps the Nonce with a mutex lock
// this struct was implemented to prevent race conditions where
// two transactions try to transact at the same time and recieve
// the same nonce, causing one to be rejected.
type Nonce struct {
	nonce uint64
	lock  *sync.Mutex
}

type Connection struct {
	cfg       Config
	ctx       context.Context
	conn      *ethclient.Client
	signer    ethtypes.Signer
	kp        *secp256k1.Keypair
	nonceLock sync.Mutex
	log       log15.Logger
}

func NewConnection(cfg *Config, kp *secp256k1.Keypair, log log15.Logger) *Connection {
	return &Connection{
		ctx:       context.Background(),
		cfg:       *cfg,
		kp:        kp,
		nonceLock: sync.Mutex{},
		log:       log,
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "url", c.cfg.endpoint, "startBlock", c.cfg.startBlock.String())
	var rpcClient *rpc.Client
	var err error
	if c.cfg.http {
		rpcClient, err = rpc.DialHTTP(c.cfg.endpoint)
	} else {
		rpcClient, err = rpc.DialWebsocket(c.ctx, c.cfg.endpoint, "/ws")
	}
	if err != nil {
		return err
	}

	c.conn = ethclient.NewClient(rpcClient)

	chainId, err := c.conn.ChainID(c.ctx)
	if err != nil {
		return err
	}
	c.signer = ethtypes.NewEIP155Signer(chainId)
	return nil
}

// Close stops the WS connection
func (c *Connection) Close() {
	c.conn.Close()
}

func (c *Connection) NetworkId() (*big.Int, error) {
	return c.conn.NetworkID(c.ctx)
}

// subscribeToEvent registers an rpc subscription for the event with the signature sig for contract at address
func (c *Connection) subscribeToEvent(query eth.FilterQuery) (*ActiveSubscription, error) {
	ch := make(chan ethtypes.Log)
	sub, err := c.conn.SubscribeFilterLogs(c.ctx, query, ch)
	if err != nil {
		close(ch)
		return nil, err
	}
	return &ActiveSubscription{
		ch:  ch,
		sub: sub,
	}, nil
}

// PendingNonceAt returns the pending nonce of the given account and the given block
func (c *Connection) PendingNonceAt(account [20]byte) (*Nonce, error) {
	c.nonceLock.Lock()
	nonce, err := c.conn.PendingNonceAt(c.ctx, ethcommon.Address(account))
	if err != nil {
		c.nonceLock.Unlock()
		return nil, err
	}

	return &Nonce{
		nonce,
		&c.nonceLock,
	}, nil
}

// NonceAt returns the nonce of the given account and the given block
func (c *Connection) NonceAt(account [20]byte, blockNum *big.Int) (uint64, error) {
	return c.conn.NonceAt(c.ctx, ethcommon.Address(account), blockNum)
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*ethtypes.Block, error) {
	return c.conn.BlockByNumber(c.ctx, nil)
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, *Nonce, error) {
	privateKey := c.kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.PendingNonceAt(address)
	if err != nil {
		return nil, nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce.nonce))
	auth.Value = big.NewInt(0)               // in wei
	auth.GasLimit = uint64(gasLimit.Int64()) // in units
	auth.GasPrice = gasPrice
	auth.Context = c.ctx

	return auth, nonce, nil
}

//getByteCode grabs the bytecode of the contract, to help determine if a contract is deployed
func (c *Connection) ensureHasBytecode(addr ethcommon.Address) error {
	code, err := c.conn.CodeAt(c.ctx, addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}
