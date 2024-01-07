-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    name
)
VALUES (
    sqlc.arg(user_id),
    sqlc.arg(name)
)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateAccount :one
UPDATE accounts
SET name = COALESCE(sqlc.narg(name), name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = sqlc.arg(id);