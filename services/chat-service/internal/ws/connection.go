package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	Socket *websocket.Conn
	mu     sync.Mutex
}

// Send Concurrency handling - sending messages
func (c *Connection) Send(data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Socket.WriteMessage(websocket.TextMessage, data)
}
