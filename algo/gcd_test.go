package algo

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	i := GCD([]int{5, 65, 0, 20})
	fmt.Println(i)
}
