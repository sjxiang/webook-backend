package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth2Claims struct {
	Identity  int 
	Field     string
	jwt.RegisteredClaims   
}


// 生成
func GenerateAuth2Token(identity int, field string, duration time.Duration, secretKey string) (string, error) {

	claims := &Auth2Claims{
		Identity:         identity,
		Field:            field,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(duration)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签发
	return token.SignedString([]byte(secretKey))
}


// 提取
func ExtractAuth2Token(accessToken, secretKey string) (claims *Auth2Claims, err error) {

	token, err := jwt.ParseWithClaims(accessToken, &Auth2Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 
	}

	claims, ok := token.Claims.(*Auth2Claims)
	if !(ok && token.Valid) {
		return 
	}

	return claims, nil
}

