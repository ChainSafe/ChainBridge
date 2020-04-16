// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"math/big"
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterErc20Resource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterErc20Resource(client, opts, erc20Handler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func RegisterGenericResource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, genericHandler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterGenericResource(client, opts, genericHandler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func DeployMintApproveErc20(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address, amount *big.Int) common.Address {
	addr, err := utils.DeployMintApproveErc20(client, opts, erc20Handler, amount)
	if err != nil {
		t.Fatal(err)
	}
	return addr
}

func AssertBalance(t *testing.T, client *ethclient.Client, amount *big.Int, erc20Contract, account common.Address) { //nolint:unused,deadcode
	actual, err := utils.GetBalance(client, erc20Contract, account)
	if err != nil {
		t.Fatal(err)
	}

	if actual.Cmp(amount) != 0 {
		t.Fatalf("Balance mismatch. Expected: %s Got: %s", amount.String(), actual.String())
	}
}

func FundErc20Handler(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, handlerAddress, erc20Address common.Address, amount *big.Int) {
	err := utils.FundErc20Handler(client, opts, handlerAddress, erc20Address, amount)
	if err != nil {
		t.Fatal(err)
	}
}
