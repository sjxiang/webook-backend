package cache

import (
	"context"

	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	XX_KEY = "ip_zone"
)

type IPZoneCache struct {
	logger  *zap.SugaredLogger
	cache   *redis.Client
	context context.Context
}

func NewXxxCache(cache *redis.Client, logger *zap.SugaredLogger) *IPZoneCache {
	return &IPZoneCache{
		logger:  logger,
		cache:   cache,
		context: context.Background(),
	}
}

func (c *IPZoneCache) SetXxx(ip string, zone string) error {
	return c.cache.HSet(c.context, XX_KEY, ip, zone).Err()
}

func (c *IPZoneCache) GetIPZone(ip string) (string, error) {
	zone, errInGet := c.cache.HGet(c.context, XX_KEY, ip).Result()
	if errInGet == redis.Nil {
		return "", nil
	} else if errInGet != nil {
		return "", errInGet
	}
	return zone, nil
}



// <!-- 

// redis 处理

// var ErrNotFound = errors.New("not found")

// func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
// 	v, err := c.conn.Get(ctx, key).Result()
// 	if errors.Is(err, redis.Nil) {
// 		return nil, ErrNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return []byte(v), nil
// }

// func (c *Cache) Set(ctx context.Context, key string, value []byte) error {
// 	return c.conn.Set(ctx, key, value, c.ttl).Err()
// } -->
