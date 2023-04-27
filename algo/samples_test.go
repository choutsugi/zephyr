package algo

import (
	"fmt"
	"testing"
)

func TestSamples(t *testing.T) {
	words := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "pear", "quince", "raspberry", "strawberry", "tangerine", "ugli fruit", "watermelon"}
	samples := Samples(words, 5)
	fmt.Println(samples)
}

func BenchmarkSamples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		words := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "pear", "quince", "raspberry", "strawberry", "tangerine", "ugli fruit", "watermelon"}
		Samples(words, 5)
	}
}
