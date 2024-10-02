package service

import (
	"context"
	"timetracker/internal/db"
	"timetracker/internal/repository"
)

type ProjectService interface {
	CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error)
	DeleteProject(ctx context.Context, id int64) error
	GetProject(ctx context.Context, id int64) (db.Project, error)
	ListProjects(ctx context.Context) ([]db.Project, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error) {
	return s.repo.CreateProject(ctx, params)
}

func (s *projectService) DeleteProject(ctx context.Context, id int64) error {
	return s.repo.DeleteProject(ctx, id)
}

func (s *projectService) GetProject(ctx context.Context, id int64) (db.Project, error) {
	return s.repo.GetProject(ctx, id)
}

func (s *projectService) ListProjects(ctx context.Context) ([]db.Project, error) {
	return s.repo.ListProjects(ctx)
}
