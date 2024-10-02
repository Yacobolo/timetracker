package repository

import (
	"context"
	"timetracker/internal/db"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error)
	DeleteProject(ctx context.Context, id int64) error
	GetProject(ctx context.Context, id int64) (db.Project, error)
	ListProjects(ctx context.Context) ([]db.Project, error)
}

type projectRepository struct {
	queries *db.Queries
}

func NewProjectRepository(queries *db.Queries) ProjectRepository {
	return &projectRepository{queries: queries}
}

func (r *projectRepository) CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error) {
	return r.queries.CreateProject(ctx, params)
}

func (r *projectRepository) DeleteProject(ctx context.Context, id int64) error {
	return r.queries.DeleteProject(ctx, id)
}

func (r *projectRepository) GetProject(ctx context.Context, id int64) (db.Project, error) {
	return r.queries.GetProject(ctx, id)
}

func (r *projectRepository) ListProjects(ctx context.Context) ([]db.Project, error) {
	return r.queries.ListProjects(ctx)
}
