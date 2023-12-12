package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sjxiang/webook-backend/internal/api/middleware"
	"github.com/sjxiang/webook-backend/internal/api/router"
	"github.com/sjxiang/webook-backend/internal/conf"
	"github.com/sjxiang/webook-backend/pkg/limiter"
)



type Server struct {
	engine  *gin.Engine
	router  *router.Router
	logger  *zap.SugaredLogger
	config  *conf.Config
	limiter  limiter.Limiter
}


func NewServer(config *conf.Config, engine *gin.Engine, router *router.Router, logger *zap.SugaredLogger, limiter limiter.Limiter) *Server {
	return &Server{
		engine:  engine,
		config:  config,
		router:  router,
		logger:  logger,
		limiter: limiter,
	}
}

func (server *Server) Start() {
	server.logger.Infow("Starting webook-backend...")
	secret := server.config.GetSecretKey()
	
	// init
	gin.SetMode(server.config.ServerMode)

	// init global middleware
	server.engine.Use(middleware.CORS())
	server.engine.Use(middleware.NewRateLimiteBuilder(server.limiter).Build())
	
	// init route
	server.router.RegisterRouters(server.engine, secret)

	// run
	err := server.engine.Run(server.config.ServerHost + ":" + server.config.ServerPort)
	if err != nil {
		server.logger.Errorw("Error in startup", "err", err)
		os.Exit(2)
	}
}