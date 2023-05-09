package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWS(ctx *gin.Context) {

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(msg) == "ping" {
			if err := ws.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
				log.Printf("failed to write msg with err: %+v\n", err)
				break
			}
			continue
		}
		log.Printf("received data: %s\n", string(msg))
		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("failed to write msg with err: %+v\n", err)
			break
		}
	}
}
