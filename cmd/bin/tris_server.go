package main

import (
	"github.com/akyrey/go-tris/internal"
	"github.com/akyrey/go-tris/internal/tris"
	"github.com/akyrey/go-tris/pkg/transport"
)

func main() {
	logger := internal.DefaultLogger()
	// TODO: handle getting the address from the command line
	tcpTransport, err := transport.NewTCPTransport(transport.WithLogger(logger))
	if err != nil {
		panic(err)
	}
	server, err := tris.NewTrisServer(tcpTransport, tris.WithLogger(logger))
	if err != nil {
		panic(err)
	}

	tcpTransport.OnNodeConnect = server.OnPlayerConnect

	server.Start()
}
