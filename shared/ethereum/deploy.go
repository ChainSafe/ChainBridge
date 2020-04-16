// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	relayer "github.com/ChainSafe/ChainBridge/bindings/Relayer"
	"github.com/ChainSafe/ChainBridge/keystore"
)

var (
	RelayerAddresses = []common.Address{
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.BobKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.DaveKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.EveKey].Address()),
	}

	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

type DeployedContracts struct {
	BridgeAddress            common.Address
	RelayerAddress           common.Address
	ERC20HandlerAddress      common.Address
	ERC721HandlerAddress     common.Address
	CentrifugeHandlerAddress common.Address
}

// DeployContracts deploys Bridge, Relayer, ERC20Handler, ERC721Handler and CentrifugeAssetHandler and returns the addresses
func DeployContracts(deployPK string, chainID uint8, url string, initialRelayerThreshold *big.Int) (*DeployedContracts, error) {

	client, opts, err := accountSetUp(url, deployPK)
	if err != nil {
		return nil, err
	}

	relayerAddr, err := deployRelayer(opts, client, RelayerAddresses, initialRelayerThreshold)
	if err != nil {
		return nil, err
	}

	bridgeAddr, err := deployBridge(opts, client, chainID, relayerAddr, initialRelayerThreshold)
	if err != nil {
		return nil, err
	}

	erc20HandlerAddr, err := deployERC20Handler(opts, client, bridgeAddr)
	if err != nil {
		return nil, err
	}

	erc721HandlerAddr, err := deployERC721Handler(opts, client, bridgeAddr)
	if err != nil {
		return nil, err
	}

	centrifugeHandlerAddr, err := deployCentrifugeHandler(opts, client, bridgeAddr)
	if err != nil {
		return nil, err
	}

	deployedContracts := DeployedContracts{bridgeAddr, relayerAddr, erc20HandlerAddr, erc721HandlerAddr, centrifugeHandlerAddr}

	return &deployedContracts, nil

}

func UpdateNonce(opts *bind.TransactOpts, client *ethclient.Client) error {
	newNonce, err := client.PendingNonceAt(context.Background(), opts.From)
	if err != nil {
		return err
	}

	opts.Nonce = big.NewInt(int64(newNonce))

	return nil
}

func accountSetUp(url string, deployPK string) (*ethclient.Client, *bind.TransactOpts, error) {

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(deployPK)
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, err
	}

	deployAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), deployAddress)
	if err != nil {
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6721975)
	auth.GasPrice = gasPrice

	return client, auth, nil

}

func deployBridge(opts *bind.TransactOpts, client *ethclient.Client, chainID uint8, relayerContract common.Address, initialRelayerThreshold *big.Int) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	bridgeAddr, _, _, err := bridge.DeployBridge(opts, client, chainID, relayerContract, initialRelayerThreshold)
	if err != nil {
		return ZeroAddress, err
	}

	return bridgeAddr, nil

}

func deployRelayer(opts *bind.TransactOpts, client *ethclient.Client, initialRelayers []common.Address, initialRelayerThreshold *big.Int) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	relAddr, _, _, err := relayer.DeployRelayer(opts, client, initialRelayers, initialRelayerThreshold)
	if err != nil {
		return ZeroAddress, err
	}

	return relAddr, nil
}

func deployERC20Handler(opts *bind.TransactOpts, client *ethclient.Client, bridgeAddress common.Address) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	erc20HandlerAddr, _, _, err := erc20Handler.DeployERC20Handler(opts, client, bridgeAddress, [][32]byte{}, []common.Address{})
	if err != nil {
		return ZeroAddress, err
	}

	return erc20HandlerAddr, nil
}

func deployERC721Handler(opts *bind.TransactOpts, client *ethclient.Client, bridgeAddress common.Address) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	erc721HandlerAddr, _, _, err := erc721Handler.DeployERC721Handler(opts, client, bridgeAddress, [][32]byte{}, []common.Address{})
	if err != nil {
		return ZeroAddress, err
	}

	return erc721HandlerAddr, nil
}

func deployCentrifugeHandler(opts *bind.TransactOpts, client *ethclient.Client, bridgeAddress common.Address) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	centrifugeHandlerAddr, _, _, err := centrifugeHandler.DeployCentrifugeAssetHandler(opts, client, bridgeAddress, [][32]byte{}, []common.Address{})
	if err != nil {
		return ZeroAddress, err
	}

	return centrifugeHandlerAddr, nil
}

func DeployErc20Contract(opts *bind.TransactOpts, client *ethclient.Client) (common.Address, error) {
	err := UpdateNonce(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	erc20Addr, _, _, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		return ZeroAddress, err
	}

	return erc20Addr, nil
}
