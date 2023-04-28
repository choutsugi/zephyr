package algo

import "math/rand"

// Samples returns N random unique items from collection.
func Samples(original []any, k int) []any {
	length := len(original)
	if k > length {
		k = length
	}
	items := make([]any, length)
	copy(items, original)
	indices := rand.Perm(length)[:k]
	samples := make([]any, k)
	for i, idx := range indices {
		samples[i] = items[idx]
		items[idx], items[length-1-i] = items[length-1-i], items[idx]
	}
	return samples
}
