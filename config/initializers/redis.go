package initializers

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/higansama/loan-apps/config"

	"github.com/go-redis/redis/v8"
)

func InitRedis(env *config.EnvConfig) config.RedisConnection {
	var tLSConfig *tls.Config

	islocalhost := env.Redis.Host == "localhost"
	password := ""
	if !islocalhost {
		password = env.Redis.Password
	}

	isSkipTLS := env.Redis.SkipTLS
	if isSkipTLS {
		tLSConfig = &tls.Config{
			InsecureSkipVerify: env.Redis.SkipTLS,
		}
	}

	// Initialize Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:      fmt.Sprintf("%s:%d", env.Redis.Host, env.Redis.Port),
		DB:        env.Redis.DB,
		Password:  password,
		TLSConfig: tLSConfig,
	})

	// Ping Redis
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return config.RedisConnection{
		RedisMaster: rdb,
	}
}
