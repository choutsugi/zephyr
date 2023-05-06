package chat

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"time"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	log.Printf("new incoming connection, addr = %s\n", ws.RemoteAddr().String())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) HandleTick(ws *websocket.Conn) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		ws.Write([]byte(fmt.Sprintf("tick => %d", time.Now().Unix())))
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("read error: %v\n", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(bytes []byte) {
	for conn := range s.conns {
		go func(conn *websocket.Conn) {
			if _, err := conn.Write(bytes); err != nil {
				log.Printf("write error: %v\n", err)
			}
		}(conn)
	}
}
