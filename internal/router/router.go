package router

import (
	"net/http"

	"github.com/Artymka/effective-mobile-test-task/internal/config"
	"github.com/Artymka/effective-mobile-test-task/internal/handlers"
	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/go-playground/validator/v10"

	_ "github.com/Artymka/effective-mobile-test-task/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(repo *repository.SubscriptionRepository,
	valid *validator.Validate,
	log *lib.Logger,
	config *config.Config) *http.ServeMux {

	subHs := handlers.SubscriptionHandlers{
		Repo:   repo,
		Valid:  valid,
		Log:    log,
		Config: config,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /subscriptions", subHs.Create)
	mux.HandleFunc("GET /subscriptions", subHs.Get)
	mux.HandleFunc("PATCH /subscriptions", subHs.Update)
	mux.HandleFunc("DELETE /subscriptions", subHs.Delete)
	mux.HandleFunc("GET /subscriptions/list", subHs.List)
	mux.HandleFunc("GET /subscriptions/total-cost", subHs.TotalCost)

	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
