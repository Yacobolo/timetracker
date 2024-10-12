-- name: GetProject :one
SELECT * FROM project
WHERE id = ? LIMIT 1;

-- name: ListProjects :many
SELECT * FROM project
ORDER BY name;

-- name: CreateProject :one
INSERT INTO project (
  name, description
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateProject :exec
UPDATE project
SET name = ?,
description = ?
WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = ?;
