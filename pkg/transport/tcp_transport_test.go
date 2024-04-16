package transport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tr, err := NewTCPTransport()
	assert.Nil(t, err)
	assert.Equal(t, *tr.listenAddr, ":3000")
	assert.NotNil(t, *tr.logger)

	assert.Nil(t, tr.ListenAndAccept())
}
