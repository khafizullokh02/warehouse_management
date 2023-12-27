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
WHERE name = sqlc.arg(name) 
LIMIT 1;

-- name: ListUsers :many
SELECT * 
FROM users
ORDER BY id
LIMIT sqlc.arg(id)
OFFSET sqlc.arg(name);

-- name: UpdateUser :one
UPDATE users
SET email = sqlc.arg(email)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = sqlc.arg(id);
