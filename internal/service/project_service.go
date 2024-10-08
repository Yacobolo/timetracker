package service

import (
	"context"
	"timetracker/internal/db"
	"timetracker/internal/dto"
	"timetracker/internal/repository"
)

type ProjectService interface {
	CreateProject(ctx context.Context, input dto.ProjectIn) (dto.ProjectOut, error)
	DeleteProject(ctx context.Context, id int64) error
	GetProject(ctx context.Context, id int64) (db.Project, error)
	ListProjects(ctx context.Context) ([]dto.ProjectOut, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

type ProjectInput struct {
	Name        string
	Description string
}

func (s *projectService) CreateProject(ctx context.Context, input dto.ProjectIn) (dto.ProjectOut, error) {

	createParams := db.CreateProjectParams{
		Name:        input.Name,
		Description: input.Description,
	}
	project, err := s.repo.CreateProject(ctx, createParams)

	if err != nil {
		return dto.ProjectOut{}, err
	}

	// Convert the project to
	projectDTO := dto.ToProjectOutDTO(project)

	return projectDTO, nil
}

func (s *projectService) DeleteProject(ctx context.Context, id int64) error {
	return s.repo.DeleteProject(ctx, id)
}

func (s *projectService) GetProject(ctx context.Context, id int64) (db.Project, error) {
	return s.repo.GetProject(ctx, id)
}

func (s *projectService) ListProjects(ctx context.Context) ([]dto.ProjectOut, error) {
	projectList, err := s.repo.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	// Pre-allocate the response slice to avoid multiple allocations
	output := make([]dto.ProjectOut, len(projectList))

	// Loop over the projectList and convert each project to a CreateProjectResponse
	for i, project := range projectList {
		output[i] = dto.ToProjectOutDTO(project)
	}

	return output, nil
}
