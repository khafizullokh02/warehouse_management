-- name: CreateAccount :one
INSERT INTO accounts (
  id,
  user_id,
  name
) VALUES (
  sqlc.arg(id),
  sqlc.arg(user_id),
  sqlc.arg(name)
) RETURNING *;

-- name: GetAccount :one
SELECT * 
FROM accounts
WHERE id = sqlc.arg(id) LIMIT 1;

-- name: ListAccounts :many
SELECT * 
FROM accounts
ORDER BY id
LIMIT sqlc.arg(id)
OFFSET sqlc.arg(user_id);

-- name: UpdateAccount :one
UPDATE accounts
SET name = sqlc.arg(name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = sqlc.arg(id);
