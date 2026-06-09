package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Artymka/effective-mobile-test-task/internal/config"
	"github.com/Artymka/effective-mobile-test-task/internal/database"
	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/Artymka/effective-mobile-test-task/internal/router"
	"github.com/go-playground/validator/v10"
)

// @title           Subscription Service API
// @termsOfService  http://swagger.io/terms/

func main() {
	// context for detecting SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM)
	defer stop()

	conf := config.Load()
	logger := lib.NewLogger()
	logger.Info("main", "config loaded")

	db, err := database.NewPostgresDB(conf.GetDBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	logger.Info("main", "connected to postgres")

	repo := repository.NewSubscriptionRepository(db.DB, conf)

	valid := validator.New()

	mux := router.New(repo, valid, logger, conf)
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort),
		Handler: *mux,
	}
	go func() {
		logger.Info("main", "starting server...")
		if err := server.ListenAndServe(); err != nil {
			logger.Error("main", err)
		}
	}()

	// graceful shutdown
	<-ctx.Done()
	logger.Info("main", "shutting down...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	// server shutdown
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("server shutdown", err)
	}
	logger.Info("main", "server shutdown done")

	// postgres shutdown
	if err := db.Close(); err != nil {
		logger.Error("postgres shutdown", err)
	}
	logger.Info("main", "postgres shutdown done")
}
