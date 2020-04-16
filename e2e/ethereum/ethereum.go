// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	centrifugeHandler "github.com/ChainSafe/ChainBridge/bindings/CentrifugeAssetHandler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc20Mintable "github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const TestEthEndpoint = "ws://localhost:8545"

var TestTimeout = time.Second * 30

var log = log15.New("e2e", "ethereum")

var AliceEthKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var CharlieEthKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var CharlieEthAddr = common.HexToAddress(CharlieEthKp.Address())

func CreateConfig(key string, chain msg.ChainId, bridgeAddress, erc20HandlerAddress, genericHandler string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:         fmt.Sprintf("ethereum(%s)", key),
		Id:           chain,
		Endpoint:     TestEthEndpoint,
		From:         "",
		KeystorePath: key,
		Insecure:     true,
		Opts: map[string]string{
			"bridge":         bridgeAddress,
			"erc20Handler":   erc20HandlerAddress,
			"genericHandler": genericHandler,
		},
	}
}

func DeployTestContracts(t *testing.T, id msg.ChainId) *ethereum.DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := ethereum.DeployContracts(
		hexutil.Encode(AliceEthKp.Encode())[2:],
		uint8(id),
		TestEthEndpoint,
		3,             // Num relayers
		big.NewInt(2), // Relayer threshold
		0,             // unused?
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("====== Contracts =======")
	fmt.Printf("Bridge: %s\n", contracts.BridgeAddress.Hex())
	fmt.Printf("Erc20Handler: %s\n", contracts.ERC20HandlerAddress.Hex())
	fmt.Printf("(Generic) Centrifuge Handler: %s\n", contracts.CentrifugeHandlerAddress.Hex())
	fmt.Println("========================")
	return contracts
}

func CreateEthClient(t *testing.T) (*ethclient.Client, *bind.TransactOpts) {
	ctx := context.Background()
	rpcClient, err := rpc.DialWebsocket(ctx, TestEthEndpoint, "/ws")
	if err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpcClient)

	nonce, err := client.PendingNonceAt(ctx, CharlieEthAddr)
	if err != nil {
		t.Fatal(err)
	}
	opts := bind.NewKeyedTransactor(CharlieEthKp.PrivateKey())
	opts.Nonce = big.NewInt(int64(nonce - 1))        // -1 since we always increment before calling
	opts.Value = big.NewInt(0)                       // in wei
	opts.GasLimit = uint64(ethereum.DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(ethereum.DefaultGasPrice)
	opts.Context = ctx

	return client, opts
}

// deployMintApproveErc20 funds the account with a newly created erc20, to prepare for initiating a transfer
func DeployMintApproveErc20(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address) common.Address {
	// Deploy
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		t.Fatal(err)
	}

	// Mint
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Mint(opts, opts.From, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	// Approve
	log.Info("Approving tokens", "who", erc20Handler.Hex(), "amount", 99)
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = erc20Instance.Approve(opts, erc20Handler, big.NewInt(99))
	if err != nil {
		t.Fatal(err)
	}

	return erc20Addr
}

func DeployErc20AndAddMinter(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address) common.Address { //nolint:unused,deadcode
	// Deploy
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Addr, _, erc20Instance, err := erc20Mintable.DeployERC20Mintable(opts, client)
	if err != nil {
		t.Fatal(err)
	}
	log.Info("Deployed new ERC20 cotntract", "address", erc20Addr)

	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	_, err = erc20Instance.AddMinter(opts, erc20Handler)
	if err != nil {
		t.Fatal(err)
	}
	log.Info("Approving handler to mint ERC20 tokens", "handler", erc20Handler, "contract", erc20Addr)
	return erc20Addr
}

func RegisterErc20Resource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address, rId msg.ResourceId, addr common.Address) {
	log.Info("Registering resource", "id", rId.Hex(), "address", addr.Hex())
	instance, err := ERC20Handler.NewERC20Handler(erc20Handler, client)
	if err != nil {
		t.Fatal(err)
	}

	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = instance.SetResourceIDAndContractAddress(opts, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func RegisterGenericResource(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, genericHandler common.Address, rId msg.ResourceId, addr common.Address) {
	log.Info("Registering resource", "id", rId.Hex(), "address", addr.Hex())
	instance, err := centrifugeHandler.NewCentrifugeAssetHandler(genericHandler, client)
	if err != nil {
		t.Fatal(err)
	}

	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = instance.SetResourceIDAndContractAddress(opts, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

// constructErc20Data constructs the data field to be passed into a deposit call
func constructErc20DepositData(rId msg.ResourceId, id msg.ChainId, destRecipient []byte, amount *big.Int) []byte {
	var data []byte
	data = append(rId[:], math.PaddedBigBytes(amount, 32)...)
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...)
	data = append(data, destRecipient...)
	return data
}

func CreateErc20Deposit(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, srcId, destId msg.ChainId, recipient []byte, amount *big.Int, contracts *ethereum.DeployedContracts, rId msg.ResourceId) {
	data := constructErc20DepositData(rId, srcId, recipient, amount)

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	// Incrememnt Nonce by one
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	if _, err := bridgeInstance.Deposit(
		opts,
		uint8(destId),
		contracts.ERC20HandlerAddress,
		data,
	); err != nil {
		t.Fatal(err)
	}
}

func WaitForEthereumEvent(t *testing.T, client *ethclient.Client, contract common.Address, subStr ethereum.EventSig) {
	query := eth.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{subStr.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case evt := <-ch:
			log.Info("Got event, continuing...", "event", subStr, "topics", evt.Topics)
			sub.Unsubscribe()
			close(ch)
			return
		case err := <-sub.Err():
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(TestTimeout):
			t.Fatalf("Test timed out waiting for event %s", subStr)
		}
	}
}

func AssertHashExistence(t *testing.T, client *ethclient.Client, hash [32]byte, contract common.Address) {
	instance, err := centrifugeHandler.NewCentrifugeAssetHandler(contract, client)
	if err != nil {
		t.Fatal(err)
	}

	exists, err := instance.GetHash(&bind.CallOpts{}, hash)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("Hash does not exist")
	}
}

func AssertBalance(t *testing.T, client *ethclient.Client, amount *big.Int, erc20Contract, account common.Address) { //nolint:unused,deadcode
	instance, err := erc20Mintable.NewERC20Mintable(erc20Contract, client)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		t.Fatal(err)
	}
	if actual.Cmp(amount) != 0 {
		t.Fatalf("Balance mismatch. Expected: %s Got: %s", amount.String(), actual.String())
	}
}
