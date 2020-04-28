// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"context"
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(t *testing.T, endpoint string, kp *secp256k1.Keypair) *utils.Client {
	client, err := utils.NewClient(endpoint, kp)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func GetLatestBlock(t *testing.T, client *ethclient.Client) *big.Int {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	return block.Number()
}
