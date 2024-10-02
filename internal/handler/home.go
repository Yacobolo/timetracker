package handler

import (
	"net/http"
	"timetracker/internal/templates/errors"
)

func RenderHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return errors.Error404().Render(r.Context(), w)
}
