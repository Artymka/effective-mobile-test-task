package repository

import (
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *SubscriptionRepository) Update(id uuid.UUID, data models.UpdateSubscription) (models.Subscription, error) {
	query := `
		UPDATE subscriptions SET 
        service_name = COALESCE($1, service_name),
        price = COALESCE($2, price),
        start_date = COALESCE($3, start_date),
        end_date = COALESCE($4, end_date)
        WHERE id = $5
		RETURNING id, service_name, user_id, price, start_date, end_date, created_at`

	var updatedSub models.Subscription
	err := r.db.Get(&updatedSub, query, data.ServiceName, data.Price, data.StartDate, data.EndDate, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return updatedSub, NotUniqueErr
			}
		}
		return updatedSub, err
	}

	return updatedSub, nil
	// rows, err := result.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	// if rows == 0 {
	// 	return sql.ErrNoRows
	// }

	// return nil
}
