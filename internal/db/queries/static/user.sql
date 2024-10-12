-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUser :many
SELECT * FROM user
ORDER BY username;

-- name: CreateUsers :one
INSERT INTO user (
  username, email, password_hash
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE user
SET username = ?,
email = ?,
password_hash = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
