package game_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"hangman/game"
	"hangman/game/mocks"
	"sync"
)

var _ = Describe("Game Service", func() {
	var mockStore *mocks.Store
	BeforeEach(func() {
		mockStore = &mocks.Store{}
	})

	It("creates a new game", func() {
		mockStore.On("Add", mock.Anything).Return(true)
		service := game.NewGameService(mockStore)
		createdGame := service.Create()
		Expect(createdGame).NotTo(BeNil())
		Expect(createdGame.Id).NotTo(BeNil())
		Expect(createdGame.Status).To(Equal(game.StatusInProgress))
		Expect(createdGame.AttemptsLeft).To(Equal(10))
		Expect(createdGame.Word).NotTo(BeNil())
		Expect(createdGame.NumberOfPlayers).To(Equal(0))
		Expect(createdGame.UsedCharacters).To(BeEmpty())
	})

	It("should process a guess attempt for a game", func() {
		service := game.NewGameService(mockStore)
		uuid, err := uuid.NewUUID()
		mockStore.On("Get", uuid).Return(generateGame([]string{}, 3, 0))
		Expect(err).NotTo(HaveOccurred())
		game, err := service.Guess(uuid, "z")
		Expect(err).ToNot(HaveOccurred())
		Expect(game.AttemptsLeft).To(Equal(2))
	})

	It("should handle a guess attempt with a used letter", func() {
		service := game.NewGameService(mockStore)
		uuid, err := uuid.NewUUID()
		Expect(err).NotTo(HaveOccurred())
		mockStore.On("Get", uuid).Return(generateGame([]string{"z"}, 3, 0))
		game, err := service.Guess(uuid, "z")
		Expect(err).To(HaveOccurred())
		Expect(game).To(BeNil())
	})

	It("should process a game as solved when word has been guessed", func() {
		service := game.NewGameService(mockStore)
		uuid, err := uuid.NewUUID()
		Expect(err).NotTo(HaveOccurred())
		mockStore.On("Get", uuid).Return(generateGame([]string{"c", "a", "z", "y"}, 6, 0))
		game, err := service.Guess(uuid, "t")
		Expect(err).ToNot(HaveOccurred())
		Expect(game.Status).To(Equal(1))
	})

	It("should process a game as lost when all attempts used", func() {
		service := game.NewGameService(mockStore)
		uuid, err := uuid.NewUUID()
		Expect(err).NotTo(HaveOccurred())
		mockStore.On("Get", uuid).Return(generateGame([]string{}, 1, 0))
		game, err := service.Guess(uuid, "t")
		Expect(err).ToNot(HaveOccurred())
		Expect(game.Status).To(Equal(2))
	})

	It("should add a player to a game", func() {
		service := game.NewGameService(mockStore)
		uuid, err := uuid.NewUUID()
		mockStore.On("IsValidGameId", uuid).Return(true)
		mockStore.On("AddPlayer", uuid).Return(1)
		Expect(err).NotTo(HaveOccurred())
		err = service.AddPlayer(uuid)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should handle invalid game id when adding player", func() {
		service := game.NewGameService(mockStore)
		uuid, err := uuid.NewUUID()
		mockStore.On("IsValidGameId", uuid).Return(false)
		mockStore.On("AddPlayer", uuid).Return(1)
		Expect(err).NotTo(HaveOccurred())
		err = service.AddPlayer(uuid)
		Expect(err).To(HaveOccurred())
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
