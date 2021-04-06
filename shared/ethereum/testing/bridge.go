// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterResource(t *testing.T, client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterResource(client, bridge, handler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func RegisterGenericResource(t *testing.T, client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address, depositSig, executeSig [4]byte) {
	err := utils.RegisterGenericResource(client, bridge, handler, rId, addr, depositSig, executeSig)
	if err != nil {
		t.Fatal(err)
	}
}

func SetBurnable(t *testing.T, client *utils.Client, bridge, handler, contract common.Address) {
	err := utils.SetBurnable(client, bridge, handler, contract)
	if err != nil {
		t.Fatal(err)
	}
}

func GetDepositNonce(t *testing.T, client *utils.Client, bridge common.Address, chain msg.ChainId) uint64 {
	count, err := utils.GetDepositNonce(client, bridge, chain)
	if err != nil {
		t.Fatal(err)
	}
	return count
}
