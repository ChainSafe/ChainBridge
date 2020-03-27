// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
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
	Name         string            // Human-readable chain name
	Id           msg.ChainId       // ChainID
	Endpoint     string            // url for rpc endpoint
	From         string            // address of key to use
	KeystorePath string            // Location of key files
	Insecure     bool              // Indicated whether the test keyring should be used
	Opts         map[string]string // Per chain options
}
