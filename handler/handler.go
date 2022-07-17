package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"hangman/game"
	"net/http"
)

type Handler struct {
	service     game.Service
	connections *ClientConnections
}

func NewHandler(service game.Service, connections *ClientConnections) *Handler {
	return &Handler{service: service, connections: connections}
}

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	createdGame := h.service.Create()
	w.WriteHeader(http.StatusCreated)
	bytes, err := json.Marshal(createdGame)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(bytes)
}

func (h *Handler) ConnectGame(w http.ResponseWriter, r *http.Request) {
	conn, err := UpgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Could not upgrade connection"))
	}
	params := mux.Vars(r)
	gameIdParam := params["id"]
	gameId, err := uuid.Parse(gameIdParam)
	if err != nil {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Id must be a UUID"))
		_ = conn.Close()
		return
	}
	h.AddClient(conn, gameId)
	h.service.AddPlayer(gameId)
}

func (h *Handler) MakeGuessForGame(w http.ResponseWriter, r *http.Request) {
	var guess game.GameGuess

	err := json.NewDecoder(r.Body).Decode(&guess)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.Guess(guess.Id, guess.Letter)

	w.WriteHeader(http.StatusOK)

}
