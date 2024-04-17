package tris

import (
	"sync"

	"github.com/akyrey/go-tris/pkg/transport"
)

// Player 1 is always X, player 2 is always O.
type TrisGame struct {
	player1 *transport.Node
	player2 *transport.Node
	lock    *sync.Mutex
	turn    string
}
