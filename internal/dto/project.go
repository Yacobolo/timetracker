package dto

import (
	"fmt"
	"timetracker/internal/db"
)

// convert dto to db model

type ProjectIn struct {
	Name        string `validate:"required"`
	Description string
}

type ProjectOut struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

func ProjectInToDB(model ProjectIn) db.CreateProjectParams {
	return db.CreateProjectParams{
		Name:        model.Name,
		Description: model.Description,
	}
}

func ToProjectOutDTO(model db.Project) ProjectOut {
	return ProjectOut{
		ID:          fmt.Sprintf("%d", model.ID),
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:   model.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}

type TimeEntryIn struct {
	ID          string `validate:"required"`
	ProjectID   string
	StartTime   string `validate:"required"`
	EndTime     string `validate:"required"`
	Description string `validate:"required"`
}
