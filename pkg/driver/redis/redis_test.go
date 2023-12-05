package redis

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewRedisConnection(t *testing.T) {
	redisConfig := &RedisConfig{
		Addr:     "localhost",
		Port:     "6379",
		Password: "",
		Database: 0,
	}
	cache, err := NewRedisConnection(redisConfig, zap.L().Sugar())
	require.NoError(t, err)
	require.NotEmpty(t, cache)
}