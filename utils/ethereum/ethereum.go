package utils

import (
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func whitelistResourceId(client *ethclient.Client, opts *bind.TransactOpts, erc20handler common.Address, rId msg.ResourceId, addr common.Address) error {
	instance, err := erc20Handler.NewERC20Handler(erc20handler, client)
	if err != nil {
		return err
	}

	_, err = instance.SetResourceIDAndContractAddress(opts, rId, addr)
	if err != nil {
		return err
	}
	return nil
}
