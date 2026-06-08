package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/Artymka/effective-mobile-test-task/internal/repository"
	"github.com/google/uuid"
)

// @Summary      Create subscription
// @Accept       json
// @Produce      json
// @Param        request body models.CreateSubscriptionRequest true "Subscription info"
// @Success      201  {object} lib.Response{data=models.SubscriptionResponse}
// @Failure      400,409,500  {object} lib.ErrResponse
// @Router       /subscriptions [post]
func (h *SubscriptionHandlers) Create(w http.ResponseWriter, r *http.Request) {
	const op = "handler.create"

	// validation
	var reqData models.CreateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		h.Log.Debug(op, err, "")
		lib.WriteError(w, "Wrong json scheme", http.StatusBadRequest)
		return
	}

	if err, msg := lib.ValidateStruct(h.Valid, &reqData, "Wrong json: "); err != nil {
		h.Log.Debug(op, err, "")
		lib.WriteError(w, msg, http.StatusBadRequest)
		return
	}

	// // if err := reqData.Validate(); err != nil {
	// if errs := h.Valid.Struct(&reqData); errs != nil {
	// 	h.Log.Debug(op, errs, "")

	// 	errSb := strings.Builder{}
	// 	errSb.WriteString("Wrong json: ")
	// 	for _, err := range errs.(validator.ValidationErrors) {
	// 		errSb.WriteString(fmt.Sprintf("%s - %s, ", err.Field(), err.Tag()))
	// 	}

	// 	lib.WriteError(w, errSb.String(), http.StatusBadRequest)
	// 	return
	// }

	// db logic
	startDate, _ := time.Parse("01-2006", reqData.StartDate)

	var endDate *time.Time
	if reqData.EndDate != nil {
		parsed, _ := time.Parse("01-2006", *reqData.EndDate)
		endDate = &parsed
	}

	userID, _ := uuid.Parse(reqData.UserID)

	subscription := &models.Subscription{
		ServiceName: reqData.ServiceName,
		Price:       *reqData.Price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := h.Repo.Create(subscription); err != nil {
		if errors.Is(err, repository.NotUniqueErr) {
			lib.WriteError(w, "Group of service, user and start date must be unique", http.StatusConflict)
		} else {
			h.Log.Error(op, err)
			lib.WriteError(w, "Error while creating subscription", http.StatusInternalServerError)
		}
		return
	}

	// response
	responseData := subscription.ToResponse()
	lib.WriteResponse(w, &responseData)
	w.WriteHeader(http.StatusCreated)
}
