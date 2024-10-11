package repository

import (
	"context"
	"timetracker/internal/db"
	queries "timetracker/internal/db/queries/dynamic"

	"github.com/jmoiron/sqlx"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error)
	DeleteProject(ctx context.Context, id int64) error
	GetProject(ctx context.Context, id int64) (db.Project, error)
	ListProjects(ctx context.Context, params queries.ProjectListQueryOpts) ([]db.Project, error)
}

type projectRepository struct {
	db      *sqlx.DB
	queries *db.Queries
}

func NewProjectRepository(db *sqlx.DB, queries *db.Queries) ProjectRepository {
	return &projectRepository{db: db, queries: queries}
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

func (r *projectRepository) ListProjects(ctx context.Context, opts queries.ProjectListQueryOpts) ([]db.Project, error) {
	// Build dynamic query
	query := queries.BuildProjectListQuery(opts)

	// Convert to SQL and arguments
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	// Execute the query and map results to the projects slice
	var projects []db.Project
	err = r.db.SelectContext(ctx, &projects, sql, args...)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
