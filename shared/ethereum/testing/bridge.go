// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterResource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, bridge, handler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterResource(client, opts, bridge, handler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func RegisterGenericResource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, bridge, handler common.Address, rId msg.ResourceId, addr common.Address, depositSig, executeSig [4]byte) {
	err := utils.RegisterGenericResource(client, opts, bridge, handler, rId, addr, depositSig, executeSig)
	if err != nil {
		t.Fatal(err)
	}
}

func SetBurnable(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, bridge, handler, contract common.Address) {
	err := utils.SetBurnable(client, opts, bridge, handler, contract)
	if err != nil {
		t.Fatal(err)
	}
}

func GetDepositNonce(t *testing.T, client *ethclient.Client, bridge common.Address, chain msg.ChainId) uint64 {
	count, err := utils.GetDepositNonce(client, bridge, chain)
	if err != nil {
		t.Fatal(err)
	}
	return count
}
