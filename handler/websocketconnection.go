package handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
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

func (c *ClientConnections) Listen(ws *WebsocketConnection) {
	for {
		if _, _, err := ws.conn.NextReader(); err != nil {
			c.Remove(ws)
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

func (c *ClientConnections) Remove(client *WebsocketConnection) {
	c.Lock()
	defer c.Unlock()
	for i, val := range c.clients {
		if val.conn == client.conn {
			c.clients[i] = c.clients[len(c.clients)-1]
			c.clients[len(c.clients)-1] = nil
			c.clients = c.clients[:len(c.clients)-1]
			log.Println("client removed")
		}
	}
}
