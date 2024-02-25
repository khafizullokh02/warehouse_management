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

-- name: GetAllSessions :many
SELECT *
FROM sessions
WHERE user_id = sqlc.arg(user_id)
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');