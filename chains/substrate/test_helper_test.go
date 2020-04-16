// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"os"
	"testing"

	"github.com/ChainSafe/ChainBridge/keystore"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

const TestEndpoint = "ws://127.0.0.1:9944"

const TestRelayerThreshold = 2
const TestChainId = 1

var AliceKey = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey].AsKeyringPair()
var BobKey = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey].AsKeyringPair()

var TestLogLevel = log15.LvlTrace
var AliceTestLogger = newTestLogger("Alice")
var BobTestLogger = newTestLogger("Bob")

func TestMain(m *testing.M) {
	ensureInitializedChain()
	os.Exit(m.Run())
}

func ensureInitializedChain() {
	conn := NewConnection(TestEndpoint, "Test", AliceKey, AliceTestLogger)
	err := conn.Connect()
	if err != nil {
		panic(err)
	}
	var exists types.Bool
	_, err = conn.queryStorage("Bridge", "Relayers", AliceKey.PublicKey, nil, &exists)
	if err != nil {
		panic(err)
	}

	if !exists {
		meta := conn.getMetadata()
		call, err := types.NewCall(&meta, string(utils.AddRelayerMethod), types.NewAccountID(AliceKey.PublicKey))
		if err != nil {
			panic(err)
		}
		err = conn.SubmitTx(utils.SudoMethod, call)
		if err != nil {
			panic(err)
		}

		call, err = types.NewCall(&meta, string(utils.AddRelayerMethod), types.NewAccountID(BobKey.PublicKey))
		if err != nil {
			panic(err)
		}
		err = conn.SubmitTx(utils.SudoMethod, call)
		if err != nil {
			panic(err)
		}

		call, err = types.NewCall(&meta, string(utils.SetThresholdMethod), types.U32(2))
		if err != nil {
			panic(err)
		}
		err = conn.SubmitTx(utils.SudoMethod, call)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("=========================================================================")
		fmt.Println("! WARNING: Running tests against an initialized chain, results may vary !")
		fmt.Println("=========================================================================")
	}
	conn.Close()
}

func newTestLogger(name string) log15.Logger {
	tLog := log15.Root().New("chain", name)
	tLog.SetHandler(log15.LvlFilterHandler(TestLogLevel, tLog.GetHandler()))
	return tLog
}

// createAliceConnection creates and starts a connection with the Alice keypair
func createAliceConnection(t *testing.T) *Connection {
	alice := NewConnection(TestEndpoint, "Alice", AliceKey, AliceTestLogger)
	err := alice.Connect()
	if err != nil {
		t.Fatal(err)
	}
	return alice
}

// createAliceAndBobConnections creates and calls `Connect()` on two Connections using the Alice and Bob keypairs
func createAliceAndBobConnections(t *testing.T) (*Connection, *Connection) {
	alice := createAliceConnection(t)

	bob := NewConnection(TestEndpoint, "Bob", BobKey, AliceTestLogger)
	err := bob.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return alice, bob
}

// getFreeBalance queries the balance for an account, storing the result in `res`
func getFreeBalance(c *Connection, res *types.U128) {
	var acct utils.AccountData

	ok, err := c.queryStorage("System", "Account", c.key.PublicKey, nil, &acct)
	if err != nil {
		panic(err)
	} else if !ok {
		panic("no account data")
	}
	*res = acct.Data.Free
}
