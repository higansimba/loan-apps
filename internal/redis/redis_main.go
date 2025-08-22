package redis

import (
	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Test()
}

type RedisClientImpl struct {
	Client *redis.Client
}

// Test implements RedisClient.
func (r *RedisClientImpl) Test() {
	panic("unimplemented")
}

func NewRedisCache(client *redis.Client) RedisClient {
	return &RedisClientImpl{
		Client: client,
	}
}
