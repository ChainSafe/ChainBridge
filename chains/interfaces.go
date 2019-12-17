package chains

import (
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/ChainBridgeV2/types"
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
	ResolveMessage(message msg.Message)
}

type Event interface {
	GetParams() []byte
}
