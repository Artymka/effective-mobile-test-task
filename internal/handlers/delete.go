package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/google/uuid"
)

// @Summary      Delete subscription by id
// @Produce      json
// @Param        id query string true "Subscription id"
// @Success      200  {object} lib.Response
// @Failure      400,404,500  {object} lib.ErrResponse
// @Router       /subscriptions [delete]
func (h *SubscriptionHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	const op = "handler.delete"

	// validation
	reqData := models.DeleteSubscriptionRequest{}
	reqData.SubscriptionID = r.URL.Query().Get("id")

	if err, msg := lib.ValidateStruct(h.Valid, &reqData, "Wrong query params: "); err != nil {
		h.Log.Debug(op, err, "")
		lib.WriteError(w, msg, http.StatusBadRequest)
		return
	}

	// db logic
	subscriptionID, _ := uuid.Parse(reqData.SubscriptionID)
	err := h.Repo.Delete(subscriptionID)

	// response
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			lib.WriteError(w, "Subscription not found", http.StatusNotFound)
		} else {
			h.Log.Error(op, err)
			lib.WriteError(w, "Error while deleting subscription", http.StatusInternalServerError)
		}
	} else {
		lib.WriteResponse(w, nil)
	}
}
