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
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployTestContracts(t *testing.T, id msg.ChainId) *DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := DeployContracts(
		hexutil.Encode(AliceKp.Encode())[2:],
		uint8(id),
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

	if err := mintErc20Tokens(conn, opts, erc20Address, opts.From, TestMintAmount); err != nil {
		t.Fatal(err)
	}

	if err := approveErc20(conn, opts, erc20Address, conn.cfg.erc20HandlerContract, TestMintAmount); err != nil {
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

func mintErc20Tokens(connection *Connection, opts *bind.TransactOpts, contractAddress, to common.Address, amount *big.Int) error {
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Instance, err := erc20Mintable.NewERC20Mintable(contractAddress, connection.conn)
	if err != nil {
		return err
	}

	_, err = erc20Instance.Mint(opts, to, amount)
	if err != nil {
		return err
	}
	log.Info("Minted tokens", "to", opts.From, "amount", amount.String())
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

// constructErc20Data constructs the data field to be passed into a deposit call
func constructErc20DepositData(erc20Contract common.Address, id msg.ChainId, destRecipient common.Address, amount *big.Int) []byte {
	var data []byte
	data = append(common.LeftPadBytes(erc20Contract.Bytes(), 31), byte(id))                        // Construct resource ID
	data = append(data, math.PaddedBigBytes(amount, 32)...)                                        // Amount
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient.Bytes()))), 32)...) // Recipient length
	data = append(data, destRecipient.Bytes()...)                                                  // Recipient
	return data
}

// createErc20Deposit deploys a new erc20 token contract mints, the sender (based on value), and creates a deposit
func createErc20Deposit(contract *Bridge.Bridge,
	txOpts *bind.TransactOpts,
	erc20Address,
	originHandler,
	destRecipient common.Address,
	destId msg.ChainId,
	amount *big.Int) error {

	chainId, err := contract.ChainID(&bind.CallOpts{})
	if err != nil {
		return err
	}

	data := constructErc20DepositData(erc20Address, msg.ChainId(chainId), destRecipient, amount)

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

func whitelistResourceId(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20handler common.Address, rId msg.ResourceId, addr common.Address) {
	_, err := updateNonce(opts, client, opts.From)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := erc20Handler.NewERC20Handler(erc20handler, client)
	if err != nil {
		t.Fatal(err)
	}

	_, err = instance.SetResourceIDAndContractAddress(opts, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
	log.Info("Whitelisted resource", "id", rId.Hex(), "addr", addr)
}
