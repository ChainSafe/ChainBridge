// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"math/big"

	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterErc20Resource(client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address, rId msg.ResourceId, addr common.Address) error {
	instance, err := ERC20Handler.NewERC20Handler(erc20Handler, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(opts, client)
	if err != nil {
		return err
	}
	_, err = instance.SetResourceIDAndContractAddress(opts, rId, addr)
	if err != nil {
		return err
	}
	return nil
}

func RegisterGenericResource(client *ethclient.Client, opts *bind.TransactOpts, genericHandler common.Address, rId msg.ResourceId, addr common.Address) error {
	instance, err := centrifugeHandler.NewCentrifugeAssetHandler(genericHandler, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	_, err = instance.SetResourceIDAndContractAddress(opts, rId, addr)
	if err != nil {
		return err
	}
	return nil
}

// DeployMintAndApprove deploys a new erc20 contract, mints to the deployer, and approves the erc20 handler to transfer those token.
func DeployMintApproveErc20(client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address, amount *big.Int) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	// Deploy
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	// Mint
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Mint(opts, opts.From, amount)
	if err != nil {
		return ZeroAddress, err
	}

	// Approve
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Approve(opts, erc20Handler, amount)
	if err != nil {
		return ZeroAddress, err
	}

	return erc20Addr, nil
}

func GetBalance(client *ethclient.Client, erc20Contract, account common.Address) (*big.Int, error) { //nolint:unused,deadcode
	instance, err := erc20Mintable.NewERC20Mintable(erc20Contract, client)
	if err != nil {
		return nil, err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return nil, err

	}
	return bal, nil

}

func ApproveErc20(client *ethclient.Client, opts *bind.TransactOpts, contractAddress, recipient common.Address, amount *big.Int) error {
	err := UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	erc20Instance, err := erc20Mintable.NewERC20Mintable(contractAddress, client)
	if err != nil {
		return err
	}

	_, err = erc20Instance.Approve(opts, recipient, amount)
	if err != nil {
		return err
	}
	return nil
}

func FundErc20Handler(client *ethclient.Client, opts *bind.TransactOpts, handlerAddress, erc20Address common.Address, amount *big.Int) error {
	err := UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	err = ApproveErc20(client, opts, erc20Address, handlerAddress, amount)
	if err != nil {
		return err
	}

	instance, err := ERC20Handler.NewERC20Handler(handlerAddress, client)
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
