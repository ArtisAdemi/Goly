package utils

import "math/rand"

var runes = []rune("012345789abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ")

func RandomURL(size int) string {
	str := make([]rune, size)

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}
