package repository

import (
	"database/sql"

	"github.com/Artymka/effective-mobile-test-task/internal/models"

	"github.com/google/uuid"
)

func (r *SubscriptionRepository) GetByID(id uuid.UUID) (*models.Subscription, error) {
	query := `SELECT * FROM subscriptions WHERE id = $1`

	var sub models.Subscription
	err := r.db.Get(&sub, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &sub, nil
}
