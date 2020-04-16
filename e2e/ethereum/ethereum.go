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
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Mintable"
	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const TestEthEndpoint = "ws://localhost:8545"

var TestTimeout = time.Second * 60

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

func DeployTestContracts(t *testing.T, id msg.ChainId) *utils.DeployedContracts {
	//deployPK string, chainID *big.Int, url string, relayers int, initialRelayerThreshold *big.Int, minCount uint8
	contracts, err := utils.DeployContracts(
		hexutil.Encode(AliceEthKp.Encode())[2:],
		uint8(id),
		TestEthEndpoint,
		big.NewInt(2), // Relayer threshold
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

func DeployErc20AndAddMinter(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, erc20Handler common.Address) common.Address { //nolint:unused,deadcode
	// Deploy
	opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	erc20Addr, _, erc20Instance, err := ERC20Mintable.DeployERC20Mintable(opts, client)
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

func CreateErc20Deposit(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, destId msg.ChainId, recipient []byte, amount *big.Int, contracts *utils.DeployedContracts, rId msg.ResourceId) {
	data := utils.ConstructErc20DepositData(rId, recipient, amount)

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

func WaitForEthereumEvent(t *testing.T, client *ethclient.Client, contract common.Address, subStr utils.EventSig) {
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
