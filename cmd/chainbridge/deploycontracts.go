// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	bridge "github.com/ChainSafe/ChainBridgeV2/contracts/Bridge"
	centrifugeHandler "github.com/ChainSafe/ChainBridgeV2/contracts/CentrifugeAssetHandler"
	erc20Handler "github.com/ChainSafe/ChainBridgeV2/contracts/ERC20Handler"
	erc721Handler "github.com/ChainSafe/ChainBridgeV2/contracts/ERC721Handler"
	relayer "github.com/ChainSafe/ChainBridgeV2/contracts/Relayer"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	log "github.com/ChainSafe/log15"
)

var deployContractsCommand = cli.Command{
	Action:      parseCommands,
	Name:        "deploycontracts",
	Usage:       "deploys contracts",
	Category:    "tests",
	Flags:       deployContractsFlags,
	Description: "\tthe deploycontracts command is used to deploy contracts on a local network for testing purposes\n",
}

var (
	// Keys generate from: when sound uniform light fee face forum huge impact talent exhaust arrow
	// DEPLOYER_PRIV_KEY = "000000000000000000000000000000000000000000000000000000416c696365"

	RELAYER_ADDRESS = []string{
		keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.BobKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.DaveKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.EveKey].Address(),
	}

	ZERO_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

type DeployedContracts struct {
	BridgeAddress            common.Address
	RelayerAddress           common.Address
	ERC20HandlerAddress      common.Address
	ERC721HandlerAddress     common.Address
	CentrifugeHandlerAddress common.Address
}

func parseCommands(ctx *cli.Context) error {
	log.Info("Deploying Contracts")
	log.Info(RELAYER_ADDRESS[0])

	port := ctx.String(PortFlag.Name)
	relayers := ctx.Int(NumRelayersFlag.Name)
	relayerThreshold := ctx.Int(RelayerThresholdFlag.Name)
	minCount := ctx.Int(MinCountFlag.Name)
	deployPK := ctx.String(PKFlag.Name)

	deployedContracts, err := deployContracts(deployPK, port, relayers, big.NewInt(int64(relayerThreshold)), uint8(minCount))
	if err != nil {
		return err
	}

	log.Info("Bridge Contract Deployed at: " + deployedContracts.BridgeAddress.Hex())
	log.Info("Relayer Contract Deployed at: " + deployedContracts.RelayerAddress.Hex())
	log.Info("ERC20 Handler Contract Deployed at: " + deployedContracts.ERC20HandlerAddress.Hex())
	log.Info("ERC721 Handler Contract Deployed at: " + deployedContracts.ERC721HandlerAddress.Hex())
	log.Info("Centrifuge Asset Handler Contract Deployed at: " + deployedContracts.CentrifugeHandlerAddress.Hex())

	return nil
}

func deployContracts(deployPK string, port string, relayers int, initialRelayerThreshold *big.Int, minCount uint8) (*DeployedContracts, error) {

	client, auth, deployAddress, initialRelayerAddresses, err := accountSetUp(port, relayers, deployPK)
	if err != nil {
		return nil, err
	}

	relayerAddr, err := deployRelayer(auth, client, initialRelayerAddresses, initialRelayerThreshold, deployAddress)
	if err != nil {
		return nil, err
	}

	bridgeAddr, err := deployBridge(auth, client, relayerAddr, initialRelayerThreshold, deployAddress)
	if err != nil {
		return nil, err
	}

	erc20HandlerAddr, err := deployERC20Handler(auth, client, bridgeAddr, deployAddress)
	if err != nil {
		return nil, err
	}

	erc721HandlerAddr, err := deployERC721Handler(auth, client, bridgeAddr, deployAddress)
	if err != nil {
		return nil, err
	}

	centrifugeHandlerAddr, err := deployCentrifugeHandler(auth, client, bridgeAddr, deployAddress)
	if err != nil {
		return nil, err
	}

	deployedContracts := DeployedContracts{bridgeAddr, relayerAddr, erc20HandlerAddr, erc721HandlerAddr, centrifugeHandlerAddr}

	return &deployedContracts, nil

}

