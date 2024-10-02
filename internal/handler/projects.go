package handler

import (
	"database/sql"
	"log"
	"net/http"
	"timetracker/internal/db"
	"timetracker/internal/service"
	"timetracker/internal/templates/components"
)

type ProjectHandler struct {
	service service.ProjectService
}

func NewProjectHandler(service service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) RenderProjectList(w http.ResponseWriter, r *http.Request) error {
	projects, err := h.service.ListProjects(r.Context())

	if err != nil {
		return err
	}

	return components.Table(projects).Render(r.Context(), w)
}

func (h *ProjectHandler) HandleProjectSubmit(w http.ResponseWriter, r *http.Request) error {
	// log form values
	log.Printf("Form values: %v", r.PostForm)

	params := db.CreateProjectParams{
		Name:        r.PostFormValue("Project Name"),
		Description: sql.NullString{String: r.PostFormValue("Description"), Valid: true},
	}

	project, err := h.service.CreateProject(r.Context(), params)
	if err == nil {
		log.Printf("Created project: %v", project)
	}
	return err

}
