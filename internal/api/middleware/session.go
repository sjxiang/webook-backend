package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化 session
func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{
			HttpOnly: true, 
			MaxAge:   7 * 86400, 
			Path:     "/",
		})

	return sessions.Sessions("gin-session", store)
}


/*

每个请求过来，给它发个碗，方便要饭（ session 结构体）

func Sessions(name string, store Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := &session{name, c.Request, store, nil, false, c.Writer}
		c.Set(DefaultKey, s)
		defer context.Clear(c.Request)
		c.Next()
	}
}

这个仓库就是对 http.Cookie 的封装，更友好

*/