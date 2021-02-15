// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package client

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var BlockRetryInterval = time.Second * 5

type Client struct {
	*ethclient.Client
	endpoint      string
	http          bool
	kp            *secp256k1.Keypair
	gasLimit      *big.Int
	maxGasPrice   *big.Int
	gasMultiplier *big.Float
	opts          *bind.TransactOpts
	callOpts      *bind.CallOpts
	nonce         uint64
	optsLock      sync.Mutex
	log           log15.Logger
	stop          chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewClient(endpoint string, http bool, kp *secp256k1.Keypair, log log15.Logger, gasLimit, gasPrice *big.Int, gasMultiplier *big.Float) *Client {
	return &Client{
		endpoint:      endpoint,
		http:          http,
		kp:            kp,
		gasLimit:      gasLimit,
		maxGasPrice:   gasPrice,
		gasMultiplier: gasMultiplier,
		log:           log,
		stop:          make(chan int),
	}
}

// Connect starts the ethereum WS connection
func (c *Client) Connect() error {
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
	c.Client = ethclient.NewClient(rpcClient)

	// Construct tx opts, call opts, and nonce mechanism
	opts, _, err := c.newTransactOpts(big.NewInt(0), c.gasLimit, c.maxGasPrice)
	if err != nil {
		return err
	}
	c.opts = opts
	c.nonce = 0
	c.callOpts = &bind.CallOpts{From: c.kp.CommonAddress()}
	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Client) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error) {
	privateKey := c.kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.Client.PendingNonceAt(context.Background(), address)
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

func (c *Client) Keypair() *secp256k1.Keypair {
	return c.kp
}

func (c *Client) Opts() *bind.TransactOpts {
	return c.opts
}

func (c *Client) CallOpts() *bind.CallOpts {
	return c.callOpts
}

func (c *Client) SafeEstimateGas(ctx context.Context) (*big.Int, error) {
	suggestedGasPrice, err := c.Client.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}

	gasPrice := multiplyGasPrice(suggestedGasPrice, c.gasMultiplier)

	// Check we aren't exceeding our limit
	if gasPrice.Cmp(c.maxGasPrice) == 1 {
		return c.maxGasPrice, nil
	} else {
		return gasPrice, nil
	}
}

func multiplyGasPrice(gasEstimate *big.Int, gasMultiplier *big.Float) *big.Int {
	gasEstimateFloat := new(big.Float).SetInt(gasEstimate)
	result := gasEstimateFloat.Mul(gasEstimateFloat, gasMultiplier)
	gasPrice := new(big.Int)
	result.Int(gasPrice)
	return gasPrice
}

// LockAndUpdateOpts acquires a lock on the opts before updating the nonce
// and gas price.
func (c *Client) LockAndUpdateOpts() error {
	c.optsLock.Lock()

	gasPrice, err := c.SafeEstimateGas(context.TODO())
	if err != nil {
		return err
	}
	c.opts.GasPrice = gasPrice

	nonce, err := c.Client.PendingNonceAt(context.Background(), c.opts.From)
	if err != nil {
		c.optsLock.Unlock()
		return err
	}
	c.opts.Nonce.SetUint64(nonce)
	return nil
}

func (c *Client) UnlockOpts() {
	c.optsLock.Unlock()
}

// LatestBlock returns the latest block from the current chain
func (c *Client) LatestBlock() (*big.Int, error) {
	header, err := c.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Client) EnsureHasBytecode(addr ethcommon.Address) error {
	code, err := c.Client.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

// WaitForBlock will poll for the block number until the current block is equal or greater.
// If delay is provided it will wait until currBlock - delay = targetBlock
func (c *Client) WaitForBlock(targetBlock *big.Int, delay *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.LatestBlock()
			if err != nil {
				return err
			}

			if delay != nil {
				currBlock.Sub(currBlock, delay)
			}

			// Equal or greater than target
			if currBlock.Cmp(targetBlock) >= 0 {
				return nil
			}
			c.log.Trace("Block not ready, waiting", "target", targetBlock, "current", currBlock, "delay", delay)
			time.Sleep(BlockRetryInterval)
			continue
		}
	}
}

// Close terminates the client connection and stops any running routines
func (c *Client) Close() {
	if c.Client != nil {
		c.Client.Close()
	}
	close(c.stop)
}
