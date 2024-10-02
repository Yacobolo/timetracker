// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: time_entries.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTimeEntry = `-- name: CreateTimeEntry :one
INSERT INTO time_entries (
  project_id, start_time, end_time, description
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, project_id, start_time, end_time, duration, description, created_at, updated_at
`

type CreateTimeEntryParams struct {
	ProjectID   int64
	StartTime   time.Time
	EndTime     time.Time
	Description sql.NullString
}

func (q *Queries) CreateTimeEntry(ctx context.Context, arg CreateTimeEntryParams) (TimeEntry, error) {
	row := q.db.QueryRowContext(ctx, createTimeEntry,
		arg.ProjectID,
		arg.StartTime,
		arg.EndTime,
		arg.Description,
	)
	var i TimeEntry
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.StartTime,
		&i.EndTime,
		&i.Duration,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTimeEntry = `-- name: DeleteTimeEntry :exec
DELETE FROM time_entries
WHERE id = ?
`

func (q *Queries) DeleteTimeEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTimeEntry, id)
	return err
}

const getTimeEntry = `-- name: GetTimeEntry :one
SELECT id, project_id, start_time, end_time, duration, description, created_at, updated_at FROM time_entries
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTimeEntry(ctx context.Context, id int64) (TimeEntry, error) {
	row := q.db.QueryRowContext(ctx, getTimeEntry, id)
	var i TimeEntry
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.StartTime,
		&i.EndTime,
		&i.Duration,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTimeEntries = `-- name: ListTimeEntries :many
SELECT id, project_id, start_time, end_time, duration, description, created_at, updated_at FROM time_entries
ORDER BY start_time
`

func (q *Queries) ListTimeEntries(ctx context.Context) ([]TimeEntry, error) {
	rows, err := q.db.QueryContext(ctx, listTimeEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TimeEntry
	for rows.Next() {
		var i TimeEntry
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.StartTime,
			&i.EndTime,
			&i.Duration,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTimeEntry = `-- name: UpdateTimeEntry :exec
UPDATE time_entries
SET start_time = ?,
end_time = ?,
description = ?
WHERE id = ?
`

type UpdateTimeEntryParams struct {
	StartTime   time.Time
	EndTime     time.Time
	Description sql.NullString
	ID          int64
}

func (q *Queries) UpdateTimeEntry(ctx context.Context, arg UpdateTimeEntryParams) error {
	_, err := q.db.ExecContext(ctx, updateTimeEntry,
		arg.StartTime,
		arg.EndTime,
		arg.Description,
		arg.ID,
	)
	return err
}
