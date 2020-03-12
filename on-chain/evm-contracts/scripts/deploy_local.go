package scripts

import (
    "flag"
    "os"
    "context"
    "crypto/ecdsa"
	"math/big"
    "strconv"
    "fmt"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	bridgeAsset "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
    simpleEmitter "github.com/ChainSafe/ChainBridgeV2/contracts/SimpleEmitter"
    erc20 "github.com/ChainSafe/ChainBridgeV2/contracts/ERC20Mintable"
    erc721 "github.com/ChainSafe/ChainBridgeV2/contracts/ERC721Mintable"
    log "github.com/ChainSafe/log15"

)


var (
	RECEIVER_ADDRESS = "0x705D4Fa884AF2Ae59C7780A0f201109947E2Bf6D";
	CENTRIFUGE_ADDRESS = "0x290f41e61374c715C1127974bf08a3993afd0145";
	EMITTER_ADDRESS = "0x3c747684333605408F9A4907DA043ee4c1A72D9c";
	TEST_EMITTER_ADDRESS = "0x8090062239c909eB9b0433F1184c7DEf6124cc78";

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
)


// dunno if the cli should look like this
// TODO: Use The Library
func main() {
	validators := flag.Int("validators", 2, "Number of validators")
	validatorThreshold := flag.Int("validator-threshold", 2, "Value of validator threshold")
	depositThreshold := flag.Int("deposit-threshold", 1, "Value of deposit threshold")
	port := flag.Int("port", 8545, "Port of RPC instance")
	depositERC20 := flag.Bool("deposit-erc", false, "Make an ERC20 deposit")
	depositNFT := flag.Bool("deposit-nft", false, "Make an ERC721 deposit")
	depositAsset := flag.Bool("deposit-asset", false, "Make a test deployment")
	testOnly := flag.Bool("test-only", false, "Skip main contract depoyments, only run tests")
	dest := flag.Int("dest", 1, "destination chain")

    err := deploy_local(*validators, *validatorThreshold, *depositThreshold, *port, *depositERC20, *depositNFT, *depositAsset, *testOnly, *dest)
    if err != nil {
        log.Error(err.Error())
		os.Exit(1)
    }

}


func createValidatorSlice(valAddr []string, numValidators int) []common.Address {
    
    validatorAddresses := make([]common.Address, numValidators)

    for i := 0; i < numValidators; i++ {
        validatorAddresses[i] = common.HexToAddress(valAddr[i])

    }

    return validatorAddresses
}


func deploy_local(validators int, validatorThreshold int, depositThreshold int, port int, erc bool, nft bool, asset bool, test bool, dest int) error {
    client, err := ethclient.Dial("https://localhost:"+strconv.Itoa(port))
    if err != nil {
        log.Error(err.Error())
        return err
    }

    privateKey, err := crypto.HexToECDSA(DEPLOYER_PRIV_KEY)
    if err != nil {
        log.Error(err.Error())
        return err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Error(err.Error())
        return err
    }

    deployAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), deployAddress)
    if err != nil {
        log.Error(err.Error())
        return err
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Error(err.Error())
        return err
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)    
    auth.GasLimit = uint64(300000) 
    auth.GasPrice = gasPrice

    validatorAddresses := createValidatorSlice(VALIDATOR_ADDRESS, validators)
    valThres := big.NewInt(int64(validatorThreshold))
    depThres := big.NewInt(int64(depositThreshold))


    if !test {
        _, _, _, err := receiver.DeployReceiver(auth, client, validatorAddresses, depThres,valThres)
        if err != nil {
            log.Error("error deploying reciever instance")
            return err
        }
        emitterAddr, _, emitterInstance, err := emitter.DeployEmitter(auth, client)
        if err != nil {
            log.Error("error deploying emitter instance")
            return err
        }
        _, _, _, err = simpleEmitter.DeploySimpleEmitter(auth, client)
        if err != nil {
            log.Error("error deploying simple emitter instance")
            return err
        }
        _,_,_, err = bridgeAsset.DeployBridgeAsset(auth, client, 10)
        if err != nil {
            log.Error("error deploying bridge asset instance")
            return err
        }

        if erc {

            erc20Addr, erc20Tx, erc20Instance, err := erc20.DeployERC20Mintable(auth, client)
            log.Info("[ERC20 Transfer] deployed token!")
    
            if err != nil {
                log.Error("error deploying ERC20 instance")
                return err
            }
    
            priKey, err := crypto.GenerateKey()
            if err != nil {
                log.Error(err.Error())
                return err
            }
        
            pubKey := privateKey.Public()
            pubKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
            if !ok {
                log.Error(err.Error())
                return err
            }
        
            // Mint and Approve tokens
            mintAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
    
            _, err = erc20Instance.Mint(auth, mintAddress, big.NewInt(100))
            if err != nil {
                log.Error(err.Error())
                return err
            }

            log.Info("[ERC20 Transfer] Minted tokens!")
    
            _, err = erc20Instance.Approve(auth, emitterAddr, big.NewInt(1))
            if err != nil {
                log.Error(err.Error())
                return err
            }

            log.Info("[ERC20 Transfer] Approved tokens!")

            // pre transfer balance
            prebal, err := emitterInstance.Balances(&bind.CallOpts{}, erc20Addr)
            if err != nil {
                log.Error(err.Error())
                return err
            }


            // make deposit
            _, err = emitterInstance.DepositGenericErc(auth, big.NewInt(0), big.NewInt(1), validatorAddresses[1], erc20Addr)
            if err != nil {
                log.Error(err.Error())
                return err
            }
            log.Info("[ERC20 Transfer] Created Deposit!")

            postbal, err := emitterInstance.Balances(&bind.CallOpts{}, erc20Addr)

            log.Info(strconv.FormatBool(prebal.Cmp(postbal)== 1))

            
    
        } 
        

    }



    return nil
	
}
