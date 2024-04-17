package tris

import (
	"log/slog"

	"github.com/akyrey/go-tris/pkg/transport"
)

type trisServerOpts struct {
	logger    *slog.Logger
	transport transport.Transport
}

// TrisServerOption is a function that updates the trisServerOpts struct.
type TrisServerOption func(*trisServerOpts) error

// WithLogger is a configuration function that updates the logger.
func WithLogger(logger *slog.Logger) TrisServerOption {
	return func(c *trisServerOpts) error {
		c.logger = logger
		return nil
	}
}
