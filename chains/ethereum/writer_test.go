package ethereum

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/common"
	"github.com/ChainSafe/ChainBridgeV2/crypto/secp256k1"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethparams "github.com/ethereum/go-ethereum/params"
)

var TestPrivateKey = "ae6a8b4518e3970a0501ecf796a51dc0dab9143a66be75e948bf352582db15d5"
var TestAddress = "C5F737aE7EaBB7226f21121E335b0949d8eA6365"
var TestContractAddress = "f563197e48e62ae4ce194e89f7fca67e05cfe5fc"

func testHandler(t *testing.T) msg.Message {
	// arbitrary hash
	data, err := hex.DecodeString("b6e25575ab25a1938070eeb64ac4cd6df49af423327877522bec719815dc5e27")
	if err != nil {
		t.Fatal(err)
	}

	return msg.Message{
		Type: msg.AssetTransferType,
		Data: data,
	}
}

func TestResolveMessage(t *testing.T) {
	m := testHandler(t)

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
		Endpoint: "http://localhost:8545",
		Keypair:  kp,
		Signer:   signer,
		HttpConn: true,
	}

	conn := NewConnection(cfg)
	err = conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	w := NewWriter(conn)
	err = w.ResolveMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteToCentrifugeContract(t *testing.T) {
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
		Endpoint: "http://localhost:8545",
		Keypair:  kp,
		Signer:   signer,
		HttpConn: true,
		Away:     TestContractAddress,
	}

	conn := NewConnection(cfg)
	err = conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	contractBytes, err := hex.DecodeString("F60D9c8AC3B9B88483cee749b25117330F927780")
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

	calldata := common.FunctionId("store(bytes32)")
	// arbitrary hash
	hash, err := hex.DecodeString("58c4b8bfc2aa5dd1da2f02da071c35c1c9f1bc95c4b8b957d44119af6abd5df9")
	if err != nil {
		t.Fatal(err)
	}
	calldata = append(calldata, hash...)

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

	err = conn.SubmitTx(data)
	if err != nil {
		t.Fatal(err)
	}
}
