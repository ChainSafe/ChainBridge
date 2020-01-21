package ethereum

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/keystore"
	centrifuge "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	msg "github.com/ChainSafe/ChainBridgeV2/message"

	//ethtypes "github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"

)

func testDepositAssetMessage(t *testing.T) msg.Message {
	// arbitrary hash
	data, err := hex.DecodeString("b6e25575ab25a1938070eeb64ac4cd6df49af423327877522bec719815dc5e27")
	if err != nil {
		t.Fatal(err)
	}

	return msg.Message{
		Type: msg.DepositAssetType,
		Data: data,
	}
}

func testCreateDepositProposalMessage(t *testing.T) msg.Message {
	// arbitrary hash
	data, err := hex.DecodeString("b6e25575ab25a1938070eeb64ac4cd6df49af423327877522bec719815dc5e27")
	if err != nil {
		t.Fatal(err)
	}

	depositId := [32]byte{0}
	originChain := [32]byte{1}

	return msg.Message{
		Type: msg.CreateDepositProposalType,
		Data: append(append(data, depositId[:]...), originChain[:]...),
	}
}

func createTestReceiverContract(t *testing.T, conn *Connection) ReceiverContract {
	addressBytes := TestReceiverContractAddress.Bytes()

	address := [20]byte{}
	copy(address[:], addressBytes)

	contract, err := receiver.NewReceiver(address, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	instance := &receiver.ReceiverRaw{
		Contract: contract,
	}

	return instance
}

func createTestCentrifugeContract(t *testing.T, conn *Connection) ReceiverContract {
	addressBytes := TestCentrifugeContractAddress.Bytes()

	address := [20]byte{}
	copy(address[:], addressBytes)

	contract, err := centrifuge.NewBridgeAsset(address, conn.conn)
	if err != nil {
		t.Fatal(err)
	}

	instance := &centrifuge.BridgeAssetRaw{
		Contract: contract,
	}

	return instance
}

func TestResolveMessage(t *testing.T) {
	m := testDepositAssetMessage(t)

	cfg := &Config{
		endpoint: TestEndpoint,
		receiver:  TestCentrifugeContractAddress,
		keystore: keystore.NewTestKeystore(),
		from:     "ethereum",
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	centrifugeContract := createTestCentrifugeContract(t, conn)
	conn.SetReceiverContract(centrifugeContract)

	w := NewWriter(conn, cfg)
	w.ResolveMessage(m)
}

// func TestWriteToCentrifugeContract(t *testing.T) {
// 	cfg := &Config{
// 		endpoint: TestEndpoint,
// 		keystore: keystore.NewTestKeystore(),
// 		from:     "ethereum",
// 	}

// 	conn := NewConnection(cfg)
// 	err := conn.Connect()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	contract := createTestCentrifugeContract(t, conn)
// 	auth := createTestAuth(t, conn)

// 	_, err = contract.Transact(auth, "store", [32]byte{1,2,3,4})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestWriteToReceiverContract(t *testing.T) {
	cfg := &Config{
		endpoint: TestEndpoint,
		keystore: keystore.NewTestKeystore(),
		from:     "ethereum",
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	contract := createTestReceiverContract(t, conn)
	auth := createTestAuth(t, conn)

	depositId := big.NewInt(0)
	originChain := big.NewInt(1)

	_, err = contract.Transact(auth, "createDepositProposal", [32]byte{1,2,3,4}, depositId, originChain)
	if err != nil {
		t.Fatal(err)
	}	

	_, err = contract.Transact(auth, "voteDepositProposal", originChain, depositId, uint8(1))
	if err != nil {
		t.Fatal(err)
	}	

	_, err = contract.Transact(auth, "executeDeposit", originChain, depositId, TestAddress, []byte("nootwashere"))
	if err != nil {
		t.Fatal(err)
	}		
}

func TestWriter_ReceiverContract(t *testing.T) {
	m := testCreateDepositProposalMessage(t)

	cfg := &Config{
		endpoint: TestEndpoint,
		receiver: TestReceiverContractAddress,
		keystore: keystore.NewTestKeystore(),
		from:     "ethereum",
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	receiver := createTestReceiverContract(t, conn)
	conn.SetReceiverContract(receiver)

	w := NewWriter(conn, cfg)
	w.ResolveMessage(m)
}