package core

import (
	msg "ChainBridgeV2/message"
)

//type Chain interface {
//	Start() error
//	Id() ChainId
//}

type Connection interface {
	Connect() error
	SubmitTx([]byte) error
}

type Listener interface {
	RegisterEventHandler()
}

type Pusher interface {
	ResolveMessage(message msg.Message) error
}

type Event interface {
	GetParams() []byte
}
