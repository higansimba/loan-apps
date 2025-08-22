package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/higansama/loan-apps/config"
	"github.com/higansama/loan-apps/config/initializers"
	"github.com/higansama/loan-apps/internal/bootstrap"
	"github.com/higansama/loan-apps/internal/logger"
	transactionmongo "github.com/higansama/loan-apps/internal/transaction_mongo"

	"github.com/higansama/loan-apps/internal/session"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	os.Setenv("TZ", "Asia/Jakarta")

	logger.InitLogger(&logger.Config{})

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

	env, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	db, err := initializers.InitDatabase(ctx, env)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	ginApp := gin.Default()

	// setup session
	ginApp = session.InitSession(ginApp)

	// Initialize application components
	txManager := transactionmongo.NewTransactionManager(db.MongoConnection.MainDB.Client())
	repos := bootstrap.InitRepositories(&db.MongoConnection)
	useCases := bootstrap.InitUseCases(repos, *txManager)

	_ = bootstrap.InitHandlers(ginApp, useCases)

	srv := &http.Server{
		Addr:    ":8888",
		Handler: ginApp.Handler(),
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	log.Println("Server running:8888")

	<-ctx.Done()

	stop()

	log.Println("Server graceful shutdown")

	nctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(nctx); err != nil {
		log.Fatal(err)
	}
}
