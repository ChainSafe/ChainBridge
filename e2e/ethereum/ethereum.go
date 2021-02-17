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
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var TestTimeout = time.Second * 30

const EthAEndpoint = "ws://localhost:8545"
const EthBEndpoint = "ws://localhost:8546"

var log = log15.New("e2e", "ethereum")

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var CharlieKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var DaveKp = keystore.TestKeyRing.EthereumKeys[keystore.DaveKey]
var EveKp = keystore.TestKeyRing.EthereumKeys[keystore.EveKey]

type TestContext struct {
	BaseContracts *utils.DeployedContracts // All the contracts required for the bridge
	TestContracts TestContracts            // Additional contracts for tests (eg. erc contracts)
	Client        *utils.Client
}

type TestContracts struct {
	Erc20Sub      common.Address // Contract configured for substrate erc20 transfers
	Erc20Eth      common.Address // Contact configured for eth to eth erc20 transfer
	Erc721Sub     common.Address // Contract configured for substrate erc721 transfers
	Erc721Eth     common.Address // Contract configured for eth to eth erc721 transfer
	AssetStoreSub common.Address // Contract configure for substrate generic transfer
	AssetStoreEth common.Address // Contract configured for eth to eth generic transfer
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
			"bridge":             contracts.BridgeAddress.String(),
			"erc20Handler":       contracts.ERC20HandlerAddress.String(),
			"erc721Handler":      contracts.ERC721HandlerAddress.String(),
			"genericHandler":     contracts.GenericHandlerAddress.String(),
			"blockConfirmations": "3",
		},
	}
}

func DeployTestContracts(t *testing.T, client *utils.Client, endpoint string, id msg.ChainId, threshold *big.Int) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		client,
		uint8(id),
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
	fmt.Printf("Erc721Handler:		%s\n", contracts.ERC721HandlerAddress.Hex())
	fmt.Printf("GenericHandler: 		%s\n", contracts.GenericHandlerAddress.Hex())
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

	id, err := client.ChainID(ctx)
	if err != nil {
		t.Fatal(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(kp.PrivateKey(), id)
	if err != nil {
		t.Fatal(err)
	}
	opts.Nonce = big.NewInt(int64(nonce - 1))        // -1 since we always increment before calling
	opts.Value = big.NewInt(0)                       // in wei
	opts.GasLimit = uint64(ethereum.DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(ethereum.DefaultGasPrice)
	opts.Context = ctx

	return client, opts
}

func CreateErc20Deposit(t *testing.T, client *utils.Client, destId msg.ChainId, recipient []byte, amount *big.Int, contracts *utils.DeployedContracts, rId msg.ResourceId) {
	data := utils.ConstructErc20DepositData(recipient, amount)

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client.Client)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := bridgeInstance.Deposit(
		client.Opts,
		uint8(destId),
		rId,
		data,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = utils.WaitForTx(client, tx)
	if err != nil {
		t.Fatal(err)
	}

}

func CreateErc721Deposit(t *testing.T, client *utils.Client, destId msg.ChainId, recipient []byte, tokenId *big.Int, contracts *utils.DeployedContracts, rId msg.ResourceId) {
	data := utils.ConstructErc721DepositData(tokenId, recipient)

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client.Client)
	if err != nil {
		t.Fatal(err)
	}

	err = utils.UpdateNonce(client)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := bridgeInstance.Deposit(
		client.Opts,
		uint8(destId),
		rId,
		data,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = utils.WaitForTx(client, tx)
	if err != nil {
		t.Fatal(err)
	}
}

func CreateGenericDeposit(t *testing.T, client *utils.Client, destId msg.ChainId, metadata []byte, contracts *utils.DeployedContracts, rId msg.ResourceId) {
	data := utils.ConstructGenericDepositData(metadata)

	bridgeInstance, err := bridge.NewBridge(contracts.BridgeAddress, client.Client)
	if err != nil {
		t.Fatal(err)
	}

	err = utils.UpdateNonce(client)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := bridgeInstance.Deposit(
		client.Opts,
		uint8(destId),
		rId,
		data,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = utils.WaitForTx(client, tx)
	if err != nil {
		t.Fatal(err)
	}
}

func WaitForProposalActive(t *testing.T, client *utils.Client, bridge common.Address, nonce uint64) {
	startBlock := ethtest.GetLatestBlock(t, client)

	query := eth.FilterQuery{
		FromBlock: startBlock,
		Addresses: []common.Address{bridge},
		Topics: [][]common.Hash{
			{utils.ProposalEvent.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.Client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()
	timeout := time.After(TestTimeout)
	for {
		select {
		case evt := <-ch:
			currentNonce := evt.Topics[2].Big()
			status := uint8(evt.Topics[3].Big().Uint64())
			// Check nonce matches
			if utils.IsActive(status) && currentNonce.Cmp(big.NewInt(int64(nonce))) == 0 {
				log.Info("Got matching ProposalCreated event, continuing...", "nonce", currentNonce, "topics", evt.Topics)
				return
			} else {
				log.Info("Incorrect ProposalCreated event", "nonce", currentNonce, "expectedNonce", nonce, "topics", evt.Topics)
			}
		case err := <-sub.Err():
			if err != nil {
				t.Fatal(err)
			}
		case <-timeout:
			t.Fatalf("Test timed out waiting for ProposalCreated event")
		}
	}
}

func WaitForProposalExecutedEvent(t *testing.T, client *utils.Client, bridge common.Address, nonce uint64) {
	startBlock := ethtest.GetLatestBlock(t, client)

	query := eth.FilterQuery{
		FromBlock: startBlock,
		Addresses: []common.Address{bridge},
		Topics: [][]common.Hash{
			{utils.ProposalEvent.GetTopic()},
		},
	}

	ch := make(chan ethtypes.Log)
	sub, err := client.Client.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Unsubscribe()

	timeout := time.After(TestTimeout)
	for {
		select {
		case evt := <-ch:
			currentNonce := evt.Topics[2].Big()
			status := uint8(evt.Topics[3].Big().Uint64())
			// Check nonce matches
			if utils.IsExecuted(status) && currentNonce.Cmp(big.NewInt(int64(nonce))) == 0 {
				log.Info("Got matching ProposalExecuted event, continuing...", "nonce", currentNonce, "topics", evt.Topics)
				return
			} else {
				log.Info("Incorrect ProposalExecuted event", "nonce", currentNonce, "expectedNonce", nonce, "topics", evt.Topics)
			}
		case err := <-sub.Err():
			if err != nil {
				t.Fatal(err)
			}
		case <-timeout:
			t.Fatalf("Test timed out waiting for ProposalExecuted event")
		}
	}
}
