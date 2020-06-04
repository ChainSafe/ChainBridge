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
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Connection struct {
	cfg       Config
	ctx       context.Context
	conn      *ethclient.Client
	signer    ethtypes.Signer
	kp        *secp256k1.Keypair
	nonceLock sync.Mutex
	log       log15.Logger
	stop      <-chan int // All routines should exit when this channel is closed
}

func NewConnection(cfg *Config, kp *secp256k1.Keypair, log log15.Logger, stop <-chan int) *Connection {
	return &Connection{
		ctx:       context.Background(),
		cfg:       *cfg,
		kp:        kp,
		nonceLock: sync.Mutex{},
		log:       log,
		stop:      stop,
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

// latestBlock returns the latest block from the current chain
func (c *Connection) latestBlock() (*big.Int, error) {
	header, err := c.conn.HeaderByNumber(c.ctx, nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error) {
	privateKey := c.kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.conn.PendingNonceAt(c.ctx, address)
	if err != nil {
		return nil, 0, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(gasLimit.Int64())
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

// waitForBlock will poll for the block number until the current block is equal or greater than
func (c *Connection) waitForBlock(block *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.latestBlock()
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
