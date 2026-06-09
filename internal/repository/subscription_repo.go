package repository

import (
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/google/uuid"
)

type SubscriptionRepository interface {
	Count() int
	Create(sub *models.Subscription) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Subscription, error)
	List(page, limit int) ([]models.Subscription, error)
	TotalCost(data models.TotalCostFilter) (int, error)
	Update(id uuid.UUID, data models.UpdateSubscription) (models.Subscription, error)
}
