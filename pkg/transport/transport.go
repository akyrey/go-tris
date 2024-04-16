package transport

// Transport is anything that handles network communication between nodes.
type Transport interface {
	// Addr returns the address of the transport.
	Addr() string
	// Close closes the transport.
	Close() error
	// ListenAndAccept listens for incoming connections and accepts them.
	ListenAndAccept() error
}
