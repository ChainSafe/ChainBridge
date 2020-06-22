// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"context"
	"fmt"
	"time"

	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (l *listener) getTransactionBlockHash(hash common.Hash) (blockHash common.Hash, err error) {

	receipt, err := l.conn.Client().TransactionReceipt(context.Background(), hash)
	if err != nil {
		return hash, fmt.Errorf("unable to get BlockHash: %s", err)
	}
	return receipt.BlockHash, nil
}

func (l *listener) getBlockTransactionsByHash(hash common.Hash) (txHashes []common.Hash, txRoot common.Hash, err error) {
	block, err := l.conn.Client().BlockByHash(context.Background(), hash)
	if err != nil {
		return []common.Hash{}, common.Hash{}, fmt.Errorf("unable to get BlockHash: %s", err)
	}

	var transactionHashes []common.Hash

	transactions := block.Transactions()
	for _, transaction := range transactions {
		transactionHashes = append(transactionHashes, transaction.Hash())
	}

	return transactionHashes, block.Root(), nil
}
