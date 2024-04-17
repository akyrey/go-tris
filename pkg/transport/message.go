package transport

// Message is a struct that represents a message sent between nodes.
type Message struct {
	From    string
	Payload []byte
}
