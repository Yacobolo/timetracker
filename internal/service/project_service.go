package service

import (
	"context"
	"timetracker/internal/db"
	queries "timetracker/internal/db/queries/dynamic"
	"timetracker/internal/dto"
	"timetracker/internal/repository"
)

type ProjectService interface {
	CreateProject(ctx context.Context, input dto.ProjectIn) (dto.ProjectOut, error)
	DeleteProject(ctx context.Context, id int32) error
	GetProject(ctx context.Context, id int32) (db.Project, error)
	ListProjects(ctx context.Context, sortBy string, sortOrder string) ([]db.Project, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
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

func (s *projectService) DeleteProject(ctx context.Context, id int32) error {
	return s.repo.DeleteProject(ctx, id)
}

func (s *projectService) GetProject(ctx context.Context, id int32) (db.Project, error) {
	return s.repo.GetProject(ctx, id)
}

func (s *projectService) ListProjects(ctx context.Context, sortBy string, sortOrder string) ([]db.Project, error) {

	opts := queries.ProjectListQueryOpts{
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Limit:     0,
		Offset:    0,
	}

	projectList, err := s.repo.ListProjects(ctx, opts)
	if err != nil {
		return nil, err
	}

	return projectList, nil
}
