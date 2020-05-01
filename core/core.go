// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	msg "github.com/ChainSafe/ChainBridge/message"
	"github.com/ChainSafe/ChainBridge/router"

	"github.com/ChainSafe/log15"
)

type Core struct {
	ctx      context.Context
	registry map[msg.ChainId]Chain
	route    *router.Router
	log      log15.Logger
}

func NewCore() *Core {
	return &Core{
		ctx:      context.Background(),
		registry: make(map[msg.ChainId]Chain),
		route:    router.NewRouter(log15.New("system", "router")),
		log:      log15.New("system", "core"),
	}
}

func (c *Core) Context() context.Context {
	return c.ctx
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
			c.log.Error(
				"failed to start chain",
				"chain", chain.Id(),
				"err", err,
			)
			return
		}
		c.log.Info(fmt.Sprintf("Started %s chain", chain.Name()))
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)

	// Block here and wait for a signal
	select {
	case <-sigc:
		c.log.Warn("Interrupt received, shutting down now.")
	case <-c.ctx.Done():
		c.log.Warn("Context cancelled, shutting down now.")
	}

	for _, ch := range c.registry {
		err := ch.Stop()
		if err != nil {
			c.log.Error(
				"failed to shutdown chain",
				"chain", ch.Id(),
				"err", err,
			)
		}
	}
}
