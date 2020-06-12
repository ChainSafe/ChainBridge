// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import ()

type Connection interface {
	Connection() error
	Stop()
}

type listener struct {
	conn Connection
}

func start(l *listener) error {
	err := l.conn.Connection()
	if err != nil {
		return err
	}
	return nil
}

func stop(l *listener) {
	l.conn.Stop()
}
