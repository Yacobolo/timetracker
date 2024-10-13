-- name: GetProject :one
SELECT * FROM project
WHERE id = $1
LIMIT 1;

-- name: ListProjects :many
SELECT * FROM project
ORDER BY name;

-- name: CreateProject :one
INSERT INTO project (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateProject :exec
UPDATE project
SET name = $1,
description = $2,
updated_at = CURRENT_TIMESTAMP
WHERE id = $3;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1;
