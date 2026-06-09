package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ServiceName string     `db:"service_name" json:"service_name"`
	Price       int        `db:"price" json:"price"`
	UserID      uuid.UUID  `db:"user_id" json:"user_id"`
	StartDate   time.Time  `db:"start_date" json:"start_date"`
	EndDate     *time.Time `db:"end_date" json:"end_date,omitempty"`
	CreatedAt   time.Time  `db:"created_at"`
}

type CreateSubscriptionRequest struct {
	ServiceName string  `json:"service_name" validate:"required,gte=1"`
	Price       *int    `json:"price" validate:"required,gte=0"`
	UserID      string  `json:"user_id" validate:"required,uuid"`
	StartDate   string  `json:"start_date" validate:"required,datetime=01-2006"`
	EndDate     *string `json:"end_date,omitempty" validate:"omitempty,datetime=01-2006"`
}

type GetSubscriptionRequest struct {
	SubscriptionID string `validate:"required,uuid"`
}

type DeleteSubscriptionRequest struct {
	SubscriptionID string `validate:"required,uuid"`
}

type UpdateSubscriptionRequest struct {
	ID          string  `json:"id" validate:"required,uuid"`
	ServiceName *string `json:"service_name" validate:"omitempty,gte=1"`
	Price       *int    `json:"price" validate:"omitempty,gte=0"`
	StartDate   *string `json:"start_date" validate:"omitempty,datetime=01-2006"`
	EndDate     *string `json:"end_date" validate:"omitempty,datetime=01-2006"`
}

func (s *UpdateSubscriptionRequest) ToDB() UpdateSubscription {
	var startDate, endDate *time.Time
	if s.StartDate != nil {
		t, _ := time.Parse("01-2006", *s.StartDate)
		startDate = &t
	}
	if s.EndDate != nil {
		t, _ := time.Parse("01-2006", *s.EndDate)
		endDate = &t
	}

	return UpdateSubscription{
		ServiceName: s.ServiceName,
		Price:       s.Price,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}

type UpdateSubscription struct {
	ServiceName *string    `db:"service_name"`
	Price       *int       `db:"price"`
	UserID      *uuid.UUID `db:"user_id"`
	StartDate   *time.Time `db:"start_date"`
	EndDate     *time.Time `db:"end_date"`
}

type SubscriptionResponse struct {
	ID          string  `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	CreatedAt   string  `json:"created_at"`
}

type SubscriptionListResponse struct {
	Subscriptions []SubscriptionResponse `json:"subscriptions"`
	PrevPage      string                 `json:"prev_page,omitempty"`
	NextPage      string                 `json:"next_page,omitempty"`
	Total         int                    `json:"total"`
	Page          int                    `json:"page"`
	Limit         int                    `json:"limit"`
}

type TotalCostRequest struct {
	StartDate   string  `json:"start_date" validate:"required,datetime=01-2006"`
	EndDate     string  `json:"end_date" validate:"required,datetime=01-2006"`
	UserID      *string `json:"user_id" validate:"omitempty,uuid"`
	ServiceName *string `json:"service_name" validate:"omitempty,gte=1"`
}

type TotalCostFilter struct {
	StartDate   time.Time
	EndDate     time.Time
	UserID      *uuid.UUID
	ServiceName *string
}

type AggregateResponse struct {
	TotalCost int `json:"total_cost"`
}

func (s *Subscription) ToResponse() SubscriptionResponse {
	var endDate *string
	if s.EndDate != nil {
		t := s.EndDate.Format("01-2006")
		endDate = &t
	}

	return SubscriptionResponse{
		ID:          s.ID.String(),
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID.String(),
		StartDate:   s.StartDate.Format("01-2006"),
		EndDate:     endDate,
		CreatedAt:   s.CreatedAt.Format(time.RFC3339),
	}
}
