package repository

import (
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/google/uuid"
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

	// var startDate interface{}
	// if req.StartDate != nil {
	// 	parsed, err := time.Parse("01-2006", *req.StartDate)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	startDate = parsed
	// } else {
	// 	startDate = nil
	// }

	// var endDate interface{}
	// if req.EndDate != nil {
	// 	if *req.EndDate == "" {
	// 		endDate = nil
	// 	} else {
	// 		parsed, err := time.Parse("01-2006", *req.EndDate)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		endDate = parsed
	// 	}
	// } else {
	// 	endDate = nil
	// }

	var updatedSub models.Subscription
	err := r.db.Get(&updatedSub, query, data.ServiceName, data.Price, data.StartDate, data.EndDate, id)
	if err != nil {
		return updatedSub, err
	}

	return updatedSub, err
	// rows, err := result.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	// if rows == 0 {
	// 	return sql.ErrNoRows
	// }

	// return nil
}
