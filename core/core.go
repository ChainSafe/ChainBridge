package core

import (
	"fmt"

	msg "github.com/ChainSafe/ChainBridgeV2/message"

	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Core struct {
	registry map[msg.ChainId]*Chain
	ks       *keystore.KeyStore
}

func NewCore(ks *keystore.KeyStore) *Core {
	return &Core{
		registry: make(map[msg.ChainId]*Chain),
		ks:       ks,
	}
}

func (c *Core) AddChain(chain *Chain) {
	c.registry[chain.Id()] = chain
}

func (c *Core) Start() {
	for _, ch := range c.registry {
		err := ch.Start()
		if err != nil {
			log15.Error(
				"failed to start chain",
				"chain", ch.Id(),
				"err", err,
			)
		} else {
			log15.Info(fmt.Sprintf("Started %s chain", ch.Id()))
		}
	}
}
