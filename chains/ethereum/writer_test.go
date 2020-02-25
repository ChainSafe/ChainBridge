package ethereum

import (
	"encoding/hex"
	"math/big"
	"testing"

	centrifuge "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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
	depositId := big.NewInt(0)
	originChain := msg.ChainId(1)

	m := &msg.Message{
		Type: msg.CreateDepositProposalType,
		Data: []byte{},
	}

	m.EncodeCreateDepositProposalData(depositId, originChain)
	return *m
}

func testVoteDepositProposalMessage(t *testing.T) msg.Message {
	depositId := big.NewInt(0)
	originChain := big.NewInt(1)
	vote := uint8(1)

	m := &msg.Message{
		Type: msg.VoteDepositProposalType,
	}

	m.EncodeVoteDepositProposalData(depositId, originChain, vote)
	return *m
}

func testExecuteDepositMessage(t *testing.T) msg.Message {
	depositId := big.NewInt(0)
	originChain := big.NewInt(1)
	address := TestAddress

	m := &msg.Message{
		Type: msg.ExecuteDepositType,
	}

	m.EncodeExecuteDepositData(depositId, originChain, address, []byte("nootwashere"))
	return *m
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
		receiver: TestCentrifugeContractAddress,
		keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	centrifugeContract := createTestCentrifugeContract(t, conn)
	w := NewWriter(conn, cfg)
	w.SetReceiverContract(centrifugeContract)
	w.ResolveMessage(m)
}

func TestWriteToReceiverContract(t *testing.T) {
	cfg := &Config{
		endpoint: TestEndpoint,
		keystore: keystore.TestKeyStoreMap[keystore.AliceKey],
		from:     keystore.AliceKey,
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	contract := createTestReceiverContract(t, conn)
	auth := createTestAuth(t, conn)

	depositId := big.NewInt(421)
	originChain := big.NewInt(1)

	data := []byte("nootwashere")
	hash := ethcrypto.Keccak256Hash(data)

	_, err = contract.Transact(auth, "createDepositProposal", hash, depositId, originChain)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriter_createDepositProposal(t *testing.T) {
	m := testCreateDepositProposalMessage(t)

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
	defer conn.Close()

	receiver := createTestReceiverContract(t, conn)
	w := NewWriter(conn, cfg)
	w.SetReceiverContract(receiver)
	w.ResolveMessage(m)
}

func TestWriter_voteDepositProposal(t *testing.T) {
	m := testVoteDepositProposalMessage(t)

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
	defer conn.Close()

	receiver := createTestReceiverContract(t, conn)
	w := NewWriter(conn, cfg)
	w.SetReceiverContract(receiver)
	w.ResolveMessage(m)
}

func TestWriter_executeDeposit(t *testing.T) {
	m := testExecuteDepositMessage(t)

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
	defer conn.Close()

	receiver := createTestReceiverContract(t, conn)
	w := NewWriter(conn, cfg)
	w.SetReceiverContract(receiver)
	w.ResolveMessage(m)
}
