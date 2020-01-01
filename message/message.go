package msg

type MessageType string

var DepositAssetType MessageType = "centrifuge_deposit_asset"

// Message is used as a generic format to communicate between chains
type Message struct {
	Source      ChainId // Source where message was initiated
	Destination ChainId // Destination chain of message
	// TODO: Add data fields
	Type MessageType // type of bridge transfer
	Data []byte      // data associated with event sequence
}
