package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

// // ConfigError represents configuration-related errors
type ConfigError struct {
	Component string
	Message   string
	Err       error
}

func (e *ConfigError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Component, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Component, e.Message)
}

type MongoConfig struct {
	Url    string
	DBName string
}

// RedisConfig holds Redis connection configuration
type RedisConfig struct {
	Host     string
	Port     int
	DB       int
	Password string
	SkipTLS  bool
}

// // Package-level variables for singleton pattern

// // EnvConfig holds all environment configurations
type EnvConfig struct {
	// Redis  RedisConfig
	Debug       bool
	MongoConfig MongoConfig
	Redis       RedisConfig
	Server      struct {
		Host string
		Port string
	}
}

// // DBConfig holds database client instances
type DBConfig struct {
	MongoConnection MongoConnections
	RedisConnection RedisConnection
}

type MongoConnections struct {
	MainDB *mongo.Database
}

type RedisConnection struct {
	RedisMaster *redis.Client
}

// // validateEnv checks if all required environment variables are present
