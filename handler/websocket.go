package handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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
	go ws.Listen()
}

func (h *Handler) Broadcast() {
	for {
		for _, client := range h.connections.clients {
			game, err := h.service.Get(client.gameId)
			if err != nil {
				client.conn.Close()
			}
			client.conn.WriteJSON(game)
		}
		time.Sleep(200 * time.Millisecond) // artificial delay so that data can be read clearly
	}
}
