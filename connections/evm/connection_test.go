// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package evm

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridge/keystore"
	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/log15"
)

const TestEndpoint = "ws://localhost:8545"

var TestConfig = &Config{
	Id:       msg.ChainId(0),
	Name:     "alice",
	Endpoint: TestEndpoint,
	From:     keystore.AliceKey,
	GasLimit: big.NewInt(DefaultGasLimit),
	GasPrice: big.NewInt(DefaultGasPrice),
}

func newTestLogger(name string) log15.Logger {
	tLog := log15.New("connection", name)
	tLog.SetHandler(log15.LvlFilterHandler(log15.LvlTrace, tLog.GetHandler()))
	return tLog
}

var testLog = newTestLogger("evm")

func TestNewConnection(t *testing.T) {
	stop_chan := make(chan int)
	connect := NewConnection(TestConfig, keystore.TestKeyRing.EthereumKeys[TestConfig.From], testLog, stop_chan)
	expected := &Connection{
		Ctx:  context.Background(),
		Cfg:  *TestConfig,
		Kp:   keystore.TestKeyRing.EthereumKeys[TestConfig.From],
		log:  testLog,
		Stop: stop_chan,
	}

	if !reflect.DeepEqual(expected, connect) {
		t.Fatalf("Mismatch.\n\tExpected: %#v\n\tGot:%#v", expected, connect)
	}
}

//TODO: Make this test more generic when we have more ethereum connections
func TestConnect(t *testing.T) {

	connect := NewConnection(TestConfig, keystore.TestKeyRing.EthereumKeys[TestConfig.From], testLog, make(chan int))
	err := connect.Connect()
	if err != nil {
		t.Fatalf("Unable to connect to connection: %s", err)
	}

	connect.Close()
}
