package tris

import (
	"log/slog"
	"sync"

	"github.com/akyrey/go-tris/internal"
	"github.com/akyrey/go-tris/pkg/transport"
)

type TrisServer struct {
	trisServerOpts
	games    []TrisGame
	gameLock sync.Mutex
}

func NewTrisServer(t transport.Transport, opts ...TrisServerOption) (*TrisServer, error) {
	options := trisServerOpts{transport: t}
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}
	if options.logger == nil {
		options.logger = internal.DefaultLogger()
	}
	return &TrisServer{
		trisServerOpts: options,
		gameLock:       sync.Mutex{},
	}, nil
}

func (s *TrisServer) Start() error {
	s.logger.Debug("starting server on ", slog.String("address", s.transport.Addr()))

	if err := s.transport.ListenAndAccept(); err != nil {
		return err
	}

	s.loop()

	return nil
}

func (s *TrisServer) OnPlayerConnect(node transport.Node) error {
	s.gameLock.Lock()
	defer s.gameLock.Unlock()

	found := false
	// Search for a game with a free slot
	for i, g := range s.games {
		if g.player1 != nil && g.player2 != nil {
			continue
		}
		if g.player1 == nil {
			s.logger.Debug("player 1 connected")
			g.player1 = &node
		} else if g.player2 == nil {
			g.player2 = &node
			s.logger.Debug("player 2 connected")
		}
		found = true
		if g.player1 != nil && g.player2 != nil {
			// TODO: Game is ready to start
			s.logger.Debug("game is ready to start", slog.Int("game_id", i))
		}
	}

	if !found {
		s.logger.Debug("new game created")
		s.games = append(s.games, TrisGame{player1: &node, player2: nil, lock: &sync.Mutex{}, turn: "X"})
	}

	return nil
}

func (s *TrisServer) loop() {
	defer func() {
		s.logger.Debug("Tris server stopped due to error or user quit action")
		s.transport.Close()
	}()
}
