// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/ChainBridgeV2/router"
)

type Chain interface {
	Start() error // Start chain
	Stop() error  // Stop chain
	SetRouter(*router.Router)
	Id() msg.ChainId
}

type ChainConfig struct {
	Id       msg.ChainId        // ChainID
	Endpoint string             // url for rpc endpoint
	From     string             // address of key to use
	Keystore *keystore.Keystore // Location of key files
	Opts     map[string]string  // Per chain options
}
