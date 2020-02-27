// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package centrifuge

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

var _ chains.Writer = &Writer{}

type Writer struct {
	conn chains.Connection
}

func NewWriter(conn chains.Connection) *Writer {
	return &Writer{
		conn: conn,
	}
}

func (w *Writer) Start() error {
	return nil
}

func (w *Writer) ResolveMessage(msg msg.Message) bool {
	panic("not implemented")
}

func (w *Writer) Stop() error {
	return nil
}
