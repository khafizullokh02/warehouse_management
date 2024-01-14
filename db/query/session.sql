-- name: CreateSession :one
INSERT INTO sessions (
  id,
  name,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  sqlc.arg(id),
  sqlc.arg(name),
  sqlc.arg(refresh_token),
  sqlc.arg(user_agent),
  sqlc.arg(client_ip),
  sqlc.arg(is_blocked),
  sqlc.arg(expires_at)
) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = sqlc.arg(id)
LIMIT 1;