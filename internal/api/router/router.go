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


func (r *Router) RegisterRouters(engine *gin.Engine) {
	// config
	engine.UseRawPath = true
	// init middleware
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

	userGroup.POST("/signup", r.Controller.Signup)
	userGroup.POST("/me", r.Controller.Me)
	
}