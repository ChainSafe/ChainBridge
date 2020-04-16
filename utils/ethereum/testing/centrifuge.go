package ethtest

import (
	"testing"

	utils "github.com/ChainSafe/ChainBridge/utils/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AssertHashExistence(t *testing.T, client *ethclient.Client, hash [32]byte, contract common.Address) {
	exists, err := utils.HashExists(client, hash, contract)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("Hash %x does not exist on chain", hash)
	}
}
