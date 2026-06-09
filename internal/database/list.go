package database

import "github.com/Artymka/effective-mobile-test-task/internal/models"

func (r *SubscriptionRepositoryPostgres) List(page, limit int) ([]models.Subscription, error) {
	offset := (page - 1) * limit

	// Get paginated results
	query := `
        SELECT * FROM subscriptions 
        ORDER BY created_at DESC 
        LIMIT $1 OFFSET $2
    `

	var subscriptions []models.Subscription
	err := r.db.Select(&subscriptions, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}
