package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/sjxiang/webook-backend/pkg/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)


func JwtAuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "未提供 authorization header",
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效 authorization header 格式",
			})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("不支持 authorization 类型 %s", authorizationType),
			})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			if errors.Is(err, token.ErrExpiredToken) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "token 过期",
				})
			}
			if errors.Is(err, token.ErrInvalidToken) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "token 被篡改，无效",
				})
				return
			}

			zap.S().Error("middleware", "校验 token 异常", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err,
			})
			return
		}

		ctx.Set("authorization_payload", payload)
		ctx.Next()
	}
}
