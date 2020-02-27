// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

import (
	"math/big"
)

type ChainId uint8

const EthereumId = ChainId(0)
const CentrifugeId = ChainId(1)

func (id ChainId) String() string {
	switch id {
	case EthereumId:
		return "ethereum"
	case CentrifugeId:
		return "centrifuge"
	default:
		return "unknown"
	}
}

func (id ChainId) Big() *big.Int {
	return big.NewInt(int64(id))
}
