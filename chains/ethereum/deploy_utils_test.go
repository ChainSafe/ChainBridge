// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func deployTestContracts(t *testing.T, id msg.ChainId) *DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := DeployContracts(
		hexutil.Encode(AliceKp.Encode())[2:],
		big.NewInt(int64(id)),
		TestEndpoint,
		TestNumRelayers,
		TestRelayerThreshold,
		0)
	if err != nil {
		t.Fatal(err)
	}

	return contracts
}

func deployMintApproveErc20(t *testing.T, conn *Connection, opts *bind.TransactOpts) common.Address {
	// Get transaction ready
	deployerAddress := ethcrypto.PubkeyToAddress(conn.kp.PrivateKey().PublicKey)

	erc20Address, err := deployErc20Contract(opts, conn.conn, deployerAddress)
	if err != nil {
		t.Fatal(err)
	}

	if err := addHandlerAsMinter(conn, opts, erc20Address); err != nil {
		t.Fatal(err)
	}

	if err := mintErc20Tokens(conn, opts, erc20Address, TestMintAmount); err != nil {
		t.Fatal(err)
	}

	if err := approveErc20(conn, opts, erc20Address, conn.cfg.erc20HandlerContract, big.NewInt(100)); err != nil {
		t.Fatal(err)
	}

	return erc20Address
}

func addHandlerAsMinter(conn *Connection, opts *bind.TransactOpts, contract common.Address) error {
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Instance, err := erc20Mintable.NewERC20Mintable(contract, conn.conn)
	if err != nil {
		return err
	}

	_, err = erc20Instance.AddMinter(opts, conn.cfg.erc20HandlerContract)
	if err != nil {
		return err
	}
	return nil
}

func mintErc20Tokens(connection *Connection, opts *bind.TransactOpts, contractAddress common.Address, amount *big.Int) error {
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Instance, err := erc20Mintable.NewERC20Mintable(contractAddress, connection.conn)
	if err != nil {
		return err
	}

	_, err = erc20Instance.Mint(opts, opts.From, amount)
	if err != nil {
		return err
	}

	return nil
}

func approveErc20(connection *Connection, opts *bind.TransactOpts, contractAddress, recipient common.Address, amount *big.Int) error {
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
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

// constructDataBytes constructs the data field to be passed into a deposit call
func constructDataBytes(erc20Address, destHandler, destTokenAddress, destRecipient common.Address, amount *big.Int) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(erc20Address.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(destHandler.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(destTokenAddress.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(destRecipient.Bytes(), 32)...)
	data = append(data, math.PaddedBigBytes(amount, 32)...)

	return data
}

// createErc20Deposit deploys a new erc20 token contract mints, the sender (based on value), and creates a deposit
func createErc20Deposit(contract *BridgeContract,
	txOpts *bind.TransactOpts,
	erc20Address,
	originHandler,
	destHandler,
	destTokenAddress,
	destRecipient common.Address,
	destId,
	amount *big.Int) error {

	data := constructDataBytes(erc20Address, destHandler, destTokenAddress, destRecipient, amount)

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if _, err := contract.Deposit(
		txOpts,
		destId,
		originHandler,
		data,
	); err != nil {
		return err
	}
	return nil
}
