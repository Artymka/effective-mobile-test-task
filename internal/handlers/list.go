package handlers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/Artymka/effective-mobile-test-task/internal/lib"
	"github.com/Artymka/effective-mobile-test-task/internal/models"
)

func (h *SubscriptionHandlers) List(w http.ResponseWriter, r *http.Request) {
	const op = "handler.list"

	// validation
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// db logic
	// count validation
	totalCount := h.Repo.Count()
	if totalCount <= limit*(page-1) {
		lib.WriteError(w, "No such page found", http.StatusNotFound)
	}

	// previous and next pages
	var prevURLs, nextURLs string
	if page > 1 {
		params := url.Values{}
		params.Add("page", strconv.Itoa(page-1))
		params.Add("limit", strconv.Itoa(limit))
		prevURL, _ := url.Parse(r.URL.String())
		prevURLs = prevURL.String()
	}
	if totalCount > limit*page {
		params := url.Values{}
		params.Add("page", strconv.Itoa(page+1))
		params.Add("limit", strconv.Itoa(limit))
		nextURL, _ := url.Parse(r.URL.String())
		nextURLs = nextURL.String()
	}

	subs, err := h.Repo.List(page, limit)
	if err != nil {
		lib.WriteError(w, "Error while extracting list of subscriptions", http.StatusInternalServerError)
		return
	}

	// response
	subsResponse := make([]models.SubscriptionResponse, len(subs))
	for i, sub := range subs {
		subsResponse[i] = sub.ToResponse()
	}

	lib.WriteResponse(w, &models.SubscriptionListResponse{
		Subscriptions: subsResponse,
		PrevPage:      prevURLs,
		NextPage:      nextURLs,
		Total:         totalCount,
		Page:          page,
		Limit:         limit,
	})
}
