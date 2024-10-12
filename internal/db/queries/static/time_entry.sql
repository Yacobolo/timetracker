-- name: GetTimeEntry :one
SELECT * FROM time_entry
WHERE id = ? LIMIT 1;

-- name: ListTimeEntries :many
SELECT * FROM time_entry
ORDER BY start_time;

-- name: CreateTimeEntry :one
INSERT INTO time_entry (
  project_id, start_time, end_time, description
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateTimeEntry :exec
UPDATE time_entry
SET start_time = ?,
end_time = ?,
description = ?
WHERE id = ?;

-- name: DeleteTimeEntry :exec
DELETE FROM time_entry
WHERE id = ?;
