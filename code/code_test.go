package code

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewCode(t *testing.T) {
	c := NewCode(1, "internal server error", errors.New("something happened"))
	fmt.Println(c)
}
