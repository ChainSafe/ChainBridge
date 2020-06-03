// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package evm

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

type connection interface { //nolint:deadcode,unused
	Connect() error
	Close()
	NetworkId() (*big.Int, error)
	LatestBlock() (*big.Int, error)
	EnsureHasBytecode(addr ethcommon.Address) error
	NewTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error)
	WaitForBlock(block *big.Int) error
}

type Connection struct {
	Cfg       Config
	Ctx       context.Context
	Conn      *ethclient.Client
	Signer    ethtypes.Signer
	Kp        *secp256k1.Keypair
	NonceLock sync.Mutex
	log       log15.Logger
	Stop      <-chan int // All routines should exit when this channel is closed
}

func NewConnection(cfg *Config, kp *secp256k1.Keypair, log log15.Logger, stop <-chan int) *Connection {
	return &Connection{
		Ctx:       context.Background(),
		Cfg:       *cfg,
		Kp:        kp,
		NonceLock: sync.Mutex{},
		log:       log,
		Stop:      stop,
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "url", c.Cfg.Endpoint, "startBlock", c.Cfg.StartBlock.String())
	var rpcClient *rpc.Client
	var err error
	if c.Cfg.http {
		rpcClient, err = rpc.DialHTTP(c.Cfg.Endpoint)
	} else {
		rpcClient, err = rpc.DialWebsocket(c.Ctx, c.Cfg.Endpoint, "/ws")
	}
	if err != nil {
		return err
	}

	c.Conn = ethclient.NewClient(rpcClient)

	chainId, err := c.Conn.ChainID(c.Ctx)
	if err != nil {
		return err
	}
	c.Signer = ethtypes.NewEIP155Signer(chainId)
	return nil
}

// Close stops the WS connection
func (c *Connection) Close() {
	c.Conn.Close()
}

func (c *Connection) NetworkId() (*big.Int, error) {
	return c.Conn.NetworkID(c.Ctx)
}

// latestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.Conn.HeaderByNumber(c.Ctx, nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) NewTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error) {
	privateKey := c.Kp.PrivateKey()
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.Conn.PendingNonceAt(c.Ctx, address)
	if err != nil {
		return nil, 0, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(gasLimit.Int64())
	auth.GasPrice = gasPrice
	auth.Context = c.Ctx

	return auth, nonce, nil
}

//getByteCode grabs the bytecode of the contract, to help determine if a contract is deployed
func (c *Connection) EnsureHasBytecode(addr ethcommon.Address) error {
	code, err := c.Conn.CodeAt(c.Ctx, addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

// waitForBlock will poll for the block number until the current block is equal or greater than
func (c *Connection) WaitForBlock(block *big.Int) error {
	for {
		select {
		case <-c.Stop:
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

			time.Sleep(BlockRetryInterval)
			continue
		}
	}
}
