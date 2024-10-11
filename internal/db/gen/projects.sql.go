// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: projects.sql

package db

import (
	"context"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
  name, description
) VALUES (
  ?, ?
)
RETURNING id, name, description, created_at, updated_at
`

type CreateProjectParams struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, arg.Name, arg.Description)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = ?
`

func (q *Queries) DeleteProject(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, name, description, created_at, updated_at FROM projects
WHERE id = ? LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, description, created_at, updated_at FROM projects
ORDER BY name
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const updateProject = `-- name: UpdateProject :exec
UPDATE projects
SET name = ?,
description = ?
WHERE id = ?
`

type UpdateProjectParams struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	ID          int64  `db:"id"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) error {
	_, err := q.db.ExecContext(ctx, updateProject, arg.Name, arg.Description, arg.ID)
	return err
}
