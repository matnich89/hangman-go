package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"hangman/game"
	"hangman/game/mocks"
	"hangman/handler"
	"net/http"
	"net/http/httptest"
	"sync"
)

var _ = Describe("Handler", func() {

	var requestHandler *handler.Handler
	var mockService mocks.Service
	BeforeEach(func() {
		mockService = mocks.Service{}
		requestHandler = handler.NewHandler(&mockService, handler.NewClientConnections())
		Expect(requestHandler).ToNot(BeNil())
	})

	It("should create a new game", func() {
		gameToCreate := generateGame([]string{}, 10, 0)
		mockService.On("Create", mock.Anything).Return(gameToCreate)
		req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
		w := httptest.NewRecorder()
		requestHandler.CreateGame(w, req)
		Expect(w.Code).To(Equal(http.StatusCreated))
		var gameBody game.Game
		err := json.Unmarshal(w.Body.Bytes(), &gameBody)
		Expect(err).NotTo(HaveOccurred())
		Expect(&gameBody).ToNot(BeNil())
		Expect(gameBody.Id).To(Equal(gameToCreate.Id))
		Expect(gameBody.Status).To(Equal(gameToCreate.Status))
		Expect(gameBody.UsedCharacters).To(Equal(gameToCreate.UsedCharacters))
		Expect(gameBody.NumberOfPlayers).To(Equal(gameToCreate.NumberOfPlayers))
		Expect(gameBody.AttemptsLeft).To(Equal(gameToCreate.AttemptsLeft))
	})

	It("should handle game guess", func() {
		guess := &game.GameGuess{
			Id:     uuid.UUID{},
			Letter: "z",
		}
		mockService.On("Guess", guess.Id, guess.Letter).Return(generateGame([]string{}, 9, 0), nil)
		b, err := json.Marshal(&guess)
		Expect(err).NotTo(HaveOccurred())
		req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/", bytes.NewReader(b))
		w := httptest.NewRecorder()
		requestHandler.MakeGuessForGame(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
		var gameBody game.Game
		err = json.Unmarshal(w.Body.Bytes(), &gameBody)
		Expect(err).NotTo(HaveOccurred())
		Expect(&gameBody).ToNot(BeNil())
	})

	It("should handle bad request for game guess", func() {
		req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
		w := httptest.NewRecorder()
		requestHandler.MakeGuessForGame(w, req)
		Expect(w.Code).To(Equal(http.StatusBadRequest))
	})

})

func generateGame(usedChars []string, attemptsLeft, status int) *game.Game {
	return &game.Game{
		Id:              uuid.UUID{},
		Word:            "cat",
		UsedCharacters:  usedChars,
		AttemptsLeft:    attemptsLeft,
		NumberOfPlayers: 1,
		Status:          status,
		RWMutex:         sync.RWMutex{},
	}
}
