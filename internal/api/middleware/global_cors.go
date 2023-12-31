package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// 全局

// Cors 跨域配置
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	// 业务请求中可以带上的头
	config.AllowHeaders = []string{
		"Origin", 
		"Content-Length", 
		"Content-Type", 
		"Cookie", 
		"Authorization",
	}
	// 把 jwt 放在 header，而不是 body 中，需要暴露（尽管真实返回了，但浏览器不给前端，所以这个设置是给浏览器看的） 
	config.ExposeHeaders = []string{"x-jwt-token"}

	// 哪些来源是允许的
	// 写法 1
	// config.AllowOrigins = []string{"http://localhost:8080", "https://www.yourcompany.com"}
	// 写法 2
	config.AllowOriginFunc = func(origin string) bool {
		if strings.HasPrefix(origin, "http://localhost") {
			// 你的开发环境
			return true
		}
		return strings.Contains(origin, "yourcompany.com")
	}
	// 是否允许带上用户认证信息，比如 cookie
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return cors.New(config)
}