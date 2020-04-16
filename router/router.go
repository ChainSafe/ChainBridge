// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package router

import (
	"fmt"
	"sync"

	"github.com/ChainSafe/ChainBridge/chains"
	msg "github.com/ChainSafe/ChainBridge/message"
	log "github.com/ChainSafe/log15"
)

// Router forwards messages from their source to their destination
type Router struct {
	registry map[msg.ChainId]chains.Writer
	lock     *sync.RWMutex
	log      log.Logger
}

func NewRouter(log log.Logger) *Router {
	return &Router{
		registry: make(map[msg.ChainId]chains.Writer),
		lock:     &sync.RWMutex{},
		log:      log,
	}
}

// Send passes a message to the destination Writer if it exists
func (r *Router) Send(msg msg.Message) error {
	r.lock.RLock()
	defer r.lock.RUnlock()

	r.log.Trace("Routing message", "src", msg.Source, "dest", msg.Destination)
	w := r.registry[msg.Destination]
	if w == nil {
		return fmt.Errorf("unknown destination chainId: %d", msg.Destination)
	}
	// TODO: Need to preserve ordering, perhaps a queue would help
	w.ResolveMessage(msg)
	return nil
}

// Listen registers a Writer with a ChainId which Router.Send can then use to propagate messages
func (r *Router) Listen(id msg.ChainId, w chains.Writer) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.log.Debug("Registering new chain in router", "id", id)
	r.registry[id] = w
}
