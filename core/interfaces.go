package core

import (
	msg "ChainBridgeV2/message"
)

type Connection interface {
	Connect() error
	SubmitTx([]byte) error
}

type Listener interface {
	RegisterEventHandler()
}

type Writer interface {
	ResolveMessage(message msg.Message) error
}

type Event interface {
	GetParams() []byte
}
