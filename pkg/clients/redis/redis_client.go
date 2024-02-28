package clients

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) Cache {
	return &RedisClient{
		client: client,
	}
}

func (c *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}
