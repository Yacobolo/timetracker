-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUserByProvider :one
SELECT * FROM "user"
WHERE provider = $1 AND provider_user_id = $2
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY created_at;

-- name: CreateUser :one
INSERT INTO "user" (
  provider, provider_user_id, email, profile_picture
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET provider = $1,
provider_user_id = $2,
email = $3,
profile_picture = $4,
updated_at = CURRENT_TIMESTAMP
WHERE id = $5;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
