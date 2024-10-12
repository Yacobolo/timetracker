package handler

import (
	"net/http"
	"timetracker/internal/service"
	"timetracker/internal/templates/pages"
)

type TimeEntryHandler struct {
	service service.TimeEntryService
}

func NewTimeEntryHandler(service service.TimeEntryService) TimeEntryHandler {
	return TimeEntryHandler{service: service}
}

func (h *TimeEntryHandler) RenderTimeEntryIndex(w http.ResponseWriter, r *http.Request) error {
	return pages.TimerPage().Render(r.Context(), w)
}
