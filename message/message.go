package msg

// Message is used as a generic format to communicate between chains
type Message struct {
	Source      ChainId // Source where message was initiated
	Destination ChainId // Destination chain of message
	// TODO: Add data fields
}
