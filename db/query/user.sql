-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password
) VALUES (
  sqlc.arg(name),
  sqlc.arg(email),
  sqlc.arg(password)
) RETURNING *;

-- name: GetUser :one
SELECT * 
FROM users
WHERE id = sqlc.arg(id) 
LIMIT 1;

-- name: ListUsers :many
SELECT * 
FROM users
WHERE name = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET
name = COALESCE(sqlc.arg(name), name),
email = COALESCE(sqlc.arg(email), email),
password = COALESCE(sqlc.arg(password), password),
created_at = COALESCE(sqlc.arg(created_at), created_at),
updated_at = COALESCE(sqlc.arg(updated_at), updated_at),
deleted_at = COALESCE(sqlc.arg(deleted_at), deleted_at)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);
