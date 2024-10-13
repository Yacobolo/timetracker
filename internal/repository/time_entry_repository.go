package repository

import (
	"context"
	"timetracker/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TimeEntryRepository interface {
	CreateTimeEntry(ctx context.Context, params db.CreateTimeEntryParams) (db.TimeEntry, error)
	DeleteTimeEntry(ctx context.Context, id int32) error
	GetTimeEntry(ctx context.Context, id int32) (db.TimeEntry, error)
	ListTimeEntries(ctx context.Context) ([]db.TimeEntry, error)
}

type timeEntryRepository struct {
	db      *pgxpool.Pool
	queries *db.Queries
}

func NewTimeEntryRepository(db *pgxpool.Pool, queries *db.Queries) TimeEntryRepository {
	return &timeEntryRepository{db: db, queries: queries}
}

func (r *timeEntryRepository) CreateTimeEntry(ctx context.Context, params db.CreateTimeEntryParams) (db.TimeEntry, error) {
	return r.queries.CreateTimeEntry(ctx, params)
}

func (r *timeEntryRepository) DeleteTimeEntry(ctx context.Context, id int32) error {
	return r.queries.DeleteTimeEntry(ctx, id)
}

func (r *timeEntryRepository) GetTimeEntry(ctx context.Context, id int32) (db.TimeEntry, error) {
	return r.queries.GetTimeEntry(ctx, id)
}

func (r *timeEntryRepository) ListTimeEntries(ctx context.Context) ([]db.TimeEntry, error) {
	return r.queries.ListTimeEntries(ctx)
}
