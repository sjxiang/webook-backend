package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			// 不填，默认全支持
			AllowMethods: []string{"POST", "GET"},  
			// 业务请求中可以带上的头
			AllowHeaders: []string{"Content-Type", "Authorization"},
			// 是否允许带上用户认证信息，比如 cookie
			AllowCredentials: true,
			// 哪些来源是允许的
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					// 你的开发环境
					return true
				}
				return strings.Contains(origin, "yourcompany.com")
			},
			// AllowOrigins: []string{"http://localhost:3000"},
			MaxAge: 12 * time.Hour,
		})
	}
}
