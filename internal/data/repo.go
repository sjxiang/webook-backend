package data

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/sjxiang/webook-backend/internal/biz"
)



type userRepo struct {
	storage *gorm.DB
	cache   *redis.Client
	logger *zap.SugaredLogger
}

func NewUserRepo(storage *gorm.DB, cache *redis.Client, logger *zap.SugaredLogger) biz.UserRepo {
	return &userRepo{
		storage: storage,
		cache:   cache,
		logger:  logger ,
	}
}