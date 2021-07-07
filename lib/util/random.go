package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefgijklmnopqrstuvwxyz1234567890.#!$&_+"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPassword() string {
	return RandomString(8)
}
