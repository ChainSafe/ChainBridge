package ethereum

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "golang.org/x/crypto/sha3"
	// "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func TestItx(t *testing.T) {
	AliceKp := keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
	to := common.HexToAddress("033eF6db9FBd0ee60E2931906b987Fe0280471a0")
	data := common.Hex2Bytes("212121")
	gas := big.NewInt(35000)
	chainId := big.NewInt(4)
	schedule := "fast"

	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	bytesTy, _ := abi.NewType("bytes", "bytes", nil)
	stringTy, _ := abi.NewType("string", "string", nil)

	// point in time recoveery for the dynamo db
	arguments := abi.Arguments{
		{
			Type: addressTy,
		},
		{
			Type: bytesTy,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: stringTy,
		},
	}

	packed, _ := arguments.Pack(to, data, gas, chainId, schedule)

	relayTxId := crypto.Keccak256Hash(packed)

	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(relayTxId), string(relayTxId.Bytes()))
	msgHash := crypto.Keccak256Hash([]byte(msg))

	// crypto.Sign()
	sig, _ := crypto.Sign(msgHash.Bytes(), AliceKp.PrivateKey())

	fmt.Println(sig)

}
