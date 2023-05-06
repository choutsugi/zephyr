package embed

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	p := NewPlayer()
	p.Teleport(100, 120)
	fmt.Println(p.poxX, p.poxY)

	p.Move(1, -2)
	fmt.Println(p.poxX, p.poxY)

}
