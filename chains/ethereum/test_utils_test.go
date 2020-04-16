// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func newTestLogger(name string) log15.Logger {
	tLog := log15.New("chain", name)
	tLog.SetHandler(log15.LvlFilterHandler(log15.LvlTrace, tLog.GetHandler()))
	return tLog
}

func newLocalConnection(t *testing.T, cfg *Config) *Connection {
	kp := keystore.TestKeyRing.EthereumKeys[cfg.from]
	conn := NewConnection(cfg, kp, TestLogger)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func deployTestContracts(t *testing.T, id msg.ChainId) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		hexutil.Encode(AliceKp.Encode())[2:],
		uint8(id),
		TestEndpoint,
		TestRelayerThreshold,
	)
	if err != nil {
		t.Fatal(err)
	}

	return contracts
}

// createErc20Deposit deploys a new erc20 token contract mints, the sender (based on value), and creates a deposit
func createErc20Deposit(contract *Bridge.Bridge,
	txOpts *bind.TransactOpts,
	rId msg.ResourceId,
	originHandler,
	destRecipient common.Address,
	destId msg.ChainId,
	amount *big.Int) error {

	data := utils.ConstructErc20DepositData(rId, destRecipient.Bytes(), amount)

	// Incrememnt Nonce by one
	txOpts.Nonce = txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1))
	if _, err := contract.Deposit(
		txOpts,
		uint8(destId),
		originHandler,
		data,
	); err != nil {
		return err
	}
	return nil
}
