package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client
var cfg *Config

func Init(c *Config) {
	cfg = c
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %s", err))
	}
}

func Get() *redis.Client {
	return client
}

// 常用操作封装
func Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return client.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	ctx := context.Background()
	return client.Get(ctx, key).Result()
}

func Del(keys ...string) error {
	ctx := context.Background()
	return client.Del(ctx, keys...).Err()
}

func Exists(keys ...string) (int64, error) {
	ctx := context.Background()
	return client.Exists(ctx, keys...).Result()
}

func Expire(key string, expiration time.Duration) error {
	ctx := context.Background()
	return client.Expire(ctx, key, expiration).Err()
}