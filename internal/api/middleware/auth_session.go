package middleware

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SessionLoginMiddlewareBuilder struct  {
	// 俺怀疑大明不知道 gin 中间件顺序的骚操作，纯属脱裤子放屁，多此一举
	publicPaths []string
}

func (impl *SessionLoginMiddlewareBuilder) IgnorePaths(path ...string) *SessionLoginMiddlewareBuilder {
	impl.publicPaths = append(impl.publicPaths, path...)
	return impl
}

func NewSessionLoginMiddlewareBuilder() *SessionLoginMiddlewareBuilder {
	return &SessionLoginMiddlewareBuilder{
		publicPaths: make([]string, 0),
	}
}

func (impl *SessionLoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Time{})
	return func(ctx *gin.Context) {

		// 不需要校验
		for _, path := range impl.publicPaths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(ctx)
		uid := sess.Get("user_id")
		lastTime := sess.Get("last_time")

		if uid == nil || lastTime == nil {
			// 没有登录
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "需要登录",
			})
			return
		}

		// // 会话状态保持
		if last, ok := lastTime.(time.Time); ok {
			if time.Since(last) > time.Minute*30 {  // 还有 30 min，刷新，再给你续一轮
				sess.Options(sessions.Options{
					MaxAge: 7 * 86400,
				})
				sess.Save()
				zap.L().Info("续了一轮")
			} 	
		}

		ctx.Set("user_id", uid.(int64))
		ctx.Next()
	}
}

