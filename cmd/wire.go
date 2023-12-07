package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sjxiang/webook-backend/internal/api/controller"
	"github.com/sjxiang/webook-backend/internal/api/router"
	"github.com/sjxiang/webook-backend/internal/biz"
	"github.com/sjxiang/webook-backend/internal/conf"
	"github.com/sjxiang/webook-backend/internal/data"
	"github.com/sjxiang/webook-backend/internal/data/cache"
	"github.com/sjxiang/webook-backend/pkg/driver/mysql"
	"github.com/sjxiang/webook-backend/pkg/driver/redis"
	"github.com/sjxiang/webook-backend/pkg/logger"
	"github.com/sjxiang/webook-backend/pkg/token"
)


func initStorage(globalConfig *conf.Config, logger *zap.SugaredLogger) *gorm.DB {
	mysqlDriver, err := mysql.NewMySQLConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, storage init failed.")
	}
	return mysqlDriver
}

func initCache(globalConfig *conf.Config, logger *zap.SugaredLogger) *cache.Cache {
	redisDriver, err := redis.NewRedisConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, cache init failed.")
	}
	return cache.NewCache(redisDriver, logger)
}

func initTokenMaker(globalConfig *conf.Config, logger *zap.SugaredLogger) token.Maker {
	tokenMaker, err := token.NewJWTMaker(globalConfig.GetRamdonKey())
	if err != nil {
		logger.Errorw("Error in startup, tokenMaker init failed.")
	}
	return tokenMaker
}

func initServer() (*Server, error) {
	globalConfig := conf.GetInstance()
	engine := gin.Default()
	sugaredLogger := logger.NewSugardLogger()

	// init token
	tokenMaker := initTokenMaker(globalConfig, sugaredLogger)

	// init storage & cache
	storage := initStorage(globalConfig, sugaredLogger)

	// init repo
	ur := data.NewUserRepo(storage, nil, sugaredLogger)

	// init usecase
	uc := biz.NewUserUsecase(ur, sugaredLogger)
	
	// init controller
	c := controller.NewControllerForBackend(uc, tokenMaker, sugaredLogger)

	router := router.NewRouter(c)
	server := NewServer(globalConfig, engine, router, sugaredLogger)
	return server, nil
}
