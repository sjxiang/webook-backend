package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sjxiang/webook-backend/pkg/util"
	"go.uber.org/zap"
)

const JWT_TOKEN_DEFAULT_EXIPRED_PERIOD = time.Minute * 30

type Auth2Claims struct {
	// 接口要实现方法太多，还是组合吧！
	jwt.RegisteredClaims   

	// 不要放敏感数据
	Identity  int 
	Field     string
}

/*

type Claims interface {
	GetExpirationTime() (*NumericDate, error)
	GetIssuedAt() (*NumericDate, error)
	GetNotBefore() (*NumericDate, error)
	GetIssuer() (string, error)
	GetSubject() (string, error)
	GetAudience() (ClaimStrings, error)
}

v5 需要实现这些

*/


// 生成 
func GenerateAuth2Token(identity int, field string, duration time.Duration, secretKey string) (string, error) {

	claims := &Auth2Claims{
		Identity:         identity,
		Field:            field,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(JWT_TOKEN_DEFAULT_EXIPRED_PERIOD)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	
	return tokenStr, nil
}


// 提取
func ExtractAuth2Token(accessToken, secretKey string) (*Auth2Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}

	token, err := jwt.ParseWithClaims(accessToken, &Auth2Claims{}, keyFunc)
	if err != nil {
		return nil, err  // 被篡改（签名和密钥对不上），无效
	}

	claims, ok := token.Claims.(*Auth2Claims)
	if !(ok && token.Valid) {
		return nil, err  // 类型断言，parse 无效
	}

	return claims, nil
}


type authHeader struct {
	AccessToken string `header:"Authorization" binding:"required,min=7"`  // 前缀 "Bearer "
}

func JWTAuth(secretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fetch content
		var req authHeader
		if err := ctx.ShouldBindHeader(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": util.Translate(err),
			})
			return
		}

		// 另一种实现
		// accessToken := c.Request.Header["Authorization"]
		
		
		segments := strings.Split(req.AccessToken, "Bearer ")

		if len(segments) != 2 { 
			ctx.AbortWithStatus(http.StatusUnauthorized)  // 格式不对
			return
		}
		claims, err := ExtractAuth2Token(segments[1], secretKey)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)  // 格式对，但内容被篡改
			return
		}

		// 会话保持
		expireTime, err := claims.GetExpirationTime()
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)  // 拿不到过期时间，不大可能发生
			return
		}
		if expireTime.Before(time.Now()) {
			ctx.AbortWithStatus(http.StatusUnauthorized)  // 已经过期
			return
		}
		if time.Until(expireTime.Time) < time.Minute*30 {  // 不足 30 min，续一轮
			zap.S().Info("快过期了，兄嘚")
			// 创建 token
			// 返回前端
		}

		ctx.Set("auth_claims", claims)
		ctx.Next()
	}
}

