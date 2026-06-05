package router

import (
	"net/http"

	"github.com/Artymka/effective-mobile-test-task/internal/handlers"
	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/go-playground/validator/v10"
)

func New(repo *repository.SubscriptionRepository,
	valid *validator.Validate,
	log *lib.Logger) *http.ServeMux {

	subHs := handlers.SubscriptionHandlers{
		Repo:  repo,
		Valid: valid,
		Log:   log,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", subHs.Create)

	return mux
}
