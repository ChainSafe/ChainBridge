// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"math/big"
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterErc721Resource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, handler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterErc721Resource(client, opts, handler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721Deploy(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts) common.Address {
	addr, err := utils.DeployErc721(client, opts)
	if err != nil {
		t.Fatal(err)
	}
	return addr
}

func Erc721Mint(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, id *big.Int, metadata []byte) {
	log15.Info("Minting erc721 token", "contract", erc721Contract.Hex(), "id", id)
	err := utils.Erc721Mint(client, opts, erc721Contract, id, metadata)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721MintMany(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, start, numberOfTokens int) {
	for i := start; i < start+numberOfTokens; i++ {
		Erc721Mint(t, client, opts, erc721Contract, big.NewInt(int64(i)), []byte{})
	}
}

func Erc721Approve(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc721Contract, recipient common.Address, tokenId *big.Int) {
	log15.Info("Approving erc721 token for transfer", "contract", erc721Contract.Hex(), "id", tokenId, "recipient", recipient)
	err := utils.ApproveErc721(client, opts, erc721Contract, recipient, tokenId)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721ApproveMany(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc721Contract, recipient common.Address, start, numberOfTokens int) {
	for i := start; i < start+numberOfTokens; i++ {
		Erc721Approve(t, client, opts, erc721Contract, recipient, big.NewInt(int64(i)))
	}
}

func Erc721AssertOwner(t *testing.T, client *ethclient.Client, erc721Contract common.Address, tokenId *big.Int, expected common.Address) {
	addr, err := utils.OwnerOf(client, erc721Contract, tokenId)
	if err != nil {
		t.Fatal(err)
	}

	if addr != expected {
		t.Fatalf("address %s does not own %x, %s does", expected.Hex(), tokenId.Bytes(), addr.Hex())
	}
	log15.Info("Asserted ownership of erc721", "tokenId", tokenId, "owner", addr)
}

func Erc721FundHandler(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, handler, erc721Contract common.Address, tokenId *big.Int) {
	err := utils.FundErc721Handler(client, opts, handler, erc721Contract, tokenId)
	if err != nil {
		t.Fatal(err)
	}
	log15.Info("Funded handler with erc721 token", "tokenId", tokenId, "handler", handler)
}

func Erc721FundHandlerMany(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, handler, erc721Contract common.Address, start, numberOfTokens int) {
	for i := start; i < start+numberOfTokens; i++ {
		Erc721FundHandler(t, client, opts, handler, erc721Contract, big.NewInt(int64(i)))
	}
}

func Erc721AddMinter(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, minter common.Address) {
	err := utils.Erc721AddMinter(client, opts, erc721Contract, minter)
	if err != nil {
		t.Fatal(err)
	}
}

func Erc721SetBurnable(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, handler, erc721Contract common.Address) {
	err := utils.Erc721SetBurnable(client, opts, handler, erc721Contract)
	if err != nil {
		t.Fatal(err)
	}
}
