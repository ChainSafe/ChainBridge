// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/common/math"
)

// constructErc20Data constructs the data field to be passed into a n erc20 deposit call
func ConstructErc20DepositData(rId msg.ResourceId, destRecipient []byte, amount *big.Int) []byte {
	var data []byte
	data = append(rId[:], math.PaddedBigBytes(amount, 32)...)
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...)
	data = append(data, destRecipient...)
	return data
}

// constructErc20Data constructs the data field to be passed into an erc721 deposit call
func ConstructErc721DepositData(rId msg.ResourceId, tokenId *big.Int, destRecipient []byte) []byte {
	var data []byte
	data = append(rId[:], math.PaddedBigBytes(tokenId, 32)...)                             // Resource Id + Token Id
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...) // Length of recipient
	data = append(data, destRecipient...)                                                  // Recipient

	return data
}

func ConstructGenericDepositData(rId msg.ResourceId, metadata []byte) []byte {
	var data []byte
	data = append(rId[:], math.PaddedBigBytes(big.NewInt(int64(len(metadata))), 32)...)
	data = append(data, metadata...)

	return data
}
