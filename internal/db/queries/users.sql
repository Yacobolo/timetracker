-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, email, password_hash
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET username = ?,
email = ?,
password_hash = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
