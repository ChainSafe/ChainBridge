package main

import (
    "context"
    "crypto/ecdsa"
	"math/big"
    "strconv"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	bridgeAsset "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
    simpleEmitter "github.com/ChainSafe/ChainBridgeV2/contracts/SimpleEmitter"
    log "github.com/ChainSafe/log15"

)


var (
	// Keys generate from: when sound uniform light fee face forum huge impact talent exhaust arrow
    DEPLOYER_PRIV_KEY = "000000000000000000000000000000000000000000000000000000416c696365";
    
    VALIDATOR_ADDRESS = []string{
        "0x0c6CD6Dc5258EF556eA7c6dab2abE302fB60e0b6", // Alice Public Address
        "0x0E17A926c6525b59921846c85E1efD7a5396a47B", // Bob Public Address
        "0x0f05849291a309EC001bbd2dAd7DC6F989c40c80", // Charlie Public Address
        "0x3003d03276434dd23429D4D33090DA5948b7b510", // Dave Public Address
        "0x251e6F841549D519dE6De1e99241695bc1000A26", // Eve Public Address
    }

    VALIDATOR_PRIV_KEYS = []string{
        "0x000000000000000000000000000000000000000000000000000000416c696365", // Alice Private Key
        "0x0000000000000000000000000000000000000000000000000000000000426f62", // Bob Private Key
        "0x00000000000000000000000000000000000000000000000000436861726c6965", // Charlie Private Key
        "0x0000000000000000000000000000000000000000000000000000000044617665", // Dave Private Key
        "0x0000000000000000000000000000000000000000000000000000000000457665", // Eve Private Key
    }

    ZERO_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

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
        log.Error(err.Error())
        return nil, err
    }

    auth.Nonce = big.NewInt(int64(newNonce))

    return auth, nil

}

func AccountSetUp(port int, validators int) (*ethclient.Client, *bind.TransactOpts, []common.Address, error) {

    client, err := ethclient.Dial("http://localhost:"+strconv.Itoa(port))
    if err != nil {
        log.Error(err.Error())
        return nil, nil, nil, err
    }

    privateKey, err := crypto.HexToECDSA(DEPLOYER_PRIV_KEY)
    if err != nil {
        log.Error(err.Error())
        return nil, nil, nil, err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Error(err.Error())
        return nil, nil, nil, err
    }

    deployAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), deployAddress)
    if err != nil {
        log.Error(err.Error())
        return nil, nil, nil, err
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Error(err.Error())
        return nil, nil, nil, err
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)    
    auth.GasLimit = uint64(6721975) 
    auth.GasPrice = gasPrice

    validatorAddresses := createValidatorSlice(VALIDATOR_ADDRESS, validators)


    return client, auth, validatorAddresses, nil

}

func DeployReceiver(auth *bind.TransactOpts, client *ethclient.Client, validatorAddresses []common.Address, depositThreshold *big.Int, validatorThreshold *big.Int, deployAddress common.Address) (common.Address, error) {
    
    auth, err := updateNonce(auth, client, deployAddress)
    
    recAddr, _, _, err := receiver.DeployReceiver(auth, client, validatorAddresses, depositThreshold,validatorThreshold)
    if err != nil {
        log.Error("error deploying reciever instance")
        return ZERO_ADDRESS, err
    }

    return recAddr, nil
}

func DeployEmitter(auth *bind.TransactOpts, client *ethclient.Client, deployAddress common.Address) (common.Address, error) {

    auth, err := updateNonce(auth, client, deployAddress)

    emitterAddr, _, _, err := emitter.DeployEmitter(auth, client)
    if err != nil {
        log.Error("error deploying emitter instance")
        return ZERO_ADDRESS,err
    }

    return emitterAddr, nil

}

func DeploySimpleEmitter(auth *bind.TransactOpts, client *ethclient.Client, deployAddress common.Address) (common.Address, error) {

    auth, err := updateNonce(auth, client, deployAddress)

    sEmitterAddr, _, _, err := simpleEmitter.DeploySimpleEmitter(auth, client)
    if err != nil {
        log.Error("error deploying simple emitter instance")
        return ZERO_ADDRESS, err
    }

    return sEmitterAddr, nil
}

func DeployBridgeAsset(auth *bind.TransactOpts, client *ethclient.Client, mc uint8, deployAddress common.Address) (common.Address, error) {

    auth, err := updateNonce(auth, client, deployAddress)


    bridgeAssetAddr,_,_, err := bridgeAsset.DeployBridgeAsset(auth, client, mc)
    if err != nil {
        log.Error("error deploying bridge asset instance")
        return ZERO_ADDRESS, err
    }

    return bridgeAssetAddr, nil

}

func DeployLocal(port int, validators int, validatorThreshold *big.Int, depositThreshold *big.Int, deployAddress common.Address, mc uint8) (common.Address, common.Address, common.Address, common.Address, error){

    client, auth, validatorAddresses, err := AccountSetUp(port, validators)
    if err != nil {
        log.Error(err.Error())
        return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
    }

    
    recieverAddr, err := DeployReceiver(auth, client, validatorAddresses, depositThreshold, validatorThreshold, deployAddress)
    if err != nil {
        return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
    }

    emitterAddr, err := DeployEmitter(auth, client, deployAddress)
    if err != nil {
        return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
    }

    simpleEmitterAddress, err := DeploySimpleEmitter(auth, client, deployAddress)
    if err != nil {
        return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
    }

    bridgeAssetAddr, err := DeployBridgeAsset(auth, client, mc, deployAddress)
    if err != nil {
        return ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, ZERO_ADDRESS, err
    }

    return recieverAddr, emitterAddr, simpleEmitterAddress, bridgeAssetAddr, nil


}

