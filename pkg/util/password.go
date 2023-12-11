package util

import (
	"fmt"
	"unicode"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	// 不可逆
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func PasswordMatches(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		switch {
		// 密码和哈希不匹配
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		// 其它错误
		default:
			return false, err
		}
	}

	return true, nil
}