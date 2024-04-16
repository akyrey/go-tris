package transport

import "net"

// Node is an interface representing a node in the network.
type Node interface {
	net.Conn
	Send([]byte) error
}
