// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/keystore"

	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/SimpleEmitter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
)

func newLocalConnection(t *testing.T, emitter ethcmn.Address) *Connection {

	cfg := &Config{
		endpoint: TestEndpoint,
		receiver: TestCentrifugeContractAddress,
		keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

// TODO: See TestListenerAndWriter
// test handler function for events
//func testTransferHandler(logi interface{}) msg.Message {
//	log := logi.(ethtypes.Log)
//	hash := [32]byte(log.Topics[1])
//	return msg.Message{
//		Destination: msg.EthereumId,
//		Type:        msg.AssetTransferType,
//		Data:        hash[:],
//	}
//}

func TestListener(t *testing.T) {
	// TODO: this works locally, but not on CI for some reason. need to debug.
	t.Skip()

	conn := newLocalConnection(t, TestCentrifugeContractAddress)
	defer conn.Close()

	// send tx to trigger event in EmitterContract
	addressBytes := TestEmitterContractAddress.Bytes()

	address := [20]byte{}
	copy(address[:], addressBytes)

	contract, err := emitter.NewSimpleEmitter(address, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	instance := &emitter.SimpleEmitterRaw{
		Contract: contract,
	}

	auth := createTestAuth(t, conn)

	eventIterator, err := instance.Contract.SimpleEmitterFilterer.FilterDepositAsset(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: context.Background(),
	}, []ethcmn.Address{TestAddress})
	if err != nil {
		t.Fatal(err)
	}

	// call fallback to trigger event
	_, err = instance.Transact(auth, "")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	eventIterator.Next()
	if eventIterator.Event == nil {
		t.Fatal("Did not get event")
	} else {
		t.Log(eventIterator.Event)
		return
	}

	<-time.After(TestTimeout)
	t.Fatal("timeout")
}

func TestListenerAndWriter(t *testing.T) {
	// TODO: Unclear what this is supposed to test
	t.Skip()
	//conn := newLocalConnection(t, TestEmitterContractAddress)
	//defer conn.Close()
	//
	//// setup writer and router
	//writer := NewWriter(conn)
	//r := router.NewRouter()
	//err := r.Listen(msg.EthereumId, writer)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// setup Emitter contract
	//contract := common.StringToAddress(TestEmitterContractAddress)
	//
	//currBlock, err := conn.LatestBlock()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//nonce, err := conn.NonceAt(common.StringToAddress(TestAddress), currBlock.Number())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//event := EventSig("Transfer(address,bytes32)")
	//cfg := &core.ChainConfig{
	//	Receiver:      TestEmitterContractAddress,
	//	Emitter:       TestEmitterContractAddress,
	//	Subscriptions: []string{string(event)},
	//}
	//listener := NewListener(conn, cfg)
	//listener.SetRouter(r)
	//err = listener.RegisterEventHandler(string(event), testTransferHandler)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer listener.Unsubscribe(event)
	//
	//// calling fallback in Emitter to trigger Transfer event
	//tx := ethtypes.NewTransaction(
	//	nonce,
	//	contract,
	//	big.NewInt(0),
	//	1000000,        // gasLimit
	//	big.NewInt(10), // gasPrice
	//	[]byte{},
	//)
	//
	//data, err := tx.MarshalJSON()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// subscribe to event in Centrifuge receiver contract
	//query := listener.buildQuery(common.StringToAddress(TestEmitterContractAddress), EventSig("AssetStored(bytes32)"))
	//subscription, err := conn.subscribeToEvent(query)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// send tx to trigger event
	//err = conn.SubmitTx(data)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//select {
	//case evt := <-subscription.ch:
	//	t.Log("got event", evt)
	//case <-time.After(TestTimeout):
	//	t.Fatal("Timed out")
	//}
}
