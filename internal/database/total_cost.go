package database

import (
	"fmt"

	"github.com/Artymka/effective-mobile-test-task/internal/models"
)

func (r *SubscriptionRepositoryPostgres) TotalCost(data models.TotalCostFilter) (int, error) {
	query := `
        SELECT
		SUM(
			(EXTRACT(YEAR FROM LEAST(COALESCE(end_date, CURRENT_DATE), $2)) * 12 +
			EXTRACT(MONTH FROM LEAST(COALESCE(end_date, CURRENT_DATE), $2)) -
			EXTRACT(YEAR FROM GREATEST(start_date, $1)) * 12 -
			EXTRACT(MONTH FROM GREATEST(start_date, $1))) * price)
		AS total_cost
		FROM subscriptions
		WHERE 1=1
    `

	args := []interface{}{data.StartDate, data.EndDate}
	paramCount := 2

	if data.UserID != nil {
		paramCount++
		query += fmt.Sprintf(" AND user_id = $%d", paramCount)
		args = append(args, *data.UserID)
	}

	if data.ServiceName != nil {
		paramCount++
		query += fmt.Sprintf(" AND service_name = $%d", paramCount)
		args = append(args, *data.ServiceName)
	}

	var totalCost *int
	err := r.db.Get(&totalCost, query, args...)
	if err != nil {
		return 0, err
	}

	if totalCost == nil {
		return 0, NotFoundErr
	}

	return *totalCost, nil
}
