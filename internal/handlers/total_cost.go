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
			lib.WriteError(w, "Wrong user_id or service_name", http.StatusConflict)
		} else {
			h.Log.Error(op, err)
			lib.WriteError(w, "Error while calculating total cost", http.StatusInternalServerError)
		}
		return
	}

	lib.WriteResponse(w, totalCost)
}
