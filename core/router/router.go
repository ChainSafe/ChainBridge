package router

import (
	"fmt"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	log "github.com/ChainSafe/log15"
)

type Router struct {
	registry map[msg.ChainId]chan<- msg.Message
}

func NewRouter() *Router {
	return &Router{registry: make(map[msg.ChainId]chan<- msg.Message)}
}

// Send passes a message to the destination channel if it exists
func (r *Router) Send(msg msg.Message) error {
	ch := r.registry[msg.Destination]
	if ch == nil {
		return fmt.Errorf("unknown destination chainId: %d", msg.Destination)
	}
	// TODO: This will block, we want to avoid this but preserve ordering, perhaps a queue would help
	ch <- msg
	return nil
}

// Listen registers a channel with a ChainId which Send can then use to forward messages
func (r *Router) Listen(id msg.ChainId, ch chan<- msg.Message) error {
	log.Debug("Registering new chain in router", "id", id)
	if r.registry[id] != nil {
		return fmt.Errorf("attempted to register chain that is already registered")
	}

	r.registry[id] = ch
	return nil
}