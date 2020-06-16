// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
)

var _ Connection = &connection.Connection{}

type Connection interface {
	Connect() error
	Close()
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
