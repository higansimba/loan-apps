package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*EnvConfig, error) {
	envPath := ".env"
	_, err := os.Stat(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			cfgFile, err := os.OpenFile(envPath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			cfgFile.Close()
		}
	}

	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	// Validate environment variables
	// fmt.Println()
	// requiredEnvVars := []string{
	// 	"ENV",
	// 	"APP_HOST",
	// 	"DEV_SERVER_PORT",
	// 	"MONGODB_URI",
	// 	// "BEARER_TOKEN",
	// 	"ADMIN_ENDPOINT",
	// }

	// for _, envVar := range requiredEnvVars {
	// 	v := os.Getenv(envVar)
	// 	if v == "" {
	// 		panic(fmt.Sprintf("Environment variable %s is required but not set", envVar))
	// 	}
	// }

	// Parse configurations
	return &EnvConfig{
		MongoConfig: parseMongo(),
		Redis:       parseRedis(),
		Debug:       os.Getenv("ENV") == "PROD",
		Server: struct {
			Host string
			Port string
		}{
			Host: GetEnv("APP_HOST"),
			Port: GetEnv("SERVER_PORT"),
		},
	}, nil
}
