-- name: CreateSession :one
INSERT INTO sessions (
    user_agent,
    client_ip,
    is_blocked,
    user_id
  )
VALUES (
    sqlc.arg(user_agent),
    sqlc.arg(client_ip),
    sqlc.arg(is_blocked),
    sqlc.arg(user_id)
  )
RETURNING *;

-- name: GetSession :one
SELECT *
FROM sessions
WHERE id = sqlc.arg(id)
LIMIT 1;