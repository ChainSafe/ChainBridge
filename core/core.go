// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/ChainBridge/router"

	"github.com/ChainSafe/log15"
)

type Core struct {
	registry map[msg.ChainId]Chain
	route    *router.Router
}

func NewCore() *Core {
	return &Core{
		registry: make(map[msg.ChainId]Chain),
		route:    router.NewRouter(),
	}
}

// AddChain registers the chain in the registry and calls Chain.SetRouter()
func (c *Core) AddChain(chain Chain) {
	c.registry[chain.Id()] = chain
	chain.SetRouter(c.route)
}

// Start will call all registered chains' Start methods and block forever (or until signal is received)
func (c *Core) Start() {
	for _, chain := range c.registry {
		err := chain.Start()
		if err != nil {
			log15.Error(
				"failed to start chain",
				"chain", chain.Id(),
				"err", err,
			)
			return
		}
		log15.Info(fmt.Sprintf("Started %s chain", chain.Name()))
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)

	// Block here and wait for a signal
	<-sigc
	log15.Info("Interrupt received, shutting down now.")
	for _, ch := range c.registry {
		err := ch.Stop()
		if err != nil {
			log15.Error(
				"failed to shutdown chain",
				"chain", ch.Id(),
				"err", err,
			)
		}
	}
}
