package msg

// Message is used as a generic format to communicate between chains
type Message struct {
	Source      ChainId
	Destination ChainId
	// TODO: Add data fields
}
