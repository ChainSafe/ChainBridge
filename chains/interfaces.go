package chains

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
)

type Connection interface {
	Connect() error
	SubmitTx([]byte) error
	Close()
}

type Listener interface {
	Start() error
	RegisterEventHandler(string, HandlerFn) error
	Stop() error
}

type HandlerFn func(interface{}) msg.Message

type Writer interface {
	Start() error
	ResolveMessage(message msg.Message)
	Stop() error
}
