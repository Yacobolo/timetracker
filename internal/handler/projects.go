package handler

import (
	"log"
	"net/http"
	"timetracker/internal/dto"
	"timetracker/internal/service"
	"timetracker/internal/templates/components"
	"timetracker/pkg/table"

	"github.com/go-playground/validator/v10"
)

type ProjectHandler struct {
	service   service.ProjectService
	validator *validator.Validate
}

func NewProjectHandler(service service.ProjectService, validator *validator.Validate) *ProjectHandler {
	return &ProjectHandler{service: service, validator: validator}
}

func (h *ProjectHandler) RenderProjectList(w http.ResponseWriter, r *http.Request) error {
	projects, err := h.service.ListProjects(r.Context())

	if err != nil {
		return err
	}

	projects_table, err := table.NewTableFromStructs(projects)
	if err != nil {
		return err
	}

	return components.Table(projects_table).Render(r.Context(), w)
}

func (h *ProjectHandler) HandleProjectSubmit(w http.ResponseWriter, r *http.Request) error {
	// Parse the form values
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return err
	}

	// log form values
	log.Printf("Form values: %v", r.PostForm)

	input := dto.ProjectIn{
		Name:        r.PostFormValue("Project Name"),
		Description: r.PostFormValue("Description"),
	}

	// Validate the form input
	if err := h.validator.Struct(input); err != nil {
		log.Printf("Validation failed: %v", err)
		return AddNotificationHeaders(w, "Faild to create project", "error")
	}

	project, err := h.service.CreateProject(r.Context(), input)
	if err != nil {
		return err
	}

	row, err := table.NewRowFromStruct(project)

	if err != nil {
		return err
	}

	// Add success notification
	if err := AddNotificationHeaders(w, "Created project successfully", "success"); err != nil {
		return err
	}

	return components.Row(row.Values).Render(r.Context(), w)

}
