-- name: CreateEntryItem :one
INSERT INTO entry_items (product_id, entry_group_id, sup_code)
VALUES (
    sqlc.arg(product_id),
    sqlc.arg(entry_group_id),
    sqlc.arg(sup_code)
  )
RETURNING *;

-- name: GetEntryItem :one
SELECT *
FROM entry_items
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListEntryItems :many
SELECT *
FROM entry_items
ORDER BY id DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UpdateEntryItem :one
UPDATE entry_items
SET product_id = COALESCE(sqlc.narg(product_id), product_id),
  entry_group_id = COALESCE(sqlc.narg(entry_group_id), entry_group_id),
  sup_code = COALESCE(sqlc.narg(sup_code), sup_code)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteEntryItem :exec
DELETE FROM entry_items
WHERE id = sqlc.arg(id);