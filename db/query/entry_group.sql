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
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateEntryGroup :one
UPDATE entry_group
SET 
quantity = COALESCE(sqlc.narg(quantity), quantity),
action_type = COALESCE(sqlc.narg(action_type), action_type),
pricing_type = COALESCE(sqlc.narg(pricing_type), pricing_type),
price = COALESCE(sqlc.narg(price), price),
currency = COALESCE(sqlc.narg(currency), currency),
entry_group_status = COALESCE(sqlc.narg(entry_group_status), entry_group_status)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteEntryGroup :exec
DELETE FROM entry_group
WHERE id = sqlc.arg(id);
