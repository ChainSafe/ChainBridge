package ethereum

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

func newLocalConnection(t *testing.T, emitter string) *Connection {
	ctx := context.Background()

	signer := ethtypes.MakeSigner(ethparams.MainnetChainConfig, big.NewInt(0))
	privBytes, err := hex.DecodeString(TestPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	priv, err := secp256k1.NewPrivateKey(privBytes)
	if err != nil {
		t.Fatal(err)
	}

	kp, err := secp256k1.NewKeypairFromPrivate(priv)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: TestEndpoint,
		Keypair:  kp,
		Signer:   signer,
		Emitter:  emitter,
	}

	conn := NewConnection(cfg)
	err = conn.Connect()
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
	conn := newLocalConnection(t, TestCentrifugeContractAddress)
	defer conn.Close()

	// send tx to trigger event in EmitterContract
	contractBytes, err := hex.DecodeString(TestEmitterContractAddress)
	if err != nil {
		t.Fatal(err)
	}

	contract := [20]byte{}
	copy(contract[:], contractBytes)

	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	nonce, err := conn.NonceAt(common.StringToAddress(TestAddress), currBlock.Number())
	if err != nil {
		t.Fatal(err)
	}

	event := EventSig("Transfer(address,bytes32)")
	cfg := &core.ChainConfig{
		Receiver:      TestEmitterContractAddress,
		Emitter:       TestCentrifugeContractAddress,
		Subscriptions: []string{string(event)},
	}
	listener := NewListener(conn, cfg)

	// calling fallback
	calldata := []byte{}

	tx := ethtypes.NewTransaction(
		nonce,
		contract,
		big.NewInt(0),
		1000000,        // gasLimit
		big.NewInt(10), // gasPrice
		calldata,
	)

	data, err := tx.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	// subscribe to event
	query := listener.buildQuery(common.StringToAddress(TestEmitterContractAddress), event)
	subscription, err := conn.subscribeToEvent(query)
	if err != nil {
		t.Fatal(err)
	}
	defer subscription.sub.Unsubscribe()

	// send tx to trigger event
	err = conn.SubmitTx(data)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case evt := <-subscription.ch:
		t.Log("got event", evt)
	case <-time.After(TestTimeout):
		t.Fatal("Timed out")
	}

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
