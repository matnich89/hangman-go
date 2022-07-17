package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"hangman/game"
	"hangman/handler"
	"net/http"
	"net/http/httptest"
	"sync"
)

type mockService struct {
}

func (s *mockService) AddPlayer(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *mockService) Get(id uuid.UUID) *game.Game {
	return &game.Game{
		Id:              uuid.UUID{},
		Word:            "cat",
		UsedCharacters:  []string{},
		AttemptsLeft:    0,
		NumberOfPlayers: 1,
		Status:          0,
		RWMutex:         sync.RWMutex{},
	}
}

func (s *mockService) Connect(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *mockService) Guess(id uuid.UUID, letter string) {
	return
}

func (s *mockService) Create() *game.Game {
	return &game.Game{}
}

var _ = Describe("Handler", func() {

	var requestHandler *handler.Handler

	BeforeEach(func() {
		requestHandler = handler.NewHandler(&mockService{}, handler.NewClientConnections())
		Expect(requestHandler).ToNot(BeNil())
	})

	It("should create a new game", func() {
		req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
		w := httptest.NewRecorder()
		requestHandler.CreateGame(w, req)
		Expect(w.Code).To(Equal(http.StatusCreated))
		var gameBody game.Game
		err := json.Unmarshal(w.Body.Bytes(), &gameBody)
		Expect(err).NotTo(HaveOccurred())
		Expect(gameBody).ToNot(BeNil())
	})

	It("should handle game guess", func() {
		guess := &game.GameGuess{
			Id:     uuid.UUID{},
			Letter: "z",
		}
		b, err := json.Marshal(&guess)
		Expect(err).NotTo(HaveOccurred())
		req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/", bytes.NewReader(b))
		w := httptest.NewRecorder()
		requestHandler.MakeGuessForGame(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	It("should handle bad request for game guess", func() {
		req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
		w := httptest.NewRecorder()
		requestHandler.MakeGuessForGame(w, req)
		Expect(w.Code).To(Equal(http.StatusBadRequest))
	})

})
