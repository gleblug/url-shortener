package random

import (
	"math/rand"
	"time"
)

// TODO: write tests
func NewRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	symbols := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@#$%&")

	buf := make([]rune, length)
	for i := range buf {
		buf[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(buf)
}
