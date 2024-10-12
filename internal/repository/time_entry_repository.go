package repository

import (
	"context"
	"timetracker/internal/db"

	"github.com/jmoiron/sqlx"
)

type TimeEntryRepository interface {
	CreateTimeEntry(ctx context.Context, params db.CreateTimeEntryParams) (db.TimeEntry, error)
	DeleteTimeEntry(ctx context.Context, id int64) error
	GetTimeEntry(ctx context.Context, id int64) (db.TimeEntry, error)
	ListTimeEntries(ctx context.Context) ([]db.TimeEntry, error)
}

type timeEntryRepository struct {
	db      *sqlx.DB
	queries *db.Queries
}

func NewTimeEntryRepository(db *sqlx.DB, queries *db.Queries) TimeEntryRepository {
	return &timeEntryRepository{db: db, queries: queries}
}

func (r *timeEntryRepository) CreateTimeEntry(ctx context.Context, params db.CreateTimeEntryParams) (db.TimeEntry, error) {
	return r.queries.CreateTimeEntry(ctx, params)
}

func (r *timeEntryRepository) DeleteTimeEntry(ctx context.Context, id int64) error {
	return r.queries.DeleteTimeEntry(ctx, id)
}

func (r *timeEntryRepository) GetTimeEntry(ctx context.Context, id int64) (db.TimeEntry, error) {
	return r.queries.GetTimeEntry(ctx, id)
}

func (r *timeEntryRepository) ListTimeEntries(ctx context.Context) ([]db.TimeEntry, error) {
	return r.queries.ListTimeEntries(ctx)
}
