package core

import (
	msg "ChainBridgeV2/message"
	"fmt"

	"github.com/ChainSafe/log15"
)

type Core struct {
	registry map[msg.ChainId]*Chain
}

func NewCore() *Core {
	return &Core{registry: make(map[msg.ChainId]*Chain)}
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
