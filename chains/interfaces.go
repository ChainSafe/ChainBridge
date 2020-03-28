// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package chains

import (
	msg "github.com/ChainSafe/ChainBridge/message"
)

type Router interface {
	Send(message msg.Message) error
}

type Writer interface {
	ResolveMessage(message msg.Message) bool
}
