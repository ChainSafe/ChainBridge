// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/crypto/secp256k1"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
)

const TestEndpoint = "ws://localhost:8545"

var TestLogger = newTestLogger("test")
var TestTimeout = time.Second * 10

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]

var TestRelayerThreshold = big.NewInt(2)

var aliceTestConfig = &Config{
	id:       msg.ChainId(0),
	name:     "alice",
	endpoint: TestEndpoint,
	from:     keystore.AliceKey,
	gasLimit: big.NewInt(DefaultGasLimit),
	gasPrice: big.NewInt(DefaultGasPrice),
}

var bobTestConfig = &Config{
	id:       msg.ChainId(0),
	name:     "bob",
	endpoint: TestEndpoint,
	from:     keystore.BobKey,
	gasLimit: big.NewInt(DefaultGasLimit),
	gasPrice: big.NewInt(DefaultGasPrice),
}

var charlieTestConfig = &Config{
	id:       msg.ChainId(0),
	name:     "charlie",
	endpoint: TestEndpoint,
	from:     keystore.CharlieKey,
	gasLimit: big.NewInt(DefaultGasLimit),
	gasPrice: big.NewInt(DefaultGasPrice),
}

func newTestLogger(name string) log15.Logger {
	tLog := log15.New("chain", name)
	tLog.SetHandler(log15.LvlFilterHandler(log15.LvlTrace, tLog.GetHandler()))
	return tLog
}

func newLocalConnection(t *testing.T, cfg *Config) *Connection {
	kp := keystore.TestKeyRing.EthereumKeys[cfg.from]
	conn := NewConnection(cfg, kp, TestLogger, make(chan int))
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func deployTestContracts(t *testing.T, client *utils.Client, id msg.ChainId, kp *secp256k1.Keypair) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		client,
		uint8(id),
		TestEndpoint,
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
	handler,
	destRecipient common.Address,
	destId msg.ChainId,
	amount *big.Int,
) {

	data := utils.ConstructErc20DepositData(rId, destRecipient.Bytes(), amount)

	// Incrememnt Nonce by one
	client.Opts.Nonce = client.Opts.Nonce.Add(client.Opts.Nonce, big.NewInt(1))
	if _, err := contract.Deposit(
		client.Opts,
		uint8(destId),
		handler,
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
	handler,
	destRecipient common.Address,
	destId msg.ChainId,
	tokenId *big.Int,
	metadata []byte,
) {

	data := utils.ConstructErc721DepositData(rId, tokenId, destRecipient.Bytes())

	// Incrememnt Nonce by one
	client.Opts.Nonce = client.Opts.Nonce.Add(client.Opts.Nonce, big.NewInt(1))
	if _, err := bridge.Deposit(
		client.Opts,
		uint8(destId),
		handler,
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
	handler common.Address,

	destId msg.ChainId,
	hash []byte) {

	data := utils.ConstructGenericDepositData(rId, hash)

	// Incrememnt Nonce by one
	client.Opts.Nonce = client.Opts.Nonce.Add(client.Opts.Nonce, big.NewInt(1))
	if _, err := bridge.Deposit(
		client.Opts,
		uint8(destId),
		handler,
		data,
	); err != nil {
		t.Fatal(err)
	}
}
