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
SELECT * 
FROM entry_items
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListEntryItems :many
SELECT * 
FROM entry_items
ORDER BY id
LIMIT sqlc.arg(id)
OFFSET sqlc.arg(product_id);

-- name: UpdateEntryItems :one
UPDATE entry_items
SET product_id = sqlc.arg(product_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteEntryItems :exec
DELETE FROM entry_items
WHERE id = sqlc.arg(id);
