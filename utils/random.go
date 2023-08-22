package utils

import (
	"math/rand"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func RandomID() int32 {
	// Generate a random int32 within a specific range
	return rand.Int31n(1000) + 1 // Add 1 to ensure the ID is not zero
}
