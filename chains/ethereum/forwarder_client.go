package ethereum

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Keeps track of forwarder info include the state of the current
// nonces
type ForwarderClientBase struct {
	client             *ethclient.Client
	forwarderAddress   common.Address
	forwarderAbi       abi.ABI
	forwarderNonceLock sync.Mutex
	forwarderNonce     *big.Int
	fromAddress        common.Address
	chainId            *big.Int
}

func (c *ForwarderClientBase) ChainId() *big.Int {
	return c.chainId
}

func (c *ForwarderClientBase) ForwarderAddress() common.Address {
	return c.forwarderAddress
}

// Unlocks nonce access and sets the provided nonce
// If transaction usage of a nonce failes the nonce shouldnt be
// updated upon unlock, instead nil should be supplied
func (c *ForwarderClientBase) UnlockAndSetNonce(nonce *big.Int) {
	if nonce != nil {
		c.forwarderNonce = nonce
	}
	c.forwarderNonceLock.Unlock()
}
