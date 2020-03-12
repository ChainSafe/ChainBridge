package scripts

import (
	"flag"
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
	"math/big"
	"strconv"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	erc20 "github.com/ChainSafe/ChainBridgeV2/contracts/ERC20"
	bridgeAsset "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	reciever "github.com/ChainSafe/ChainBridgeV2/contracts/Reciever"
	simpleEmitter "github.com/ChainSafe/ChainBridgeV2/contracts/SimpleEmitter"

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

	deploy_local(validators, validatorThreshold, depositThreshold, port, depositERC20, depositNFT, depositAsset, testOnly, dest)

}


func deploy_local(validators *int, validatorThreshold *int, depositThreshold *int, port *int, erc20 *bool, nft *bool, asset *bool, test *bool, dest *int) {
    client, err := ethclient.Dial("https://localhost:"+*port)
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(DEPLOYER_PRIV_KEY)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    deployAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), deployAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)    
    auth.GasLimit = uint64(300000) 
    auth.GasPrice = gasPrice



	_ = instance
	
	if !testOnly {
		recieverAddress, recieverTx, recieverInstance, recieverErr := deployReciever()
		emitterAddress, emitterTx, emitterInstance, emitterErr := deployEmitter()
		deploySimpleEmitter()
		deployBridgeAsset()
	}
}

func deployReciever(auth *bind.TransactOpts, client *bind.ContractBackend) (common.Address, *types.Transaction, *Receiver, error) {
    //_addrs []common.Address, _depositThreshold *big.Int, _validatorThreshold *big.Int
    return reciever.deployReciever(auth, client,)


}

func deployEmitter(auth *bind.TransactOpts, client *bind.ContractBackend) {}
func deploySimpleEmitter(auth *bind.TransactOpts, client *bind.ContractBackend) {}

func deployBridgeAsset(auth *bind.TransactOpts, client *bind.ContractBackend) {}