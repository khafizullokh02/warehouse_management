-- name: CreateEntryGroup :one
INSERT INTO entry_group (
  quantity,
  action_type,
  pricing_type,
  price,
  currency,
  entry_group_status
) VALUES (
  sqlc.arg(quantity),
  sqlc.arg(action_type),
  sqlc.arg(pricing_type),
  sqlc.arg(price),
  sqlc.arg(currency),
  sqlc.arg(entry_group_status)
) RETURNING *;

-- name: GetEntryGroup :one
SELECT * 
FROM entry_group
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListEntryGroups :many
SELECT * 
FROM entry_group
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntryGroup :one
UPDATE entry_group
SET price = sqlc.narg(price)
WHERE id = sqlc.narg(id)
RETURNING *;

-- name: DeleteEntryGroup :exec
DELETE FROM entry_group
WHERE id = sqlc.narg(id);
