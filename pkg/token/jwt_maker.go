package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("无效密钥长度：最少 %d 个字符", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(id int64, email string, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(id, email, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", payload, err
	}

	return token, payload, nil 
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken  // 过期
		}
		return nil, ErrInvalidToken  // 可能签名或者密钥不对，无效
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken  // 类型断言拉了，无效
	}

	return payload, nil
}
