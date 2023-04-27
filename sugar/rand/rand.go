package rand

import (
	"crypto/rand"
	"math/big"
	"unsafe"
)

type Charset = string

const (
	Digits                         Charset = "0123456789"
	LowerAlph                      Charset = "abcdefghijklmnopqrstuvwxyz"
	UpperAlph                      Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerAndUpperAlph                      = LowerAlph + UpperAlph
	DigitsAndLowerAlphAndUpperAlph         = Digits + LowerAndUpperAlph
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

// String generates a random string.
func String(length int, charset Charset) (string, error) {
	charsetLen := int64(len(charset))
	bytes := make([]byte, length)
	for i := range bytes {
		idx, err := Int64(charsetLen)
		if err != nil {
			return "", err
		}
		bytes[i] = charset[idx]
	}
	return *(*string)(unsafe.Pointer(&bytes)), nil
}
