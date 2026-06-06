package main

import (
	"net/http"

	"github.com/Artymka/effective-mobile-test-task/internal/config"
	"github.com/Artymka/effective-mobile-test-task/internal/database"
	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/Artymka/effective-mobile-test-task/internal/router"
	"github.com/go-playground/validator/v10"
)

func main() {
	conf := config.Load()

	db, err := database.NewPostgresDB(conf.GetDBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewSubscriptionRepository(db.DB, conf)

	valid := validator.New()

	logger := lib.NewLogger()

	mux := router.New(repo, valid, logger, conf)

	http.ListenAndServe("localhost:8000", mux)
}
