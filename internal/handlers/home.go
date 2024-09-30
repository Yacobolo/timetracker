package handlers

import (
	"net/http"
	"timetracker/internal/templates/errors"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return errors.Error404().Render(r.Context(), w)
}
