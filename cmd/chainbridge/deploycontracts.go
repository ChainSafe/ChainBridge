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

	bridgeAsset "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	log "github.com/ChainSafe/log15"
)

var deployContractsLocalCommand = cli.Command{
	Action:      parseCommands,
	Name:        "deploycontractslocal",
	Usage:       "deploys contracts",
	Category:    "tests",
	Flags:       deployContractLocalFlags,
	Description: "\tthe deploycontractslocal command is used to deploy contracts on a local network for testing purposes\n",
}

var (
	// Keys generate from: when sound uniform light fee face forum huge impact talent exhaust arrow
	DEPLOYER_PRIV_KEY = "000000000000000000000000000000000000000000000000000000416c696365"

	VALIDATOR_ADDRESS = []string{
		keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.BobKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.DaveKey].Address(),
		keystore.TestKeyRing.EthereumKeys[keystore.EveKey].Address(),
	}

	ZERO_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func parseCommands(ctx *cli.Context) error {
	log.Info("Deploying Contracts")
	log.Info(VALIDATOR_ADDRESS[0])

	port := ctx.String(PortFlag.Name)
	validators := ctx.Int(NumValidatorsFlag.Name)
	validatorThreshold := ctx.Int(ValidatorThresholdFlag.Name)
	depositThreshold := ctx.Int(DepositThresholdFlag.Name)
	minCount := ctx.Int(MinCountFlag.Name)
	deployPK := ctx.String(PKFlag.Name)

	recieverAddr, emitterAddr, bridgeAssetAddr, err := deployContractsLocal(deployPK, port, validators, big.NewInt(int64(validatorThreshold)), big.NewInt(int64(depositThreshold)), uint8(minCount))
	if err != nil {
		return err
	}

	log.Info("Reciever Contract Deployed at: " + recieverAddr.Hex())
	log.Info("Emitter Contract Deployed at: " + emitterAddr.Hex())
	log.Info("Bridge Asset Contract Deployed at: " + bridgeAssetAddr.Hex())

	return nil
}

func deployContractsLocal(deployPK string, port string, validators int, validatorThreshold *big.Int, depositThreshold *big.Int, minCount uint8) (common.Address, common.Address, common.Address, error) {

	client, auth, deployAddress, validatorAddresses, err := accountSetUp(port, validators, deployPK)
	if err != nil {
		return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
	}

	recieverAddr, err := deployReceiver(auth, client, validatorAddresses, depositThreshold, validatorThreshold, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
	}

	emitterAddr, err := deployEmitter(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
	}

	bridgeAssetAddr, err := deployBridgeAsset(auth, client, minCount, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
	}

	return recieverAddr, emitterAddr, bridgeAssetAddr, nil

}

func createValidatorSlice(valAddr []string, numValidators int) []common.Address {

	validatorAddresses := make([]common.Address, numValidators)

	for i := 0; i < numValidators; i++ {
		validatorAddresses[i] = common.HexToAddress(valAddr[i])

	}

	return validatorAddresses
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

func accountSetUp(port string, validators int, deployPK string) (*ethclient.Client, *bind.TransactOpts, common.Address, []common.Address, error) {

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

	validatorAddresses := createValidatorSlice(VALIDATOR_ADDRESS, validators)

	return client, auth, deployAddress, validatorAddresses, nil

}

func deployReceiver(auth *bind.TransactOpts, client *ethclient.Client, validatorAddresses []common.Address, depositThreshold *big.Int, validatorThreshold *big.Int, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	recAddr, _, _, err := receiver.DeployReceiver(auth, client, validatorAddresses, depositThreshold, validatorThreshold)
	if err != nil {
		log.Error("error deploying reciever instance")
		return ZERO_ADDRESS, err
	}

	return recAddr, nil
}

func deployEmitter(auth *bind.TransactOpts, client *ethclient.Client, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	emitterAddr, _, _, err := emitter.DeployEmitter(auth, client)
	if err != nil {
		log.Error("error deploying emitter instance")
		return ZERO_ADDRESS, err
	}

	return emitterAddr, nil

}

func deployBridgeAsset(auth *bind.TransactOpts, client *ethclient.Client, mc uint8, deployAddress common.Address) (common.Address, error) {

	auth, err := updateNonce(auth, client, deployAddress)
	if err != nil {
		return ZERO_ADDRESS, err
	}

	bridgeAssetAddr, _, _, err := bridgeAsset.DeployBridgeAsset(auth, client, mc)
	if err != nil {
		log.Error("error deploying bridge asset instance")
		return ZERO_ADDRESS, err
	}

	return bridgeAssetAddr, nil

}
