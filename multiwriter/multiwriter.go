package multiwriter

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"
)

type Conn struct {
	io.Writer
}

func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}

func (c *Conn) Write(p []byte) (int, error) {
	fmt.Printf("time=%s write data: %s\n", time.Now().Format(time.RFC3339Nano), string(p))
	return c.Writer.Write(p)
}

type Server struct {
	peers map[*Conn]bool
}

func NewServer() *Server {
	s := &Server{
		peers: make(map[*Conn]bool),
	}
	for i := 0; i < 10; i++ {
		s.peers[NewConn()] = true
	}
	return s
}

func (s *Server) broadcast(msg []byte) error {

	var peers []io.Writer
	for peer := range s.peers {
		peers = append(peers, peer)
	}
	if _, err := io.MultiWriter(peers...).Write(msg); err != nil {
		log.Fatal(err)
	}

	/*
		for peer := range s.peers {
			peer.Write(msg)
		}
	*/

	return nil
}
