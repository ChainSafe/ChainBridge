// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
)

const TestEndpoint = "ws://localhost:8545"

var TestLogger = newTestLogger("test")
var TestTimeout = time.Second * 30

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]

var TestNumRelayers = 2
var TestRelayerThreshold = big.NewInt(2)
var TestMintAmount = big.NewInt(1000)

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
