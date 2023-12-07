package util

import (
	"math/rand"
)

var numbers    = []rune("0123456789")
var letters    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var mixLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateCode(size int) string {
	l := make([]rune, 3)
	n := make([]rune, 3)
	for i := range l {
		l[i] = letters[rand.Intn(len(letters))]
	}
	for i := range n {
		n[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(l) + "-" + string(n)
}


func RandomString(size int) string {
	m := make([]rune, size)
	for i := range m {
		m[i] = letters[rand.Intn(len(letters))]
	}
	return string(m)
}