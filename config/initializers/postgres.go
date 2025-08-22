package initializers

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/higansama/loan-apps/config"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func InitPostgres(env *config.EnvConfig) config.PostgresConnections {
// 	// Initialize multiple Postgres connections
// 	connections := config.PostgresConnections{}

// 	// LoyalPoint Master database connection
// 	sortHubMasterDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
// 		env.PostgresLoyalPointMaster.Host,
// 		env.PostgresLoyalPointMaster.User,
// 		env.PostgresLoyalPointMaster.Password,
// 		env.PostgresLoyalPointMaster.Database,
// 		env.PostgresLoyalPointMaster.Port)

// 	sortHubMasterDB, err := initializePostgresDB(sortHubMasterDSN)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to postgres sorthub master database: %v", err)
// 	}
// 	connections.LoyalPointMaster = sortHubMasterDB

// 	return connections
// }

// func initializePostgresDB(dsn string) (*gorm.DB, error) {
// 	// Configure GORM with query logging
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %v", err)
// 	}

// 	// Configure connection pool
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to configure connection pool: %v", err)
// 	}

// 	// Set connection pool settings
// 	sqlDB.SetMaxIdleConns(10)           // Maximum number of idle connections
// 	sqlDB.SetMaxOpenConns(10)           // Maximum number of open connections
// 	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection

// 	return db, nil
// }
