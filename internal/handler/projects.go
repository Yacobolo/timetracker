package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"timetracker/internal/config"
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

func (h *ProjectHandler) RenderProjectForm(w http.ResponseWriter, r *http.Request) error {
	input_fields := BuildInputFields(*config.ProjectFieldConfigManager, []FieldError{})
	return components.ModalForm(input_fields).Render(r.Context(), w)
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
		Name:        r.PostFormValue("name"),
		Description: r.PostFormValue("description"),
	}

	// Validate the form input
	if err := h.validator.Struct(input); err != nil {

		fmt.Println("Validation error")
		var field_errors []FieldError // Correctly declare the slice
		for _, err := range err.(validator.ValidationErrors) {
			field_id := strings.ToLower(err.Field())
			errorMsg := CustomErrorMessage(err)
			field_errors = append(field_errors, FieldError{FieldID: field_id, Error: errorMsg})

		}

		input_fields := BuildInputFields(*config.ProjectFieldConfigManager, field_errors)
		return components.ModalForm(input_fields).Render(r.Context(), w)
	}

	// Create the project
	project, err := h.service.CreateProject(r.Context(), input)
	if err != nil {

		if handled := handleSQLiteError(err, config.ProjectFieldConfigManager); handled != nil {
			input_fields := BuildInputFields(*config.ProjectFieldConfigManager, handled)
			return components.ModalForm(input_fields).Render(r.Context(), w)

		}
		return err
	}

	row, err := table.NewRowFromStruct(project)

	if err != nil {
		return err
	}

	// Add success notification
	if err := AddHxNotificationTrigger(w, "Created project successfully", "success"); err != nil {
		return err
	}

	if err := AddHxTrigger(w, "close-modal", nil); err != nil {
		return err
	}

	return components.HxRow(row.Values).Render(r.Context(), w)

}
