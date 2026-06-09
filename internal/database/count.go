package database

func (r *SubscriptionRepositoryPostgres) Count() int {
	var total int
	query := `SELECT COUNT(1) FROM subscriptions`
	if err := r.db.Get(&total, query); err != nil {
		return 0
	}
	return total
}
