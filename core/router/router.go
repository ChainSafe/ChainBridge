package router

import (
	"fmt"
	"sync"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	log "github.com/ChainSafe/log15"
)

// Router forwards messages from their source to their destination
type Router struct {
	registry map[msg.ChainId]chan<- msg.Message
	lock sync.RWMutex
}

func NewRouter() *Router {
	return &Router{
		registry: make(map[msg.ChainId]chan<- msg.Message),
		lock: sync.RWMutex{},
	}
}

// Send passes a message to the destination channel if it exists
func (r *Router) Send(msg msg.Message) error {
	r.lock.RLock()
	defer r.lock.RUnlock()
	log.Trace("Sending message", "src", msg.Source.String(), "dest", msg.Destination.String())
	ch := r.registry[msg.Destination]
	if ch == nil {
		return fmt.Errorf("unknown destination chainId: %d", msg.Destination)
	}
	// TODO: Need to preserve ordering, perhaps a queue would help
	go func (){ch <- msg}()
	return nil
}

// Listen registers a channel with a ChainId which Send can then use to forward messages
func (r *Router) Listen(id msg.ChainId, ch chan<- msg.Message) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	log.Debug("Registering new chain in router", "id", id)
	if r.registry[id] != nil {
		return fmt.Errorf("attempted to register chain that is already registered")
	}

	r.registry[id] = ch
	return nil
}