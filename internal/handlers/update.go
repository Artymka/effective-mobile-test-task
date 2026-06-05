package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/models"
	"github.com/google/uuid"
)

func (h *SubscriptionHandlers) Update(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.update"

	// request data
	reqData := models.UpdateSubscriptionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		lib.WriteError(w, "Wrong json scheme", http.StatusBadRequest)
		return
	}

	reqData.ID = r.URL.Query().Get("id")

	// validation
	if err, msg := lib.ValidateStruct(h.Valid, &reqData, "Wrong query params or json: "); err != nil {
		h.Log.Debug(op, err, "")
		lib.WriteError(w, msg, http.StatusBadRequest)
		return
	}

	// db logic
	updateSubs := reqData.ToDB()
	id, _ := uuid.Parse(reqData.ID)
	err := h.Repo.Update(id, updateSubs)

	// response
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			lib.WriteError(w, "Subscription not found", http.StatusNotFound)
		} else {
			h.Log.Error(op, err)
			lib.WriteError(w, "Error while updating subscription", http.StatusInternalServerError)
		}
	} else {
		lib.WriteResponse(w, nil)
	}
}
