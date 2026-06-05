package repository

import (
	"github.com/Artymka/effective-mobile-test-task/internal/models"
)

/*
нужно расчитать произведение месячной стоимости на длительность подписки
Причем длительность подписки должна вычисляться как Min(record.EndDate, param.EndDate) - startDate

У меня есть таблица с подписками, записи которой содержат id, id пользователя, название сервиса, месячную стоимость подписки, дату начала подписки и дату окончания подписки/NULL.
У меня есть обязательные параметры для фильтрации: дата начала и дата окончания.
Также у меня есть два необязвательных параметра: id пользователя и названия сервиса.
Мне нужно посчитать суммарную стоимость подписок с настренными фильтрами.
Как это сделать?

Создать новую таблицу, в которой будут количества месяцев для каждой записи (или для фильтрованных записей)
После
*/

func (r *SubscriptionRepository) TotalCost(data models.TotalCostFilter) (int, error) {
	query := `
        SELECT
		SUM(
			(EXTRACT(YEAR FROM LEAST(COALESCE(end_date, CURRENT_DATE), $2)) * 12 +
			EXTRACT(MONTH FROM LEAST(COALESCE(end_date, CURRENT_DATE), $2)) -
			EXTRACT(YEAR FROM GREATEST(start_date, $1)) * 12 -
			EXTRACT(MONTH FROM GREATEST(start_date, $1))) * price)
		AS total_cost
		FROM subscriptions
		WHERE
		($3 IS NULL OR user_id = $3) AND
		($4 IS NULL OR service_name = $4);
    `
	// args := []interface{}{startDate, endDate}
	// argIndex := 3

	// if userID != "" {
	// 	userUUID, err := uuid.Parse(userID)
	// 	if err != nil {
	// 		return 0, fmt.Errorf("invalid user_id format: %w", err)
	// 	}
	// 	query += fmt.Sprintf(" AND user_id = $%d", argIndex)
	// 	args = append(args, userUUID)
	// 	argIndex++
	// }

	// if serviceName != "" {
	// 	query += fmt.Sprintf(" AND service_name = $%d", argIndex)
	// 	args = append(args, serviceName)
	// }

	var totalCost int
	err := r.db.Get(&totalCost, query, data.StartDate, data.EndDate, data.UserID, data.ServiceName)
	if err != nil {
		return 0, err
	}

	return totalCost, nil
}
