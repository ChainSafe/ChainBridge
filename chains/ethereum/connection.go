// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"sync"

	"github.com/ChainSafe/ChainBridgeV2/contracts/Bridge"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
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
}

func NewConnection(cfg *Config, kp *secp256k1.Keypair) *Connection {
	signer := ethtypes.HomesteadSigner{}
	return &Connection{
		ctx: context.Background(),
		cfg: *cfg,
		// TODO: add network to use to config
		signer:    signer,
		kp:        kp,
		nonceLock: sync.Mutex{},
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	log15.Info("Connecting to ethereum chain...", "chain", c.cfg.name, "url", c.cfg.endpoint)
	rpcClient, err := rpc.DialWebsocket(c.ctx, c.cfg.endpoint, "/ws")
	if err != nil {
		return err
	}

	c.conn = ethclient.NewClient(rpcClient)
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

	signedTx, err := ethtypes.SignTx(tx, c.signer, c.kp.PrivateKey())
	if err != nil {
		log15.Trace("Signing tx failed", "err", err)
		return err
	}

	return c.conn.SendTransaction(c.ctx, signedTx)
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
func (c *Connection) getByteCode(account ethcommon.Address) ([]byte, error) {
	return c.conn.CodeAt(c.ctx, account, nil)
}

func (c *Connection) checkBridgeContract(contract ethcommon.Address) error {
	byteCode, err := c.getByteCode(contract)
	if err != nil {
		return err
	}

	if len(byteCode) == 0 || Bridge.RuntimeBytecode != "0x"+hex.EncodeToString(byteCode) {
		return errors.New("Bridge Bytecode does not match.")
	}

	return nil
}
