// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"github.com/ChainSafe/log15"

	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
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
func createErc20Deposit(contract BridgeContract, conn *Connection, txOpts *bind.TransactOpts, deployerAddress, originHandler, destHandler, destTokenAddress, destRecipient common.Address, destId, amount *big.Int) error {
	erc20Address, err := deployErc20Contract(txOpts, conn.conn, deployerAddress)
	if err != nil {
		log15.Info("deployErc20Contract")
		return err
	}

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if err := mintErc20Tokens(conn, txOpts, erc20Address, deployerAddress, amount); err != nil {
		log15.Info("MintErc20Tokens")
		return err
	}

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if err := approveErc20(conn, txOpts, erc20Address, originHandler, amount); err != nil {
		log15.Info("ApproveErc20")
		return err
	}

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))

	data := constructDataBytes(erc20Address, destHandler, destTokenAddress, destRecipient, destId, amount)

	if _, err := contract.Deposit(
		txOpts,
		originHandler,
		data,
	); err != nil {
		log15.Info("Deposit")

		return err
	}
	return nil
}

func constructDataBytes(erc20Address, destHandler, destTokenAddress, destRecipient common.Address, destId, amount *big.Int) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(erc20Address.Bytes(), 32)...)
	data = append(data, math.PaddedBigBytes(destId, 32)...)
	data = append(data, common.LeftPadBytes(destHandler.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(destHandler.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(destTokenAddress.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(amount.Bytes(), 32)...)

	return data
}

func createDepositProposal(contract BridgeContract, handlerAddress common.Address, conn *Connection, txOpts *bind.TransactOpts, destChain, depositNonce *big.Int, metadata [32]byte) error {
	if _, err := contract.BridgeRaw.Transact(
		txOpts,
		"createDepositProposal",
		destChain,
		handlerAddress,
		depositNonce,
		metadata,
	); err != nil {
		return err
	}
	return nil
}
