package ethereum

import (
	"encoding/hex"
	"math/big"
	"testing"

	//"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/keystore"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	centrifuge "github.com/ChainSafe/ChainBridgeV2/contracts/BridgeAsset"
	receiver "github.com/ChainSafe/ChainBridgeV2/contracts/Receiver"
	//emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	msg "github.com/ChainSafe/ChainBridgeV2/message"

	//ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

)

func testMessage(t *testing.T) msg.Message {
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

func TestResolveMessage(t *testing.T) {
	m := testMessage(t)

	cfg := &Config{
		endpoint: TestEndpoint,
		emitter:  TestCentrifugeContractAddress,
		keystore: keystore.NewTestKeystore(),
		from:     "ethereum",
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	w := NewWriter(conn, cfg)
	w.ResolveMessage(m)

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

func createTestAuth(t *testing.T, conn *Connection) *bind.TransactOpts {
	currBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	nonce, err := conn.NonceAt(TestAddress, currBlock.Number())
	if err != nil {
		t.Fatal(err)
	}

	privateKey := conn.kp.Private().(*secp256k1.PrivateKey).Key()
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = big.NewInt(10)

	return auth
}

func TestWriteToCentrifugeContract(t *testing.T) {
	cfg := &Config{
		endpoint: TestEndpoint,
		receiver: TestCentrifugeContractAddress,
		keystore: keystore.NewTestKeystore(),
		from:     "ethereum",
	}

	conn := NewConnection(cfg)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	contract := createTestCentrifugeContract(t, conn)
	auth := createTestAuth(t, conn)

	_, err = contract.Transact(auth, "store", [32]byte{1,2,3,4})
	if err != nil {
		t.Fatal(err)
	}
}
