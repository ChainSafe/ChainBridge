// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/chains/ethereum"
	"github.com/ChainSafe/ChainBridge/core"
	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
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

var TestTimeout = time.Second * 30

var log = log15.New("e2e", "ethereum")

// TODO: Remove extra addrs vars when PR #339 lands
var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var AliceAddr = common.HexToAddress(AliceKp.Address())
var CharlieKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var CharlieAddr = common.HexToAddress(CharlieKp.Address())

type TestContext struct {
	BaseContracts *utils.DeployedContracts // All the contracts required for the bridge
	TestContracts TestContracts            // Additional contracts for tests (eg. erc contracts)
	Client        *ethclient.Client
	Opts          *bind.TransactOpts
}

type TestContracts struct {
	Erc20Sub common.Address // Contract configured for substrate transfers
	Erc20Eth common.Address // Contact configured for eth to eth transfer
}

func CreateConfig(key string, chain msg.ChainId, contracts *utils.DeployedContracts, endpoint string) *core.ChainConfig {
	return &core.ChainConfig{
		Name:           fmt.Sprintf("ethereum(%s,%d)", key, chain),
		Id:             chain,
		Endpoint:       endpoint,
		From:           "",
		KeystorePath:   key,
		Insecure:       true,
		FreshStart:     true,
		BlockstorePath: os.TempDir(),
		Opts: map[string]string{
			"bridge":         contracts.BridgeAddress.String(),
			"erc20Handler":   contracts.ERC20HandlerAddress.String(),
			"genericHandler": contracts.CentrifugeHandlerAddress.String(),
		},
	}
}

func DeployTestContracts(t *testing.T, endpoint string, id msg.ChainId, threshold *big.Int) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		hexutil.Encode(AliceKp.Encode())[2:],
		uint8(id),
		endpoint,
		threshold,
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("======================== Contracts Deployed ========================")
	fmt.Printf("Endpoint:			%s\n", endpoint)
	fmt.Printf("Deployer:			%s\n", AliceKp.Address())
	fmt.Printf("Chain ID:			%d\n", id)
	fmt.Printf("Relayer Threshold:	%s\n", threshold.String())
	fmt.Printf("Bridge:				%s\n", contracts.BridgeAddress.Hex())
	fmt.Printf("Erc20Handler:		%s\n", contracts.ERC20HandlerAddress.Hex())
	fmt.Printf("Generic/Centrifuge Handler: %s\n", contracts.CentrifugeHandlerAddress.Hex())
	fmt.Println("====================================================================")
	return contracts
}

func CreateEthClient(t *testing.T, endpoint string, kp *secp256k1.Keypair) (*ethclient.Client, *bind.TransactOpts) {
	ctx := context.Background()
	rpcClient, err := rpc.DialWebsocket(ctx, endpoint, "/ws")
	if err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpcClient)

	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(kp.Address()))
	if err != nil {
		t.Fatal(err)
	}
	opts := bind.NewKeyedTransactor(kp.PrivateKey())
	opts.Nonce = big.NewInt(int64(nonce - 1))        // -1 since we always increment before calling
	opts.Value = big.NewInt(0)                       // in wei
	opts.GasLimit = uint64(ethereum.DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(ethereum.DefaultGasPrice)
	opts.Context = ctx

	return client, opts
}

func CreateErc20Deposit(t *testing.T, client *ethclient.Client, opts *bind.TransactOpts, destId msg.ChainId, recipient []byte, amount *big.Int, contracts *utils.DeployedContracts, rId msg.ResourceId) {
	data := utils.ConstructErc20DepositData(rId, recipient, amount)

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client)
	if err != nil {
		t.Fatal(err)
	}

	err = utils.UpdateNonce(opts, client)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := bridgeInstance.Deposit(
		opts,
		uint8(destId),
		contracts.ERC20HandlerAddress,
		data,
	); err != nil {
		t.Fatal(err)
	}
}

func WaitForEthereumEvent(t *testing.T, client *ethclient.Client, contract common.Address, subStr utils.EventSig, startBlock *big.Int) {
	query := eth.FilterQuery{
		FromBlock: startBlock,
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
