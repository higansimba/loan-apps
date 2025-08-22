package initializers

import (
	"context"

	"github.com/higansama/loan-apps/config"
)

// InitDatabase initializes all database connections and returns a DBConfig
func InitDatabase(ctx context.Context, env *config.EnvConfig) (*config.DBConfig, error) {
	// Initialize Redis
	redisClient := InitRedis(env)

	// Initialize Postgres
	mongoConnection := InitMongoDB(ctx, env)

	// Return combined database configuration
	return &config.DBConfig{
		MongoConnection: mongoConnection,
		RedisConnection: redisClient,
	}, nil
}
