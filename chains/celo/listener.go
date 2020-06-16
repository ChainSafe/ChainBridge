// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Connection = &connection.Connection{}

type Connection interface {
	Connect() error
	Close()
	Client() *ethclient.Client
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
	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}
	return tx.Hash()
}

func (l *listener) getBlockTransactionsByHash(hash common.Hash) (tx types.Transactions, txRoot common.Hash) {
	block, err := l.conn.Client().BlockByHash(context.Background(), hash)
	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}
	return block.Transactions(), block.Root()
}
