// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const BridgePalletName = "ChainBridge"
const BridgeStoragePrefix = "Bridge"

type Erc721Token struct {
	Id       types.U256
	Metadata types.Bytes
}
