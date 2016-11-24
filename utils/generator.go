package utils

import "math/rand"

// GeneratorUID64 used to generated uid 64 bit
func GeneratorUID64() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	runes := make([]rune, 11)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}

	return string(runes)
}
