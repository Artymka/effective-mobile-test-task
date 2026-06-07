package main

import (
	"fmt"
	"net/http"

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
	conf := config.Load()

	logger := lib.NewLogger()

	// logger.Info("check db config", conf.GetDBConnectionString())
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

	logger.Info("main", "starting server...")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", conf.ServerHost, conf.ServerPort), *mux); err != nil {
		logger.Error("main", err)
	}
}
