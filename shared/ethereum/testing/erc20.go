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

func Erc20DeployMint(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, amount *big.Int) common.Address {
	addr, err := utils.DeployAndMintErc20(client, opts, amount)
	if err != nil {
		t.Fatal(err)
	}
	return addr
}

func Erc20Approve(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Contract, recipient common.Address, amount *big.Int) {
	err := utils.Erc20Approve(client, opts, erc20Contract, recipient, amount)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc20AssertBalance(t *testing.T, client *ethclient.Client, amount *big.Int, erc20Contract, account common.Address) { //nolint:unused,deadcode
	actual, err := utils.Erc20GetBalance(client, erc20Contract, account)
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

func Erc20BalanceOf(t *testing.T, client *ethclient.Client, erc20Contract, acct common.Address) *big.Int {
	balance, err := utils.Erc20GetBalance(client, erc20Contract, acct)
	if err != nil {
		t.Fatal(err)
	}
	return balance
}

func Erc20AddMinter(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Contract, handler common.Address) {
	err := utils.Erc20AddMinter(client, opts, erc20Contract, handler)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc20AssertAllowance(t *testing.T, client *ethclient.Client, erc20Contract, owner, spender common.Address, expected *big.Int) {
	amount, err := utils.Erc20GetAllowance(client, erc20Contract, owner, spender)
	if err != nil {
		t.Fatal(err)
	}

	if amount.Cmp(expected) != 0 {
		t.Fatalf("Allowance mismatch. Expected: %s Got: %s", expected.String(), amount.String())
	}
}

func Erc20AssertResourceMapping(t *testing.T, client *ethclient.Client, handler common.Address, rId msg.ResourceId, expected common.Address) {
	addr, err := utils.Erc20GetResourceId(client, handler, rId)
	if err != nil {
		t.Fatal(err)
	}

	if addr.String() != expected.String() {
		t.Fatalf("Unexpected address for resource ID %x. Expected: %x Got: %x", rId, expected, addr)
	}
}