func createRelayerSlice(valAddr []string, numRelayers int) []common.Address {

	relayerAddresses := make([]common.Address, numRelayers)

	for i := 0; i < numRelayers; i++ {
		relayerAddresses[i] = common.HexToAddress(valAddr[i])

	}

	return relayerAddresses
}

func updateNonce(auth *bind.TransactOpts, client *ethclient.Client, deployAddress common.Address) (*bind.TransactOpts, error) {

	newNonce, err := client.PendingNonceAt(context.Background(), deployAddress)
	if err != nil {
		log.Error("error fetching latest nonce")
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(newNonce))

	return auth, nil

}

func accountSetUp(port string, relayers int, deployPK string) (*ethclient.Client, *bind.TransactOpts, common.Address, []common.Address, error) {

	client, err := ethclient.Dial("http://localhost:" + port)
	if err != nil {
		log.Error("error connecting to client")
		return nil, nil, ZERO_ADDRESS, nil, err
	}

	privateKey, err := crypto.HexToECDSA(deployPK)
	if err != nil {
		log.Error("error casting private key to ECDSA")
		return nil, nil, ZERO_ADDRESS, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("error casting public key to EDCSA")
		return nil, nil, ZERO_ADDRESS, nil, err
	}

	deployAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), deployAddress)
	if err != nil {
		return nil, nil, ZERO_ADDRESS, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("error fetching latest gas price")
		return nil, nil, ZERO_ADDRESS, nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6721975)
	auth.GasPrice = gasPrice

	relayerAddresses := createRelayerSlice(RELAYER_ADDRESS, relayers)

	return client, auth, deployAddress, relayerAddresses, nil

}

func deployBridge(auth *bind.TransactOpts, client *ethclient.Client, relayerContract common.Address, initialRelayerThreshold *big.Int, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	bridgeAddr, _, _, err := bridge.DeployBridge(auth, client, relayerContract, initialRelayerThreshold)
	if err != nil {
		log.Error("error deploying bridge instance")
		return ZERO_ADDRESS, err
	}

	return bridgeAddr, nil

}

func deployRelayer(auth *bind.TransactOpts, client *ethclient.Client, initialRelayers []common.Address, initialRelayerThreshold *big.Int, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	relAddr, _, _, err := relayer.DeployRelayer(auth, client, initialRelayers, initialRelayerThreshold)
	if err != nil {
		log.Error("error deploying relayer instance")
		return ZERO_ADDRESS, err
	}

	return relAddr, nil
}

func deployERC20Handler(auth *bind.TransactOpts, client *ethclient.Client, bridgeAddress common.Address, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	erc20HandlerAddr, _, _, err := erc20Handler.DeployERC20Handler(auth, client, bridgeAddress)
	if err != nil {
		log.Error("error deploying ERC20 Handler instance")
		return ZERO_ADDRESS, err
	}

	return erc20HandlerAddr, nil

}

func deployERC721Handler(auth *bind.TransactOpts, client *ethclient.Client, bridgeAddress common.Address, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	erc721HandlerAddr, _, _, err := erc721Handler.DeployERC721Handler(auth, client, bridgeAddress)
	if err != nil {
		log.Error("error deploying ERC20 Handler instance")
		return ZERO_ADDRESS, err
	}

	return erc721HandlerAddr, nil

}

func deployCentrifugeHandler(auth *bind.TransactOpts, client *ethclient.Client, bridgeAddress common.Address, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	centrifugeHandlerAddr, _, _, err := centrifugeHandler.DeployCentrifugeAssetHandler(auth, client, bridgeAddress)
	if err != nil {
		log.Error("error deploying Centrifuge Asset Handler instance")
		return ZERO_ADDRESS, err
	}

	return centrifugeHandlerAddr, nil
}
