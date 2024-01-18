-- name: CreateUser :one
INSERT INTO users (
  name,
  role,
  email,
  password
) VALUES (
  sqlc.arg(name),
  sqlc.arg(role),
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
password = COALESCE(sqlc.arg(password), password)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);
