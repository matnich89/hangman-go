package handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

type ClientConnections struct {
	sync.RWMutex
	clients []*WebsocketConnection
	wg      sync.WaitGroup
}

func NewClientConnections() *ClientConnections {
	return &ClientConnections{
		RWMutex: sync.RWMutex{},
		clients: []*WebsocketConnection{},
	}
}

type WebsocketConnection struct {
	conn   *websocket.Conn
	gameId uuid.UUID
}

func (ws *WebsocketConnection) Listen() {
	for {
		if _, _, err := ws.conn.NextReader(); err != nil {
			ws.conn.Close()
			break
		}
	}
}

func (c *ClientConnections) Append(client *WebsocketConnection) {
	c.Lock()
	defer c.Unlock()
	c.clients = append(c.clients, client)
}
