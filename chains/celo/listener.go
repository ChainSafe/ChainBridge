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

func (l *listener) getTransactionBlockHash(hash common.Hash) (tx *types.Transaction) {
	client := l.conn.Client()
	blockHash, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Errorf("unable to get BlockHash: %s", err)
	}
	return blockHash
}

func (l *listener) getBlockTransactionsByHash(hash common.Hash) *types.Block {
	client := l.conn.Client()
	block, err := client.BlockByHash(context.Background(), hash)
	if err != nil {
		return block
	}

}
