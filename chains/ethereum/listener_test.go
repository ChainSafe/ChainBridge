package ethereum

import (
	"context"
	"encoding/hex"
	//"fmt"
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	//msg "github.com/ChainSafe/ChainBridgeV2/message"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

var TestEmitterContractAddress = "1a5eA9856dE8966DF6f3aE55B5cC3CD5207537c8"

func newLocalConnection(t *testing.T, away string) *Connection {
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

	// TODO: add run ./scripts/local_test/start_ganache.sh and ./scripts/local_test/ethereum_start.sh
	// to CI so that contract gets deployed
	cfg := &ConnectionConfig{
		Ctx:      ctx,
		Endpoint: "ws://localhost:8545",
		Keypair:  kp,
		Signer:   signer,
		HttpConn: true,
		Away:     away,
	}

	conn := NewConnection(cfg)
	err = conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

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
		Home:          TestEmitterContractAddress,
		Away:          TestCentrifugeContractAddress,
		Subscriptions: []string{string(event)},
	}
	listener := NewListener(conn, cfg)
	// err = listener.RegisterEventHandler(string(event), testTransferHandler)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// calling fallback
	calldata := []byte{} ///common.FunctionId("arbitrary")

	tx := ethtypes.NewTransaction(
		nonce,
		ethcommon.Address(contract),
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

	// send tx to trigger event
	err = conn.SubmitTx(data)
	if err != nil {
		t.Fatal(err)
	}

	evt := <-subscription.ch
	t.Log("got event", evt)
}

// func TestListenerAndWriter(t *testing.T) {
// 	conn := newLocalConnection(t, TestCentrifugeContractAddress)
// 	defer conn.Close()

// 	// send tx to trigger event in EmitterContract
// 	contractBytes, err := hex.DecodeString(TestEmitterContractAddress)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	contract := [20]byte{}
// 	copy(contract[:], contractBytes)

// 	currBlock, err := conn.LatestBlock()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	nonce, err := conn.NonceAt(common.StringToAddress(TestAddress), currBlock.Number())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	event := EventSig("Transfer(address,bytes32)")
// 	cfg := &core.ChainConfig{
// 		Home:	TestEmitterContractAddress,
// 		Away:    TestCentrifugeContractAddress,
// 		Subscriptions: []string{string(event)},
// 	}
// 	listener := NewListener(conn, cfg)
// 	err = listener.RegisterEventHandler(string(event), testTransferHandler)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// calling fallback
// 	calldata := []byte{}///common.FunctionId("arbitrary")

// 	tx := ethtypes.NewTransaction(
// 		nonce,
// 		ethcommon.Address(contract),
// 		big.NewInt(0),
// 		1000000,        // gasLimit
// 		big.NewInt(10), // gasPrice
// 		calldata,
// 	)

// 	data, err := tx.MarshalJSON()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// subscribe to event in Centrifuge receiver contract
// 	query := listener.buildQuery(common.StringToAddress(TestCentrifugeContractAddress), EventSig("store(bytes32)"))
// 	subscription, err := conn.subscribeToEvent(query)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// send tx to trigger event
// 	err = conn.SubmitTx(data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	evt := <-subscription.ch
// 	t.Log("got event", evt)
// }

// func testTransferHandler(logi interface{}) msg.Message {
// 	log := logi.(ethtypes.Log)
// 	fmt.Printf("got log %v\n", log)
// 	return msg.Message{}
// }
