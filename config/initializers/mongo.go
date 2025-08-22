package initializers

import (
	"context"
	"os"
	"time"

	"github.com/higansama/loan-apps/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(ctx context.Context, env *config.EnvConfig) config.MongoConnections {
	option := options.Client().
		ApplyURI(os.Getenv("MONGODB_URI")).
		SetServerSelectionTimeout(10 * time.Second).
		SetConnectTimeout(10 * time.Second).
		SetSocketTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		panic(err)
	}

	return config.MongoConnections{
		MainDB: client.Database(env.MongoConfig.DBName),
	}
}
