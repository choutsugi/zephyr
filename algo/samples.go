package algo

import "math/rand"

// Samples returns N random unique items from collection.
func Samples(original []interface{}, k int) []interface{} {
	length := len(original)
	if k > length {
		k = length
	}
	items := make([]interface{}, length)
	copy(items, original)
	indices := rand.Perm(length)[:k]
	samples := make([]interface{}, k)
	for i, idx := range indices {
		samples[i] = items[idx]
		items[idx], items[length-1-i] = items[length-1-i], items[idx]
	}
	return samples
}
