package ws

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer(t *testing.T) {
	r := gin.Default()
	r.GET("/ws", handleWS)
	r.Run(":8000")
}
