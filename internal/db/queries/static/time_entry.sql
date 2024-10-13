-- name: GetTimeEntry :one
SELECT * FROM time_entry
WHERE id = $1
LIMIT 1;

-- name: ListTimeEntries :many
SELECT * FROM time_entry
ORDER BY start_time;

-- name: CreateTimeEntry :one
INSERT INTO time_entry (
  project_id, start_time, end_time, description
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTimeEntry :exec
UPDATE time_entry
SET start_time = $1,
end_time = $2,
description = $3,
updated_at = CURRENT_TIMESTAMP
WHERE id = $4;

-- name: DeleteTimeEntry :exec
DELETE FROM time_entry
WHERE id = $1;
