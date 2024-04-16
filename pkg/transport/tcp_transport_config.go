package transport

import "log/slog"

// tcpTransportOpts is a struct that holds the configuration for a TCPTransport.
type tcpTransportOpts struct {
	listenAddr *string
	logger     *slog.Logger
}

// If no listen address is provided, the default listen address is used.
var defaultListenAddr = ":3000"

// TCPTransportOption is a function that updates the TCPTransportConfig struct.
type TCPTransportOption func(*tcpTransportOpts) error

// WithListenAddr is a configuration function that updates the listen address.
func WithListenAddr(addr string) TCPTransportOption {
	return func(c *tcpTransportOpts) error {
		c.listenAddr = &addr
		return nil
	}
}

// WithLogger is a configuration function that updates the logger.
func WithLogger(logger *slog.Logger) TCPTransportOption {
	return func(c *tcpTransportOpts) error {
		c.logger = logger
		return nil
	}
}
