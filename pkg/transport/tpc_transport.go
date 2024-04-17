package transport

import (
	"errors"
	"log/slog"
	"net"

	"github.com/akyrey/go-tris/internal"
	"github.com/akyrey/go-tris/pkg/assert"
)

type TCPTransport struct {
	listener net.Listener
	tcpTransportOpts
}

// Create a new TCPTransport with the given configuration.
func NewTCPTransport(opts ...TCPTransportOption) (*TCPTransport, error) {
	var options tcpTransportOpts
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}
	if options.listenAddr == nil {
		options.listenAddr = &defaultListenAddr
	}
	if options.logger == nil {
		options.logger = internal.DefaultLogger()
	}
	return &TCPTransport{tcpTransportOpts: options}, nil
}

// Implement the Transport interface by returning the address of the TCPTransport.
func (t *TCPTransport) Addr() string {
	assert.Assert(t.listenAddr != nil, "TCPTransport.Addr: listenAddr cannot be nil")
	return *t.listenAddr
}

// Implement the Transport interface closing the listener.
func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

// Implement the Transport interface listening and handling incoming connections.
func (t *TCPTransport) ListenAndAccept() error {
	assert.Assert(t.listenAddr != nil, "TCPTransport.ListenAndAccept: listenAddr cannot be nil")
	var err error

	t.listener, err = net.Listen("tcp", *t.listenAddr)
	if err != nil {
		return err
	}

	t.logger.Debug("listening on ", slog.String("address", *t.listenAddr))

	for {
		conn, err := t.listener.Accept()
		// If the connection is closed, we stop looping
		if errors.Is(err, net.ErrClosed) {
			t.logger.Debug("TCP listener closed")
			return err
		}

		if err != nil {
			t.logger.Error("TCP accept error: ", slog.Any("error", err))
			continue
		}

		t.logger.Debug("accepted connection from ", slog.String("address", conn.RemoteAddr().String()))
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	defer func() {
		t.logger.Debug("dropping peer connection", slog.Any("error", err))
		conn.Close()
	}()

	node := NewTCPNode(conn)

	if t.OnNodeConnect != nil {
		if err := t.OnNodeConnect(node); err != nil {
			t.logger.Error("failed to perform on node connection action", slog.Any("error", err))
			return
		}
	}

	t.logger.Debug("handling connection from ", slog.String("address", conn.RemoteAddr().String()))
	for {
		buf := make([]byte, 8)
		if _, err = conn.Read(buf); err != nil {
			t.logger.Error("error reading from connection: ", slog.Any("error", err))
			return
		}
		t.logger.Debug("received message", slog.String("message", string(buf)))
		// node.wg.Add(1)
		// node.wg.Wait()
	}
}
