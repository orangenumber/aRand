package arand

import (
	"math/rand"
	"time"
)

const (
	R_ALPHA              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	R_NUMERIC            = "0123456789"
	R_ALPHANUMERIC       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	R_ALPHANUMERIC_LOWER = "abcdefghijklmnopqrstuvwxyz0123456789"
	R_ALPHANUMERIC_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	R_HEX                = "0123456789ABCDEF"
)

func init() {
	UpdateSeed()
}

func UpdateSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandStr(chars string, length int) string {
	bytes := make([]byte, length)
	charsLen := len(chars)
	for i := 0; i < length; i++ {
		bytes[i] = chars[RandInt(0, charsLen)]
	}
	return string(bytes)
}
