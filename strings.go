package random

import (
	"math/rand"
	"time"
)

const (
	alphaNum string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a random alphanumeric string of a given length
func RandomAlphaNum(length int) string {
	bytes := make([]byte, length)
	for i, _ := range bytes {
		r := rand.Int()
		bytes[i] = byte(alphaNum[r%len(alphaNum)])
	}

	return string(bytes)
}
