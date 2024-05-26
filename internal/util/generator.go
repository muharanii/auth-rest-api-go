package util

import "math/rand"

func GenerateRandomString(n int) string {
	var charsets = []rune("abcdefhigklmopqrstupwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(letters)
}
