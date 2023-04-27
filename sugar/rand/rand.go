package rand

import (
	"crypto/rand"
	"math/big"
)

// Int64 returns a uniform random value in [0, max). It panics if max <= 0.
func Int64(max int64) (int64, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return random.Int64(), nil
}

// Float64 returns a uniform random value in [0.0, 1.0).
func Float64() (float64, error) {
	random, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0, err
	}
	return float64(random.Int64()) / (1 << 53), nil
}
