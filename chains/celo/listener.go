// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"time"

	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Connection = &connection.Connection{}

type Connection interface {
	Connect() error
	Close()
	Client() *ethclient.Client
	Opts() *bind.TransactOpts
}

var ExpectedBlockTime = time.Second

type listener struct {
	conn Connection
}

// WaitForTx will query the chain at ExpectedBlockTime intervals, until a receipt is returned.
// Returns an error if the tx failed.
func WaitForTx(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	retry := 10
	for retry > 0 {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			retry--
			time.Sleep(ExpectedBlockTime)
			continue
		}

		if receipt.Status != 1 {
			return nil, fmt.Errorf("transaction failed on chain")
		}
		return receipt, nil
	}
	return nil, fmt.Errorf("transaction after retries failed")
}

func NewListener(conn Connection) *listener {
	return &listener{conn: conn}
}

func (l *listener) start() error {
	err := l.conn.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (l *listener) close() {
	l.conn.Close()
}

func (l *listener) getTransactionBlockHash(hash common.Hash) (blockHash common.Hash) {
	tx, _, err := l.conn.Client().TransactionByHash(context.Background(), hash)

	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}

	receipt, err := WaitForTx(l.conn.Client(), tx)
	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}

	return receipt.BlockHash
}

func (l *listener) getBlockTransactionsByHash(hash common.Hash) (tx []common.Hash, txRoot common.Hash) {
	time.Sleep(30 * time.Second)
	block, err := l.conn.Client().BlockByHash(context.Background(), hash)
	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}
	var transactionHashes []common.Hash

	transactions := block.Transactions()
	for _, transaction := range transactions {
		transactionHashes = append(transactionHashes, transaction.Hash())
	}

	return transactionHashes, block.Root() //should only return hash of transactions or all transactions
}
