// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"math/big"

	"github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC721MinterBurnerPauser"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DeployMintAndApprove deploys a new erc721 contract, mints to the deployer, and approves the erc20 handler to transfer those token.
func DeployErc721(client *ethclient.Client, opts *bind.TransactOpts) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	// Deploy
	addr, tx, _, err := ERC721MinterBurnerPauser.DeployERC721MinterBurnerPauser(opts, client, "", "", "")
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return addr, nil
}

func Erc721Mint(client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, id *big.Int, metadata []byte) error {
	instance, err := ERC721MinterBurnerPauser.NewERC721MinterBurnerPauser(erc721Contract, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	// Mint
	tx, err := instance.Mint(opts, opts.From, id, string(metadata))
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}
	return nil
}

func ApproveErc721(client *ethclient.Client, opts *bind.TransactOpts, contractAddress, recipient common.Address, tokenId *big.Int) error {
	err := UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	instance, err := ERC721MinterBurnerPauser.NewERC721MinterBurnerPauser(contractAddress, client)
	if err != nil {
		return err
	}

	tx, err := instance.Approve(opts, recipient, tokenId)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func FundErc721Handler(client *ethclient.Client, opts *bind.TransactOpts, handlerAddress, erc721Address common.Address, tokenId *big.Int) error {
	err := ApproveErc721(client, opts, erc721Address, handlerAddress, tokenId)
	if err != nil {
		return err
	}

	instance, err := ERC721Handler.NewERC721Handler(handlerAddress, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	tx, err := instance.FundERC721(opts, erc721Address, opts.From, tokenId)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func OwnerOf(client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, tokenId *big.Int) (common.Address, error) {
	instance, err := ERC721MinterBurnerPauser.NewERC721MinterBurnerPauser(erc721Contract, client)
	if err != nil {
		return ZeroAddress, err
	}
	return instance.OwnerOf(&bind.CallOpts{From: opts.From}, tokenId)
}

func Erc721GetTokenURI(client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, tokenId *big.Int) (string, error) {
	instance, err := ERC721MinterBurnerPauser.NewERC721MinterBurnerPauser(erc721Contract, client)
	if err != nil {
		return "", err
	}

	return instance.TokenURI(&bind.CallOpts{From: opts.From}, tokenId)
}

func Erc721AddMinter(client *ethclient.Client, opts *bind.TransactOpts, erc721Contract common.Address, minter common.Address) error {
	instance, err := ERC721MinterBurnerPauser.NewERC721MinterBurnerPauser(erc721Contract, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	role, err := instance.MINTERROLE(&bind.CallOpts{})
	if err != nil {
		return err
	}

	tx, err := instance.GrantRole(opts, role, minter)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}
