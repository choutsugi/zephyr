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
}
