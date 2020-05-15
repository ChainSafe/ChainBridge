// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
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
	BridgeAddress         common.Address
	ERC20HandlerAddress   common.Address
	ERC721HandlerAddress  common.Address
	GenericHandlerAddress common.Address
}

// DeployContracts deploys Bridge, Relayer, ERC20Handler, ERC721Handler and CentrifugeAssetHandler and returns the addresses
func DeployContracts(deployPK string, chainID uint8, url string, initialRelayerThreshold *big.Int) (*DeployedContracts, error) {

	client, err := accountSetUp(url, deployPK)
	if err != nil {
		return nil, err
	}

	bridgeAddr, err := deployBridge(client, chainID, RelayerAddresses, initialRelayerThreshold)
	if err != nil {
		return nil, err
	}

	erc20HandlerAddr, err := deployERC20Handler(client, bridgeAddr)
	if err != nil {
		return nil, err
	}

	erc721HandlerAddr, err := deployERC721Handler(client, bridgeAddr)
	if err != nil {
		return nil, err
	}

	genericHandlerAddr, err := deployGenericHandler(client, bridgeAddr)
	if err != nil {
		return nil, err
	}

	deployedContracts := DeployedContracts{bridgeAddr, erc20HandlerAddr, erc721HandlerAddr, genericHandlerAddr}

	return &deployedContracts, nil

}

func UpdateNonce(client *Client) error {
	newNonce, err := client.PendingNonceAt(context.Background(), client.CallOpts.From)
	if err != nil {
		return err
	}

	client.CallOpts.Nonce = big.NewInt(int64(newNonce))

	return nil
}

func accountSetUp(url string, deployPK string) (*Client, error) {

	newKeyPair, err := secp256k1.NewKeypairFromString(deployPK)
	if err != nil {
		return nil, nil, err
	}

	client, err := NewClient(url, newKeyPair)
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

func deployBridge(client *Client, chainID uint8, relayerAddrs []common.Address, initialRelayerThreshold *big.Int) (common.Address, error) {
	err := UpdateNonce(client)
	if err != nil {
		return ZeroAddress, err
	}

	bridgeAddr, tx, _, err := bridge.DeployBridge(client, chainID, relayerAddrs, initialRelayerThreshold, big.NewInt(0))
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return bridgeAddr, nil

}

func deployERC20Handler(client *Client, bridgeAddress common.Address) (common.Address, error) {
	err := UpdateNonce(client)
	if err != nil {
		return ZeroAddress, err
	}

	erc20HandlerAddr, tx, _, err := erc20Handler.DeployERC20Handler(client, bridgeAddress, [][32]byte{}, []common.Address{}, []common.Address{})
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return erc20HandlerAddr, nil
}

func deployERC721Handler(client *Client, bridgeAddress common.Address) (common.Address, error) {
	err := UpdateNonce(client)
	if err != nil {
		return ZeroAddress, err
	}

	erc721HandlerAddr, tx, _, err := erc721Handler.DeployERC721Handler(client, bridgeAddress, [][32]byte{}, []common.Address{}, []common.Address{})
	if err != nil {
		return ZeroAddress, err
	}
	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return erc721HandlerAddr, nil
}

func deployGenericHandler(client *Client, bridgeAddress common.Address) (common.Address, error) {
	err := UpdateNonce(client)
	if err != nil {
		return ZeroAddress, err
	}

	addr, tx, _, err := GenericHandler.DeployGenericHandler(client, bridgeAddress, [][32]byte{}, []common.Address{}, [][4]byte{}, [][4]byte{})
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return addr, nil
}
