package database

import (
	"errors"

	"github.com/Artymka/effective-mobile-test-task/internal/config"
	"github.com/jmoiron/sqlx"
)

var (
	NotFoundErr  = errors.New("Record not found")
	NotUniqueErr = errors.New("Unique constraint violated")
)

type SubscriptionRepositoryPostgres struct {
	db     *sqlx.DB
	config *config.Config
}

func NewSubscriptionRepositoryPostgres(db *sqlx.DB, conf *config.Config) *SubscriptionRepositoryPostgres {
	return &SubscriptionRepositoryPostgres{db: db}
}
