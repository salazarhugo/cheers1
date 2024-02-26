package tests

import (
	"crypto/rand"
	"math/big"
)

func CreateRandomUsername(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	var result string
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result += string(charset[randomIndex.Int64()])
	}
	return result
}
