-- name: CreateEntryGroup :one
INSERT INTO entry_group (
  id,
  quantity,
  action_type,
  pricing_type,
  price,
  currency,
  entry_group_status
) VALUES (
  sqlc.arg(id),
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
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateEntryGroup :one
UPDATE entry_group
SET price = sqlc.arg(price)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteEntryGroup :exec
DELETE FROM entry_group
WHERE id = sqlc.arg(id);
