package chains

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/chains/substrate"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const TestEthEndpoint = "http://localhost:8545"
const TestSubEndpoint = "http://localhost:9944"
const EthChainId = msg.ChainId(0)
const SubChainId = msg.ChainId(1)

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var AliceAddr = common.HexToAddress(AliceKp.Address())
var BobAddr = common.HexToAddress(BobKp.Address())

func deployTestContracts(t *testing.T, id msg.ChainId) *ethereum.DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := ethereum.DeployContracts(
		hexutil.Encode(AliceKp.Encode())[2:],
		big.NewInt(int64(id)),
		TestEthEndpoint,
		2, // Num relayers
		2, // Relayer threshold
		0, // unused?
	)
	if err != nil {
		t.Fatal(err)
	}

	return contracts
}

func createEthClient(t *testing.T) (*ethclient.Client, *bind.TransactOpts) {
	ctx := context.Background()
	rpcClient, err := rpc.DialWebsocket(ctx, TestEthEndpoint, "/ws")
	if err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpcClient)

	nonce, err := client.PendingNonceAt(ctx, AliceAddr)
	opts := bind.NewKeyedTransactor(AliceKp.PrivateKey())
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)               // in wei
	opts.GasLimit = uint64(ethereum.DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(ethereum.DefaultGasPrice)
	opts.Context = ctx

	return client, opts
}

func deployApproveErc20(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address) common.Address {


	// Deploy
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		t.Fatal(err)
	}

	// Approve
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Approve(opts, erc20Handler, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	return erc20Addr
}

func createEthConfig(key, bridgeAddress, erc20HandlerAddress string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("ethereum(%s)", key),
		Id:           EthChainId,
		Endpoint:     TestEthEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts:         map[string]string{
			"contract": bridgeAddress,
			"erc20Handler": erc20HandlerAddress,
		},
	}
}

func createSubConfig(key, bridgeAddress, erc20HandlerAddress string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("substrate(%s)", key),
		Id:           SubChainId,
		Endpoint:    TestSubEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts:         map[string]string{
			"contract": bridgeAddress,
			"erc20Handler": erc20HandlerAddress,
		},
	}
}

func createBridge(t *testing.T, name, bridgeAddress, erc20HandlerAddres string) *core.Core {
	eth, err := ethereum.InitializeChain(createEthConfig(name, bridgeAddress, erc20HandlerAddres))
	if err != nil {
		t.Fatal(err)
	}
	sub, err := substrate.InitializeChain(createSubConfig(name, bridgeAddress, erc20HandlerAddres))
	if err != nil {
		t.Fatal(err)
	}

	bridge := core.NewCore()
	bridge.AddChain(eth)
	bridge.AddChain(sub)

	return bridge
}

func createErc20Deposit(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, contracts ethereum.DeployedContracts, erc20Contract common.Address) {
	data := constructDataBytes(erc20Contract, destHandler, destTokenAddress, destRecipient, amount)

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	// Incrememnt Nonce by one
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	if _, err := bridgeInstance.Deposit(
		opts,
		big.NewInt(int64(SubChainId)),
		contracts.ERC20HandlerAddress,
		data,
	); err != nil {
		t.Fatal(err)
	}
}

func TestErc20ToSubstrate(t *testing.T) {

	// Deploy contracts, mint, approve

	// Creates an alice client we can use to initiate everything
	client, opts := createEthClient(t)
	contracts := deployTestContracts(t, EthChainId)
	erc20Contract := deployApproveErc20(t, client, opts, contracts.ERC20HandlerAddress)
	// Create two bridges with chains
	bridgeA := createBridge(t, "A", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex())
	bridgeB := createBridge(t, "A", contracts.BridgeAddress.Hex(), contracts.ERC20HandlerAddress.Hex())

	// Initiate transfer
	createErc20Deposit(t, erc20Contract)

	// Check for success event

}