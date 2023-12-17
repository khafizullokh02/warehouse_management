-- name: CreateEntryItems :one
INSERT INTO entry_items (
  id,
  product_id,
  entry_group_id,
  sup_code
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetEntryItems :one
SELECT * FROM entry_items
WHERE id = $1 LIMIT 1;

-- name: ListEntryItems :many
SELECT * FROM entry_items
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntryItems :one
UPDATE entry_items
SET product_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntryItems :exec
DELETE FROM entry_items
WHERE id = $1;
