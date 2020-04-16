// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// accoutData is the on-chain representation of an account
// TODO: Could live in GSRPC
type AccountData struct {
	Nonce    types.U32
	Refcount types.UCompact
	Data     struct {
		Free       types.U128
		Reserved   types.U128
		MiscFrozen types.U128
		FreeFrozen types.U128
	}
}
