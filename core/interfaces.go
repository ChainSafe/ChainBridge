package core

import (
	msg "ChainBridgeV2/message"
	"ChainBridgeV2/types"
)

type Connection interface {
	Connect() error
	SubmitTx([]byte) error
	Close()
}

type Listener interface {
	RegisterEventHandler(types.EventName, func())
}

type Writer interface {
	ResolveMessage(message msg.Message) error
}

type Event interface {
	GetParams() []byte
}
