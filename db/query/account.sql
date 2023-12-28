-- name: CreateAccount :one
INSERT INTO accounts (
  user_id,
  name
) VALUES (
  sqlc.arg(user_id),
  sqlc.arg(name)
) RETURNING *;

-- name: GetAccount :one
SELECT * 
FROM accounts
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListAccounts :many
SELECT * 
FROM accounts
WHERE name = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts
SET 
name = COALESCE(sqlc.arg(name), name),
user_id = COALESCE(sqlc.arg(user_id), user_id),
created_at = COALESCE(sqlc.arg(created_at), created_at)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = sqlc.arg(id);
