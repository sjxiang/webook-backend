package limiter

import (
	"context"
	_ "embed"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed slide_window.lua
var luaScript string

type RedisSlidingWindowLimiter struct {
	cmd      redis.Cmdable
	interval time.Duration
	// 阈值
	rate     int
}

func NewRedisSlidingWindowLimiter(cmd redis.Cmdable, interval time.Duration, rate int) *RedisSlidingWindowLimiter {
	return &RedisSlidingWindowLimiter{
		cmd:      cmd,
		interval: interval, // 比方一分钟上限 100 个请求
		rate:     rate,
	}
}

// 检查是否需要限流
func (b *RedisSlidingWindowLimiter) Limit(ctx context.Context, key string) (bool, error) {
	return b.cmd.Eval(
		ctx, 
		luaScript, 
		[]string{key},
		b.interval.Milliseconds(),  // arg 1
		b.rate,                     // arg 2
		time.Now().UnixMilli(),     // arg 3
	).Bool()
}