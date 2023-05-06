package multiwriter

import "testing"

func TestMultiWriter(t *testing.T) {
	s := NewServer()
	s.broadcast([]byte("foo"))
}
