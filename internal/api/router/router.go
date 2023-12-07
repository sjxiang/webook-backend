package router

import (
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
	// config
	engine.UseRawPath = true
	// init middleware
	engine.Use(middleware.Session(secret))
	engine.Use(middleware.CORS())
	
	engine.NoRoute(func(c *gin.Context) {
		c.Data(404, "text/plain", []byte("404 page not found"))
		c.Abort()
	})
	
	engine.NoMethod(func(c *gin.Context) {
		c.Data(405, "text/plain", []byte("Method Not Allowed"))
		c.Abort()
	})


	// init route
	routerGroup := engine.Group("/api/v1")
	userGroup := routerGroup.Group("/user")

	// session + cookie 认证
	{
		// userGroup.Use(middleware.NewSessionLoginMiddlewareBuilder().IgnorePaths("/api/v1/user/signup", "/api/v1/user/login").Build())
		userGroup.POST("/signup", r.Controller.Signup)
		userGroup.POST("/login", r.Controller.Login)
		// 中间件顺序，不要乱
		userGroup.Use(middleware.NewSessionLoginMiddlewareBuilder().Build())
		userGroup.POST("/profile", r.Controller.Profile)
		userGroup.POST("/edit", r.Controller.Edit)
	}
	

	// init route
	otherRouterGroup := engine.Group("/api/v2")
	otherUserGroup := otherRouterGroup.Group("/user")
	
	// jwt 认证
	{
		otherRouterGroup.Use(middleware.JwtAuthMiddleware(r.Controller.ExportTokenMaker()))
		otherUserGroup.POST("/signup", nil)
	}

}