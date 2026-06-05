package handlers

import (
	"net/http"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/google/uuid"
)

func (h *SubscriptionHandlers) Get(w http.ResponseWriter, r *http.Request) {
	const op = "handler.get"

	// validation
	reqData := models.GetSubscriptionRequest{}
	reqData.SubscriptionID = r.URL.Query().Get("id")

	if err, msg := lib.ValidateStruct(h.Valid, &reqData, "Wrong query params: "); err != nil {
		h.Log.Debug(op, err, "")
		lib.WriteError(w, msg, http.StatusBadRequest)
		return
	}

	// db logic
	subscriptionID, _ := uuid.Parse(reqData.SubscriptionID)

	subscription, err := h.Repo.GetByID(subscriptionID)
	if err != nil {
		h.Log.Error(op, err)
		lib.WriteError(w, "Error while getting subscription", http.StatusInternalServerError)
		return
	}

	// response
	if subscription == nil {
		lib.WriteError(w, "Subscription not found", http.StatusNotFound)
		return
	}

	respData := subscription.ToResponse()
	lib.WriteResponse(w, &respData)
}
