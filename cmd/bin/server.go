package main

import (
	"github.com/akyrey/go-tris/pkg/transport"
)

func main() {
	// TODO: handle getting the address from the command line
	tcpTransport, err := transport.NewTCPTransport()
	if err != nil {
		panic(err)
	}
	if err := tcpTransport.ListenAndAccept(); err != nil {
		panic(err)
	}
}
