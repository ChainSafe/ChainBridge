package router

import (
	"fmt"
	"sync"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	log "github.com/ChainSafe/log15"
)

// Router forwards messages from their source to their destination
type Router struct {
	registry map[msg.ChainId]chains.Writer
	lock     *sync.RWMutex
}

func NewRouter() *Router {
	return &Router{
		registry: make(map[msg.ChainId]chains.Writer),
		lock:     &sync.RWMutex{},
	}
}

// Send passes a message to the destination Writer if it exists
func (r *Router) Send(msg msg.Message) error {
	r.lock.RLock()
	defer r.lock.RUnlock()

	log.Trace("[Router.go] Sending message", "src", msg.Source, "dest", msg.Destination)
	w := r.registry[msg.Destination]
	if w == nil {
		return fmt.Errorf("unknown destination chainId: %d", msg.Destination)
	}
	// TODO: Need to preserve ordering, perhaps a queue would help
	w.ResolveMessage(msg)
	return nil
}

// Listen registers a Writer with a ChainId which Router.Send can then use to propagate messages
func (r *Router) Listen(id msg.ChainId, w chains.Writer) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	log.Debug("Registering new chain in router", "id", id)
	if r.registry[id] != nil {
		return fmt.Errorf("attempted to register chain that is already registered")
	}

	r.registry[id] = w
	return nil
}
