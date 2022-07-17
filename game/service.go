package game

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"hangman/util"
	"log"
	"strings"
)

type Service interface {
	Create() *Game
	AddPlayer(id uuid.UUID) error
	Guess(id uuid.UUID, letter string)
	Get(id uuid.UUID) *Game
}

type GameService struct {
	store Store
}

func NewGameService(store Store) *GameService {
	return &GameService{store: store}
}

func (s *GameService) Create() *Game {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Println("Could not create uuid")
	}
	game := Game{
		Id:              id,
		Word:            util.GenerateWord(),
		UsedCharacters:  []string{},
		AttemptsLeft:    10,
		NumberOfPlayers: 0,
		Status:          StatusInProgress,
	}
	s.store.Add(&game)
	return &game
}

func (s *GameService) AddPlayer(id uuid.UUID) error {
	if s.store.IsValidGameId(id) == false {
		return errors.New(fmt.Sprintf("game with id %s does not exist", id.String()))
	}
	s.store.AddPlayer(id)
	return nil
}

func (s *GameService) Guess(id uuid.UUID, letter string) *Game {
	game := s.store.Get(id)
	game.Lock()
	defer game.Unlock()
	game.AttemptsLeft--
	game.UsedCharacters = append(game.UsedCharacters, letter)
	if isSolved(game.Word, game.UsedCharacters) {
		game.Status = StatusFinishedWin
		return game
	}
	if game.AttemptsLeft <= 0 {
		game.Status = StatusFinishedLose
		return game
	}
	return game
}

func (s *GameService) Get(id uuid.UUID) *Game {
	return s.store.Get(id)
}

func isSolved(word string, usedCharacters []string) bool {
	wordArray := strings.Split(word, "")
	solved := true
	for _, letter := range wordArray {
		letterFound := false
		for _, usedLetter := range usedCharacters {
			if usedLetter == letter {
				letterFound = true
				break
			}
		}
		if !letterFound {
			solved = false
			break
		}
	}
	return solved
}
