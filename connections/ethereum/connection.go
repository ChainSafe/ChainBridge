// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var BlockRetryInterval = time.Second * 5

type Connection struct {
	endpoint string
	http     bool
	kp       *secp256k1.Keypair
	gasLimit *big.Int
	gasPrice *big.Int
	conn     *ethclient.Client
	// signer    ethtypes.Signer
	opts      *bind.TransactOpts
	callOpts  *bind.CallOpts
	nonce     uint64
	nonceLock sync.Mutex
	log       log15.Logger
	stop      chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, http bool, kp *secp256k1.Keypair, log log15.Logger, gasLimit, gasPrice *big.Int) *Connection {
	return &Connection{
		endpoint: endpoint,
		http:     http,
		kp:       kp,
		gasLimit: gasLimit,
		gasPrice: gasPrice,
		log:      log,
		stop:     make(chan int),
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "url", c.endpoint)
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	if c.http {
		rpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = rpc.DialWebsocket(context.Background(), c.endpoint, "/ws")
	}
	if err != nil {
		return err
	}
	c.conn = ethclient.NewClient(rpcClient)

	// Construct tx opts, call opts, and nonce mechanism
	opts, _, err := c.newTransactOpts(big.NewInt(0), c.gasLimit, c.gasPrice)
	if err != nil {
		return err
	}
	c.opts = opts
	c.nonce = 0
	c.callOpts = &bind.CallOpts{From: c.kp.CommonAddress()}
	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error) {
	privateKey := c.kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, 0, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(gasLimit.Int64())
	auth.GasPrice = gasPrice
	auth.Context = context.Background()

	return auth, nonce, nil
}

func (c *Connection) Keypair() *secp256k1.Keypair {
	return c.kp
}

func (c *Connection) Client() *ethclient.Client {
	return c.conn
}

func (c *Connection) Opts() *bind.TransactOpts {
	return c.opts
}

func (c *Connection) CallOpts() *bind.CallOpts {
	return c.callOpts
}

func (c *Connection) LockAndUpdateNonce() error {
	c.nonceLock.Lock()
	nonce, err := c.conn.PendingNonceAt(context.Background(), c.opts.From)
	if err != nil {
		c.nonceLock.Unlock()
		return err
	}
	c.opts.Nonce.SetUint64(nonce)
	return nil
}

func (c *Connection) UnlockNonce() {
	c.nonceLock.Unlock()
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Connection) EnsureHasBytecode(addr ethcommon.Address) error {
	code, err := c.conn.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

// WaitForBlock will poll for the block number until the current block is equal or greater than
func (c *Connection) WaitForBlock(block *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.LatestBlock()
			if err != nil {
				return err
			}

			// Equal or greater than target
			if currBlock.Cmp(block) >= 0 {
				return nil
			}
			c.log.Trace("Block not ready, waiting", "target", block, "current", currBlock)
			time.Sleep(BlockRetryInterval)
			continue
		}
	}
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
	close(c.stop)
}
