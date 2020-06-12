// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package celo

import (
	"testing"
)

func Test_Start(t *testing.T) {
	l := listener{}
	err := l.conn.Connection()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Stop(t *testing.T) {
	l := listener{}
	err := l.conn.Connection()
	if err != nil {
		t.Fatal(err)
	}
	l.conn.Stop()
}
