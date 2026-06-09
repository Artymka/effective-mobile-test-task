package database

import (
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/lib/pq"
)

func (r *SubscriptionRepositoryPostgres) Create(sub *models.Subscription) error {
	query := `
        INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at
    `

	err := r.db.QueryRowx(
		query,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate,
		sub.EndDate,
		// sub.CreatedAt,
	// ).Scan(&sub.ID, &sub.CreatedAt)
	).Scan(&sub.ID, &sub.CreatedAt)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return NotUniqueErr
			}
		}
		return err
	}
	return nil
}
