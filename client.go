package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn // Websocket for this client
	send   chan []byte     // is a channel on which messages are sent
	room   *room           // is the room this client is chatting in
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
