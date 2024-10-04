package dto

import (
	"strconv"
	"timetracker/internal/db"
)

// convert dto to db model

type ProjectIn struct {
	Name        string `form:"name" validate:"required"`
	Description string `form:"description"`
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
		ID:          strconv.FormatInt(model.ID, 10),
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   model.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
