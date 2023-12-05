package util

import (
	"golang.org/x/crypto/bcrypt"
	"unicode"
)

func ValidatePassword(password string) (minSize, digit, special, lowercase, uppercase bool) {
	for _, c := range password {
		switch {
		// 数字
		case unicode.IsNumber(c):
			digit = true
		// 大写字母
		case unicode.IsUpper(c):
			uppercase = true
		// 小写字母
		case unicode.IsLower(c):
			lowercase = true
		// 特殊字符
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}
	minSize = len(password) >= 8
	return
}


func ValidatePasswordV1(password string) (minSize, digit, special, letter bool) {
	for _, c := range password {
		switch {
		// 数字
		case unicode.IsNumber(c):
			digit = true
		// 字母
		case unicode.IsUpper(c) || unicode.IsLower(c):
			letter = true
		// 特殊字符
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}
	minSize = len(password) >= 8
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}