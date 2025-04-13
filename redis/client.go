package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Redis *redis.Client

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg Config) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	return &RedisClient{Client: rdb}
}

func SetupRedisClient() error {
	cfg := LoadConfig()
	rdb := NewRedisClient(cfg)

	// Test the connection
	if err := rdb.Client.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	Redis = rdb.Client
	return nil
}

func (r *RedisClient) Set(key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
