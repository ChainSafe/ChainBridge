// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"github.com/ChainSafe/chainbridge-utils/msg"
)

// need to figure out what function to call here
func (l *listener) handleErc721PutDataEvent(userAddress string, key string) (msg.Message, error) {
	l.log.Info("Handling nonfungible put data event")

	return msg.NewNonFungibleTransfer(l.cfg.id, userAddress, key), nil // need to alter message type
}
