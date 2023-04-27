package rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	{
		i, err := Int64(255)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(i)
	}
	{
		i, err := Float64()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(i)
	}
	{
		s, err := String(18, DigitsAndLowerAlphAndUpperAlph)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(s)
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := String(16, DigitsAndLowerAlphAndUpperAlph)
		if err != nil {
			return
		}
		fmt.Println(s)
	}
}
