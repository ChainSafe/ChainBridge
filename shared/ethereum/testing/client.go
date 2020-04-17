// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlock(t *testing.T, client *ethclient.Client) *big.Int {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	return block.Number()
}
