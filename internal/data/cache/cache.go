package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"goex1/internal/conf"
)

var redisClient *redis.Client

func Redis() *redis.Client {
	return redisClient
}

func InitCache(cfg *conf.Conf) {
	// redisConf := config.Conf.Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.Db,
		PoolSize:     cfg.PoolSize,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		PoolTimeout:  10 * time.Second,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
