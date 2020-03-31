// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func mintErc20Tokens(connection *Connection, opts *bind.TransactOpts, contractAddress, recipient common.Address, amount *big.Int) error {
	erc20Instance, err := erc20Mintable.NewERC20Mintable(contractAddress, connection.conn)
	if err != nil {
		return err
	}

	_, err = erc20Instance.Mint(opts, recipient, amount)
	if err != nil {
		return err
	}

	return nil
}

func approveErc20(connection *Connection, opts *bind.TransactOpts, contractAddress, recipient common.Address, amount *big.Int) error {
	erc20Instance, err := erc20Mintable.NewERC20Mintable(contractAddress, connection.conn)
	if err != nil {
		return err
	}

	_, err = erc20Instance.Approve(opts, recipient, amount)
	if err != nil {
		return err
	}

	return nil
}

// createErc20Deposit deploys a new erc20 token contract mints, the sender (based on value), and creates a deposit
func createErc20Deposit(contract BridgeContract, conn *Connection, txOpts *bind.TransactOpts, deployerAddress, originHandler, destHandler, destRecipient common.Address, destId, amount *big.Int) error {
	erc20Address, err := deployErc20Contract(txOpts, conn.conn, deployerAddress)
	if err != nil {
		return err
	}

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if err := mintErc20Tokens(conn, txOpts, erc20Address, deployerAddress, amount); err != nil {
		return err
	}

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if err := approveErc20(conn, txOpts, erc20Address, originHandler, amount); err != nil {
		return err
	}

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if _, err := contract.DepositERC20(
		txOpts,
		erc20Address,
		originHandler,
		destId,
		destHandler,
		destRecipient,
		amount,
	); err != nil {
		return err
	}
	return nil
}

func createDepositProposal(contract BridgeContract, conn *Connection, txOpts *bind.TransactOpts, destChain, depositNonce *big.Int, metadata [32]byte) error {
	if _, err := contract.BridgeRaw.Transact(
		txOpts,
		"createDepositProposal",
		destChain,
		depositNonce,
		metadata,
	); err != nil {
		return err
	}
	return nil
}
