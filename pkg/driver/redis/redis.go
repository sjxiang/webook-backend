package redis

import (
	"context"

	redis "github.com/redis/go-redis/v9"
	"github.com/sjxiang/webook-backend/internal/conf"
	"go.uber.org/zap"
)

const RETRY_TIMES = 6

type RedisConfig struct {
	Addr     string 
	Port     string 
	Password string 
	Database int    
}

func NewRedisConnectionByGlobalConfig(config *conf.Config, logger *zap.SugaredLogger) (*redis.Client, error) {
	redisConfig := &RedisConfig{
		Addr:     config.GetRedisAddr(),
		Port:     config.GetRedisPort(),
		Password: config.GetRedisPassword(),
		Database: config.GetRedisDatabase(),
	}
	return NewRedisConnection(redisConfig, logger)
}

func NewRedisConnection(config *RedisConfig, logger *zap.SugaredLogger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Database,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		logger.Errorw("error in connecting redis ", "redis", config, "err", err)
		return nil, err 
	}

	logger.Infow("connected with redis", "redis", config)

	return rdb, nil
}
