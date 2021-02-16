package cache

//go:generate docker run -d -p 6379:6379 --name mo-redis --rm redis:6.0.10
import (
	"context"
	"github.com/ggdream/mochat-server/config"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// Get 获取redis句柄
func Get() *redis.Client {
	return rdb
}

// GetCtx 获取ctx
func GetCtx() context.Context {
	return ctx
}

// Init 初始化redis
func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr: config.RedisURI,
		DB: 0,
	})

	return rdb.Ping(ctx).Err()
}
