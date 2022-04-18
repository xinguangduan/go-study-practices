package performance

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func rand01(n int) string {
	bytes := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Int31()%52]
	}
	return string(result)
}

func rand02(n int) string {
	result := make([]byte, n/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}
