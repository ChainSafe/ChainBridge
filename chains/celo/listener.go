// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

type listener struct {
	conn Connection
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
	//utils.WaitForTx(l.conn.Client(), tx)
	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}

	receipt, err := l.conn.Client().TransactionReceipt(context.Background(), tx.Hash())
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
