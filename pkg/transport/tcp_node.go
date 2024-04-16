package transport

import (
	"net"
	"sync"
)

// TCPNode represents the remote node over a TCP established connection.
type TCPNode struct {
	// The underlying TCP connection of the peer
	net.Conn
	wg *sync.WaitGroup
}

func NewTCPNode(conn net.Conn) *TCPNode {
	return &TCPNode{
		Conn: conn,
		wg:   &sync.WaitGroup{},
	}
}

func (p *TCPNode) CloseStream() {
	p.wg.Done()
}

func (p *TCPNode) Send(b []byte) error {
	_, err := p.Write(b)
	return err
}
