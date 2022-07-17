package game

import (
	"github.com/google/uuid"
	"sync"
)

const (
	StatusInProgress   int = iota
	StatusFinishedWin      = iota
	StatusFinishedLose     = iota
)

type Game struct {
	Id              uuid.UUID `json:"id"`
	Word            string    `json:"word"`
	UsedCharacters  []string  `json:"usedCharacters"`
	AttemptsLeft    int       `json:"attemptsLeft"`
	NumberOfPlayers int       `json:"numberOfPlayers"`
	Status          int       `json:"status"`
	sync.RWMutex
}

type GameGuess struct {
	Id     uuid.UUID `json:"id"`
	Letter string    `json:"letter"`
}
