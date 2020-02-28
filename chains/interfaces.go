// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package chains

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type Connection interface {
	Connect() error
	SubmitTx([]byte) error
	Close()
}

type Router interface {
	Send(message msg.Message) error
}

type Listener interface {
	RegisterEventHandler(string, EvtHandlerFn) error
}

type EvtHandlerFn func(interface{}) msg.Message

type Writer interface {
	ResolveMessage(message msg.Message) bool
}
