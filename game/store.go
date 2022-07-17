package game

import (
	"github.com/google/uuid"
	"sync"
)

type Store interface {
	Add(game *Game) bool
	IsValidGameId(id uuid.UUID) bool
	AddPlayer(id uuid.UUID) int
	Get(id uuid.UUID) *Game
}

type GameStore struct {
	currentGames map[uuid.UUID]*Game
	sync.RWMutex
}

func NewGameStore() *GameStore {
	return &GameStore{currentGames: map[uuid.UUID]*Game{}}
}

func (s *GameStore) Add(game *Game) bool {
	s.Lock()
	defer s.Unlock()
	s.currentGames[game.Id] = game
	return true
}

func (s *GameStore) IsValidGameId(id uuid.UUID) bool {
	s.Lock()
	defer s.Unlock()
	if s.currentGames[id] == nil {
		return false
	}
	return true
}

func (s *GameStore) AddPlayer(id uuid.UUID) int {
	s.Lock()
	defer s.Unlock()
	game := s.currentGames[id]
	game.NumberOfPlayers++
	return game.NumberOfPlayers
}

func (s *GameStore) Get(id uuid.UUID) *Game {
	s.Lock()
	defer s.Unlock()
	game := s.currentGames[id]
	return game
}
