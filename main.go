package main

import (
	"fmt"
	"github.com/luvcurt/zephyr/sugar/rand"
)

func main() {
	s, err := rand.String(16, rand.DigitsAndLowerAlphAndUpperAlph)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
