// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

const TestEndpoint = "ws://localhost:8545"

var TestLogger = newTestLogger("test")
var TestTimeout = time.Second * 30

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]

var TestRelayerThreshold = big.NewInt(2)
var TestChainId = msg.ChainId(0)

var aliceTestConfig = createConfig("alice", nil, nil)

func createConfig(name string, startBlock *big.Int, contracts *utils.DeployedContracts) *Config {
	cfg := &Config{
		name:                   name,
		id:                     0,
		endpoint:               TestEndpoint,
		from:                   name,
		keystorePath:           "",
		blockstorePath:         "",
		freshStart:             true,
		bridgeContract:         common.Address{},
		erc20HandlerContract:   common.Address{},
		erc721HandlerContract:  common.Address{},
		genericHandlerContract: common.Address{},
		gasLimit:               big.NewInt(DefaultGasLimit),
		maxGasPrice:            big.NewInt(DefaultGasPrice),
		gasMultiplier:          big.NewFloat(DefaultGasMultiplier),
		http:                   false,
		startBlock:             startBlock,
		blockConfirmations:     big.NewInt(3),
	}

	if contracts != nil {
		cfg.bridgeContract = contracts.BridgeAddress
		cfg.erc20HandlerContract = contracts.ERC20HandlerAddress
		cfg.erc721HandlerContract = contracts.ERC721HandlerAddress
		cfg.genericHandlerContract = contracts.GenericHandlerAddress
	}

	return cfg
}

func newTestLogger(name string) log15.Logger {
	tLog := log15.New("chain", name)
	tLog.SetHandler(log15.LvlFilterHandler(log15.LvlError, tLog.GetHandler()))
	return tLog
}

func newLocalConnection(t *testing.T, cfg *Config) *connection.Connection {
	kp := keystore.TestKeyRing.EthereumKeys[cfg.from]
	conn := connection.NewConnection(TestEndpoint, false, kp, TestLogger, big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice), big.NewInt(DefaultMinGasPrice), big.NewFloat(DefaultGasMultiplier), "", "")
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func deployTestContracts(t *testing.T, client *utils.Client, id msg.ChainId) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		client,
		uint8(id),
		TestRelayerThreshold,
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("=======================================================")
	fmt.Printf("Bridge: %s\n", contracts.BridgeAddress.Hex())
	fmt.Printf("Erc20Handler: %s\n", contracts.ERC20HandlerAddress.Hex())
	fmt.Printf("ERC721Handler: %s\n", contracts.ERC721HandlerAddress.Hex())
	fmt.Printf("GenericHandler: %s\n", contracts.GenericHandlerAddress.Hex())
	fmt.Println("========================================================")

	return contracts
}

func createErc20Deposit(
	t *testing.T,
	contract *Bridge.Bridge,
	client *utils.Client,
	rId msg.ResourceId,
	destRecipient common.Address,
	destId msg.ChainId,
	amount *big.Int,
) {

	data := utils.ConstructErc20DepositData(destRecipient.Bytes(), amount)

	// Incrememnt Nonce by one
	client.Opts.Nonce = client.Opts.Nonce.Add(client.Opts.Nonce, big.NewInt(1))
	if _, err := contract.Deposit(
		client.Opts,
		uint8(destId),
		rId,
		data,
	); err != nil {
		t.Fatal(err)
	}
}

func createErc721Deposit(
	t *testing.T,
	bridge *Bridge.Bridge,
	client *utils.Client,
	rId msg.ResourceId,
	destRecipient common.Address,
	destId msg.ChainId,
	tokenId *big.Int,
) {

	data := utils.ConstructErc721DepositData(tokenId, destRecipient.Bytes())

	// Incrememnt Nonce by one
	client.Opts.Nonce = client.Opts.Nonce.Add(client.Opts.Nonce, big.NewInt(1))
	if _, err := bridge.Deposit(
		client.Opts,
		uint8(destId),
		rId,
		data,
	); err != nil {
		t.Fatal(err)
	}
}

func createGenericDeposit(
	t *testing.T,
	bridge *Bridge.Bridge,
	client *utils.Client,
	rId msg.ResourceId,
	destId msg.ChainId,
	hash []byte) {

	data := utils.ConstructGenericDepositData(hash)

	// Incrememnt Nonce by one
	client.Opts.Nonce = client.Opts.Nonce.Add(client.Opts.Nonce, big.NewInt(1))
	if _, err := bridge.Deposit(
		client.Opts,
		uint8(destId),
		rId,
		data,
	); err != nil {
		t.Fatal(err)
	}
}
