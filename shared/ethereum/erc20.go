// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"math/big"

	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20PresetMinterPauser"
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
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20PresetMinterPauser(opts, client, "", "")
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

func DeployAndMintErc20(client *ethclient.Client, opts *bind.TransactOpts, amount *big.Int) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	// Deploy
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20PresetMinterPauser(opts, client, "", "")
	if err != nil {
		return ZeroAddress, err
	}

	// Mint
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Mint(opts, opts.From, amount)
	if err != nil {
		return ZeroAddress, err
	}

	return erc20Addr, nil
}

func Erc20Approve(client *ethclient.Client, opts *bind.TransactOpts, erc20Contract, recipient common.Address, amount *big.Int) error {
	err := UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	instance, err := erc20Mintable.NewERC20PresetMinterPauser(erc20Contract, client)
	if err != nil {
		return err
	}

	_, err = instance.Approve(opts, recipient, amount)
	if err != nil {
		return err
	}
	return nil
}

func Erc20GetBalance(client *ethclient.Client, erc20Contract, account common.Address) (*big.Int, error) { //nolint:unused,deadcode
	instance, err := erc20Mintable.NewERC20PresetMinterPauser(erc20Contract, client)
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

	erc20Instance, err := erc20Mintable.NewERC20PresetMinterPauser(contractAddress, client)
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

func Erc20AddMinter(client *ethclient.Client, opts *bind.TransactOpts, erc20Contract, handler common.Address) error {
	err := UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	instance, err := erc20Mintable.NewERC20PresetMinterPauser(erc20Contract, client)
	if err != nil {
		return err
	}

	role, err := instance.MINTERROLE(&bind.CallOpts{})
	if err != nil {
		return err
	}

	_, err = instance.GrantRole(opts, role, handler)
	if err != nil {
		return err
	}
	return nil
}

func Erc20GetAllowance(client *ethclient.Client, erc20Contract, owner, spender common.Address) (*big.Int, error) {
	instance, err := erc20Mintable.NewERC20PresetMinterPauser(erc20Contract, client)
	if err != nil {
		return nil, err
	}

	amount, err := instance.Allowance(&bind.CallOpts{}, owner, spender)
	if err != nil {
		return nil, err
	}

	return amount, nil
}

func Erc20GetResourceId(client *ethclient.Client, handler common.Address, rId msg.ResourceId) (common.Address, error) {
	instance, err := ERC20Handler.NewERC20Handler(handler, client)
	if err != nil {
		return ZeroAddress, err
	}

	addr, err := instance.ResourceIDToTokenContractAddress(&bind.CallOpts{}, rId)
	if err != nil {
		return ZeroAddress, err
	}

	return addr, nil
}

func Erc20SetBurnable(client *ethclient.Client, opts *bind.TransactOpts, handler, erc20Contract common.Address) error {
	instance, err := ERC20Handler.NewERC20Handler(handler, client)
	if err != nil {
		return err
	}

	err = UpdateNonce(opts, client)
	if err != nil {
		return err
	}

	_, err = instance.SetBurnable(opts, erc20Contract)
	if err != nil {
		return err
	}
	return nil
}
