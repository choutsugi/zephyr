package chat

import (
	"golang.org/x/net/websocket"
	"net/http"
	"testing"
)

/*
	let socket = new WebSocket("ws://localhost:3000/ws")
	socket.onmessage = (event) => { console.log("received from server: ", event.data) }
	socket.send("halo")
*/

func TestChat(t *testing.T) {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleWS))
	http.Handle("/tick", websocket.Handler(server.HandleTick))
	http.ListenAndServe(":3000", nil)
}
