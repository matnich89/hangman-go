package handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var UpgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // prevents CORS errors... would not use in production system
	},
}

func (h *Handler) AddClient(client *websocket.Conn, gameId uuid.UUID) {
	ws := &WebsocketConnection{
		conn:   client,
		gameId: gameId,
	}
	h.connections.Append(ws)
	go h.connections.Listen(ws)
}

func (h *Handler) RemoveClient(client *WebsocketConnection) {
	h.connections.Remove(client)
}

func (h *Handler) Broadcast() {
	for {
		for _, client := range h.connections.clients {
			game, err := h.service.Get(client.gameId)
			if err != nil {
				log.Println("error occured when broadcasting ")
				client.conn.Close()
				h.connections.Remove(client)
			}
			client.conn.WriteJSON(game)
		}
		time.Sleep(200 * time.Millisecond) // artificial delay so that data can be read clearly
	}
}
