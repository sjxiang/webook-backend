package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	return func(ctx *gin.Context) {

		// 不需要校验
		for _, path := range impl.publicPaths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(ctx)
		uid := sess.Get("user_id")
		
		if uid == nil  {
			// 没有登录
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "需要登录",
			})
			return
		}

		ctx.Set("user_id", uid)
		ctx.Next()
	}
}




		
// 		// 如果是空字符串，你可以预期后面 Parse 就会报错
// 		tokenStr := impl.ExtractTokenString(ctx)
// 		uc := ijwt.UserClaims{}
// 		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
// 			return ijwt.AccessTokenKey, nil
// 		})
// 		if err != nil || !token.Valid {
// 			// 不正确的 token
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		expireTime, err := uc.GetExpirationTime()
// 		if err != nil {
// 			// 拿不到过期时间
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		if expireTime.Before(time.Now()) {
// 			// 已经过期
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		if ctx.GetHeader("User-Agent") != uc.UserAgent {
// 			// 换了一个 User-Agent，可能是攻击者
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		err = impl.CheckSession(ctx, uc.Ssid)
// 		if err != nil {
// 			// 系统错误或者用户已经主动退出登录了
// 			// 这里也可以考虑说，如果在 Redis 已经崩溃的时候，
// 			// 就不要去校验是不是已经主动退出登录了。
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		// 说明 token 是合法的
// 		// 我们把这个 token 里面的数据放到 ctx 里面，后面用的时候就不用再次 Parse 了
// 		ctx.Set("user", uc)
// 	}
// }