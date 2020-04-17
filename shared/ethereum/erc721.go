package utils

import (
	"math/big"

	"github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC721Mintable"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterErc721Resource(client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address, rId msg.ResourceId, addr common.Address) error {
	instance, err := ERC721Handler.NewERC721Handler(erc20Handler, client)
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

// DeployMintAndApprove deploys a new erc721 contract, mints to the deployer, and approves the erc20 handler to transfer those token.
func DeployMintApproveErc721(client *ethclient.Client, opts *bind.TransactOpts, erc721Handler common.Address, id *big.Int) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	// Deploy
	addr, _, instance, err := ERC721Mintable.DeployERC721Mintable(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	// Mint
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = instance.Mint(opts, opts.From, id)
	if err != nil {
		return ZeroAddress, err
	}

	// Approve
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = instance.Approve(opts, erc721Handler, id)
	if err != nil {
		return ZeroAddress, err
	}

	return addr, nil
}

func OwnerOf(client *ethclient.Client, erc721Contrct common.Address, tokenId *big.Int) (common.Address, error) {
	instance, err := ERC721Mintable.NewERC721Mintable(erc721Contrct, client)
	if err != nil {
		return ZeroAddress, err
	}

	return instance.OwnerOf(&bind.CallOpts{}, tokenId)
}
