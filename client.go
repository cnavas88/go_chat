package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn // Websocket for this client
	send   chan []byte     // is a channel on which messages are sent
	room   *room           // is the room this client is chatting in
}
