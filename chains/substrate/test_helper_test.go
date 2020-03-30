// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"testing"

	"github.com/ChainSafe/ChainBridge/keystore"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

var TestEndpoint = "ws://127.0.0.1:9944"

var AliceKey = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey].AsKeyringPair()
var BobKey = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey].AsKeyringPair()

var TestLogger = newTestLogger()

func newTestLogger() log15.Logger {
	tLog := log15.New("test_chain", "substrate")
	tLog.SetHandler(log15.LvlFilterHandler(log15.LvlInfo, tLog.GetHandler()))
	return tLog
}

// createAliceConnection creates and starts a connection with the Alice keypair
func createAliceConnection(t *testing.T) *Connection {
	alice := NewConnection(TestEndpoint, "Alice", AliceKey, TestLogger)
	err := alice.Connect()
	if err != nil {
		t.Fatal(err)
	}
	return alice
}

// createAliceAndBobConnections creates and calls `Connect()` on two Connections using the Alice and Bob keypairs
func createAliceAndBobConnections(t *testing.T) (*Connection, *Connection) {
	alice := createAliceConnection(t)

	bob := NewConnection(TestEndpoint, "Bob", BobKey, TestLogger)
	err := bob.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return alice, bob
}

// getFreeBalance queries the balance for an account, storing the result in `res`
func getFreeBalance(c *Connection, res *types.U128) {
	var acct accountData

	ok, err := c.queryStorage("System", "Account", c.key.PublicKey, nil, &acct)
	if err != nil {
		panic(err)
	} else if !ok {
		panic("no account data")
	}
	*res = acct.Data.Free
}
