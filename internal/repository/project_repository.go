package repository

import (
	"context"
	"timetracker/internal/db"
	queries "timetracker/internal/db/queries/dynamic"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error)
	DeleteProject(ctx context.Context, id int32) error
	GetProject(ctx context.Context, id int32) (db.Project, error)
	ListProjects(ctx context.Context, params queries.ProjectListQueryOpts) ([]db.Project, error)
}

type projectRepository struct {
	db      *pgxpool.Pool
	queries *db.Queries
}

func NewProjectRepository(db *pgxpool.Pool, queries *db.Queries) ProjectRepository {
	return &projectRepository{db: db, queries: queries}
}

func (r *projectRepository) CreateProject(ctx context.Context, params db.CreateProjectParams) (db.Project, error) {
	return r.queries.CreateProject(ctx, params)
}

func (r *projectRepository) DeleteProject(ctx context.Context, id int32) error {
	return r.queries.DeleteProject(ctx, id)
}

func (r *projectRepository) GetProject(ctx context.Context, id int32) (db.Project, error) {
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
	pgxscan.Select(ctx, r.db, &projects, sql, args...)

	return projects, nil
}
