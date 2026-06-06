package repository

import (
	"errors"

	"github.com/Artymka/effective-mobile-test-task/internal/config"
	"github.com/jmoiron/sqlx"
)

var (
	NotFoundErr = errors.New("Record not found")
)

type SubscriptionRepository struct {
	db     *sqlx.DB
	config *config.Config
}

func NewSubscriptionRepository(db *sqlx.DB, conf *config.Config) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}
