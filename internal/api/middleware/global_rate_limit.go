package middleware

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/webook-backend/pkg/limiter"
)

// 全局

type ctxKey string
const (
	XStressKey ctxKey = "x-stress"
)

type RateLimitBuilder struct {
	prefix  string
	limiter limiter.Limiter

}

func NewRateLimiteBuilder(l limiter.Limiter) *RateLimitBuilder {
	return &RateLimitBuilder{
		prefix:  "ip-limiter",
		limiter: l,
	}
}

func (b *RateLimitBuilder) Prefix(prefix string) *RateLimitBuilder {
	b.prefix = prefix
	return b
}

func (b *RateLimitBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.GetHeader("x-stress") == "true" {
			// 用 context.Context 来带这个标记位
			newCtx := context.WithValue(ctx, XStressKey, true)
			ctx.Request = ctx.Request.Clone(newCtx)
			ctx.Next()
			return
		}

		limited, err := b.limiter.Limit(ctx, fmt.Sprintf("%s:%s", b.prefix, ctx.ClientIP()))
		if err != nil {
			log.Println(err)
			// 这一步很有意思，就是如果这边出错了
			// 要怎么办？
			// 保守做法：因为借助于 Redis 来做限流，那么 Redis 崩溃了，为了防止系统崩溃，直接限流
			ctx.AbortWithStatus(http.StatusInternalServerError)
			// 激进做法：虽然 Redis 崩溃了，但是这个时候还是要尽量服务正常的用户，所以不限流
			// ctx.Next()
			return
		}
		if limited {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "猴急，慢点",
			})
			return
		}
		ctx.Next()
	}
}