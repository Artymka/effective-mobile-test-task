package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/google/uuid"
)

// @Summary      Сalculate the total cost of subscriptions for a specified period with the ability to filter by user or service
// @Produce      json
// @Param        start_date query string true "Subscription start date"
// @Param        end_date query string true "Subscription end date"
// @Param        service_name query string false "Subscription service name"
// @Param        user_id query string false "Subscription user id"
// @Success      200  {object} lib.Response{data=integer}
// @Failure      400,500  {object} lib.ErrResponse
// @Router       /subscriptions/total-cost [get]
func (h *SubscriptionHandlers) TotalCost(w http.ResponseWriter, r *http.Request) {
	const op = "handler.total_cost"

	// validation
	var userID, serviceName *string
	if t := r.URL.Query().Get("user_id"); t != "" {
		userID = &t
	}
	if t := r.URL.Query().Get("service_name"); t != "" {
		serviceName = &t
	}

	reqData := models.TotalCostRequest{
		StartDate:   r.URL.Query().Get("start_date"),
		EndDate:     r.URL.Query().Get("end_date"),
		UserID:      userID,
		ServiceName: serviceName,
	}

	if err, msg := lib.ValidateStruct(h.Valid, &reqData, "Wrong params: "); err != nil {
		h.Log.Debug(op, err, "")
		lib.WriteError(w, msg, http.StatusBadRequest)
		return
	}

	// db logic
	var procUserID *uuid.UUID
	if reqData.UserID != nil {
		t, _ := uuid.Parse(*reqData.UserID)
		procUserID = &t
	}

	startDate, _ := time.Parse("01-2006", reqData.StartDate)
	endDate, _ := time.Parse("01-2006", reqData.EndDate)

	totalCost, err := h.Repo.TotalCost(models.TotalCostFilter{
		StartDate:   startDate,
		EndDate:     endDate,
		UserID:      procUserID,
		ServiceName: reqData.ServiceName,
	})

	if err != nil {
		if errors.Is(err, repository.NotFoundErr) {
			lib.WriteError(w, "Wrong user_id or service_name", http.StatusBadRequest)
		} else {
			h.Log.Error(op, err)
			lib.WriteError(w, "Error while calculating total cost", http.StatusInternalServerError)
		}
		return
	}

	lib.WriteResponse(w, totalCost)
}
