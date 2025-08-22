package config

import (
	"os"
	"strconv"
)

func parseMongo() MongoConfig {
	return MongoConfig{
		Url:    os.Getenv("MONGODB_URI"),
		DBName: "dev_main",
	}
}

func parseRedis() RedisConfig {
	port := 6379
	if os.Getenv("REDIS_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	}

	return RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: port,
	}
}

func GetEnv(attribute string) string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "DEV"
	}
	return env + "_" + os.Getenv(attribute)
}
