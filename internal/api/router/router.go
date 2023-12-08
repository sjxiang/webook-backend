package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/webook-backend/internal/api/controller"
	"github.com/sjxiang/webook-backend/internal/api/middleware"
)

type Router struct {
	Controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{
		Controller: controller,
	}
}


func (r *Router) RegisterRouters(engine *gin.Engine, secret string) {
	// 404
	engine.NoRoute(func(c *gin.Context) {
		c.Data(404, "text/plain", []byte("404 page not found"))
		c.Abort()
	})
	engine.NoMethod(func(c *gin.Context) {
		c.Data(405, "text/plain", []byte("Method Not Allowed"))
		c.Abort()
	})

	// config
	engine.UseRawPath = true
	// init middleware
	engine.Use(middleware.CORS())
	

	// init v1 route
	routerGroup := engine.Group("/api/v1")
	userGroup := routerGroup.Group("/user")
	
	// init session
	store := cookie.NewStore([]byte(secret))
	// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{
			HttpOnly: true, 
			MaxAge:   7 * 86400,  // 7 天
			Path:     "/",
		})
	userGroup.Use(sessions.Sessions("gin-session", store))
	
	// Session + Cookie
	{
		// userGroup.Use(middleware.NewSessionLoginMiddlewareBuilder().IgnorePaths("/api/v1/user/signup", "/api/v1/user/login").Build())
		userGroup.POST("/signup", r.Controller.Signup)
		userGroup.POST("/login", r.Controller.Login)
		// 中间件顺序，不要乱
		userGroup.Use(middleware.NewSessionLoginMiddlewareBuilder().Build())
		userGroup.POST("/profile", r.Controller.Profile)
		userGroup.POST("/edit", r.Controller.Edit)
		userGroup.POST("/logout", r.Controller.Logout)
	}
	

	// init v2 route
	otherRouterGroup := engine.Group("/api/v2")
	otherUserGroup := otherRouterGroup.Group("/user")
	
	// JWT
	{
		otherRouterGroup.Use(middleware.JwtAuthMiddleware(r.Controller.ExportTokenMaker()))
		otherUserGroup.POST("/signup", nil)
	}

}