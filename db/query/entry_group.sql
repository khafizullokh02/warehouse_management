-- name: CreateEntryGroup :one
INSERT INTO entry_group (
  id,
  quantity,
  action_type,
  pricing_type,
  price,
  currency
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetEntryGroup :one
SELECT * FROM entry_group
WHERE id = $1 LIMIT 1;

-- name: ListEntryGroups :many
SELECT * FROM entry_group
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntryGroup :one
UPDATE entry_group
SET price = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntryGroup :exec
DELETE FROM entry_group
WHERE id = $1;
