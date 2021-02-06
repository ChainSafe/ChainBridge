// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"os"
	"testing"

	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const TestEndpoint = "ws://127.0.0.1:9944"

const TestRelayerThreshold = 2
const TestChainId = 1

var AliceKey = keystore.TestKeyRing.SubstrateKeys[keystore.AliceKey].AsKeyringPair()
var BobKey = keystore.TestKeyRing.SubstrateKeys[keystore.BobKey].AsKeyringPair()

var TestLogLevel = log15.LvlTrace
var AliceTestLogger = newTestLogger("Alice")
var BobTestLogger = newTestLogger("Bob")

var ThisChain msg.ChainId = 1
var ForeignChain msg.ChainId = 2

var relayers = []types.AccountID{
	types.NewAccountID(AliceKey.PublicKey),
	types.NewAccountID(BobKey.PublicKey),
}

var resources = map[msg.ResourceId]utils.Method{
	// These are taken from the Polkadot JS UI (Chain State -> Constants)
	msg.ResourceIdFromSlice(hexutil.MustDecode("0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00")): utils.ExampleTransferMethod,
	msg.ResourceIdFromSlice(hexutil.MustDecode("0x000000000000000000000000000000e389d61c11e5fe32ec1735b3cd38c69501")): utils.ExampleMintErc721Method,
	msg.ResourceIdFromSlice(hexutil.MustDecode("0x000000000000000000000000000000f44be64d2de895454c3467021928e55e01")): utils.ExampleRemarkMethod,
}

const relayerThreshold = 2

type testContext struct {
	client         *utils.Client
	listener       *listener
	router         *mockRouter
	writerAlice    *writer
	writerBob      *writer
	latestOutNonce msg.Nonce
	latestInNonce  msg.Nonce
	lSysErr        chan error
	wSysErr        chan error
}

var context testContext

func TestMain(m *testing.M) {
	client, err := utils.CreateClient(AliceKey, TestEndpoint)
	if err != nil {
		panic(err)
	}

	var nativeTokenId, hashId, nftTokenId []byte

	err = utils.QueryConst(client, "Example", "NativeTokenId", &nativeTokenId)
	if err != nil {
		panic(err)
	}

	err = utils.QueryConst(client, "Example", "HashId", &hashId)
	if err != nil {
		panic(err)
	}
	err = utils.QueryConst(client, "Example", "Erc721Id", &nftTokenId)
	if err != nil {
		panic(err)
	}

	err = utils.InitializeChain(client, relayers, []msg.ChainId{ForeignChain}, resources, relayerThreshold)
	if err != nil {
		panic(err)
	}

	aliceConn, bobConn, wSysErr, err := createAliceAndBobConnections()
	if err != nil {
		panic(err)
	}
	l, lSysErr, r, err := newTestListener(client, aliceConn)
	if err != nil {
		panic(err)
	}
	alice := NewWriter(aliceConn, AliceTestLogger, wSysErr, nil, true)
	bob := NewWriter(bobConn, BobTestLogger, wSysErr, nil, true)
	context = testContext{
		client:         client,
		listener:       l,
		router:         r,
		writerAlice:    alice,
		writerBob:      bob,
		latestInNonce:  0,
		latestOutNonce: 0,
		lSysErr:        lSysErr,
		wSysErr:        wSysErr,
	}

	os.Exit(m.Run())
}

func newTestLogger(name string) log15.Logger {
	tLog := log15.Root().New("chain", name)
	tLog.SetHandler(log15.LvlFilterHandler(TestLogLevel, tLog.GetHandler()))
	return tLog
}

// createAliceConnection creates and starts a connection with the Alice keypair
func createAliceConnection() (*Connection, chan error, error) {
	sysErr := make(chan error)
	alice := NewConnection(TestEndpoint, "Alice", AliceKey, AliceTestLogger, make(chan int), sysErr)
	err := alice.Connect()
	if err != nil {
		return nil, nil, err
	}
	return alice, sysErr, err
}

// createAliceAndBobConnections creates and calls `Connect()` on two Connections using the Alice and Bob keypairs
func createAliceAndBobConnections() (*Connection, *Connection, chan error, error) {
	alice, sysErr, err := createAliceConnection()
	if err != nil {
		return nil, nil, nil, err
	}

	bob := NewConnection(TestEndpoint, "Bob", BobKey, AliceTestLogger, make(chan int), sysErr)
	err = bob.Connect()
	if err != nil {
		return nil, nil, nil, err
	}

	return alice, bob, sysErr, nil
}

// getFreeBalance queries the balance for an account, storing the result in `res`
func getFreeBalance(c *Connection, res *types.U128) {
	var acct types.AccountInfo

	ok, err := c.queryStorage("System", "Account", c.key.PublicKey, nil, &acct)
	if err != nil {
		panic(err)
	} else if !ok {
		panic("no account data")
	}
	*res = acct.Data.Free
}
