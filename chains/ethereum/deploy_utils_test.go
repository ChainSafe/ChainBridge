// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/utils/ethereum"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func deployTestContracts(t *testing.T, id msg.ChainId) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		hexutil.Encode(AliceKp.Encode())[2:],
		uint8(id),
		TestEndpoint,
		TestRelayerThreshold,
	)
	if err != nil {
		t.Fatal(err)
	}

	return contracts
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
	log.Info("Approved tokens", "spender", recipient, "amount", amount.String())
	return nil
}

func fundErc20Handler(conn *Connection, opts *bind.TransactOpts, handlerAddress, erc20Address common.Address, amount *big.Int) error {
	err := approveErc20(conn, opts, erc20Address, handlerAddress, amount)
	if err != nil {
		return err
	}

	instance, err := erc20Handler.NewERC20Handler(handlerAddress, conn.conn)
	if err != nil {
		return err
	}

	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = instance.FundERC20(opts, erc20Address, opts.From, amount)
	if err != nil {
		return err
	}
	return nil
}

// createErc20Deposit deploys a new erc20 token contract mints, the sender (based on value), and creates a deposit
func createErc20Deposit(contract *Bridge.Bridge,
	txOpts *bind.TransactOpts,
	rId msg.ResourceId,
	originHandler,
	destRecipient common.Address,
	destId msg.ChainId,
	amount *big.Int) error {

	data := utils.ConstructErc20DepositData(rId, destRecipient.Bytes(), amount)

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if _, err := contract.Deposit(
		txOpts,
		uint8(destId),
		originHandler,
		data,
	); err != nil {
		return err
	}
	return nil
}
