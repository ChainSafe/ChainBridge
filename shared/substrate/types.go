// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const BridgePalletName = "ChainBridge"
const BridgeStoragePrefix = "ChainBridge"

type Erc721Token struct {
	Id       types.U256
	Metadata types.Bytes
}

type RegistryId types.H160
type TokenId types.U256

type AssetId struct {
	RegistryId RegistryId
	TokenId    TokenId
}
