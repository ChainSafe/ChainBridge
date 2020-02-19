package chains

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type Connection interface {
	Connect() error
	SubmitTx([]byte) error
	Close()
}

type Router interface {
	Send(message msg.Message) error
}

type Listener interface {
	Start() error
	SetRouter(Router)
	RegisterEventHandler(string, EvtHandlerFn) error
	Stop() error
}

type EvtHandlerFn func(interface{}) msg.Message

type Writer interface {
	Start() error
	ResolveMessage(message msg.Message)
	Stop() error
}
