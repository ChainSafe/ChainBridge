package ethtest

import (
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetDepositNonce(t *testing.T, client *ethclient.Client, bridge common.Address, chain msg.ChainId) uint64 {
	count, err := utils.GetDepositNonce(client, bridge, chain)
	if err != nil {
		t.Fatal(err)
	}
	return count
}
