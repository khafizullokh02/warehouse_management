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

-- name: GetUserByParams :one
SELECT *
FROM users
WHERE email = sqlc.arg(email)
LIMIT 1;

-- name: GetUser :one
SELECT * 
FROM users
WHERE id = sqlc.arg(id) 
LIMIT 1;

-- name: ListUsers :many
SELECT * 
FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
name = COALESCE(sqlc.arg(name), name),
password = COALESCE(sqlc.arg(password), password)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);
