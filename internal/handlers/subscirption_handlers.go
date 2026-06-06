package handlers

import (
	"github.com/Artymka/effective-mobile-test-task/internal/config"
	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/go-playground/validator/v10"
)

type SubscriptionHandlers struct {
	Repo   *repository.SubscriptionRepository
	Valid  *validator.Validate
	Log    *lib.Logger
	Config *config.Config
}
