// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

func DeployAssetStore(t *testing.T, client *utils.Client) common.Address {
	addr, err := utils.DeployAssetStore(client)
	if err != nil {
		t.Fatal(err)
	}
	return addr
}

func AssertHashExistence(t *testing.T, client *utils.Client, hash [32]byte, contract common.Address) {
	exists, err := utils.HashExists(client, hash, contract)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("Hash %x does not exist on chain", hash)
	}
	log15.Info("Assert existence in asset store", "hash", hash, "assetStore", contract)
}
