package redis

import (
	"context"
	"fmt"
	"time"

	"easy-fiber-admin/pkg/config"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// InitRedis initializes the Redis client
func InitRedis(cfg config.RedisConfig) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Ping the Redis server to ensure the connection is established
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return rdb, nil
}

// GetClient returns the global Redis client
func GetClient() *redis.Client {
	return rdb
}

// Set stores a key-value pair in Redis with an expiration time
func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value from Redis by key
func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", fmt.Errorf("failed to get key %s from Redis: %w", key, err)
	}
	return val, nil
}

// Del deletes a key from Redis
func Del(key string) error {
	return rdb.Del(ctx, key).Err()
}
