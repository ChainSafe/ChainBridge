// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/common/math"
)

// constructErc20Data constructs the data field to be passed into a deposit call
func ConstructErc20DepositData(rId msg.ResourceId, destRecipient []byte, amount *big.Int) []byte {
	var data []byte
	data = append(rId[:], math.PaddedBigBytes(amount, 32)...)
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...)
	data = append(data, destRecipient...)
	return data
}
