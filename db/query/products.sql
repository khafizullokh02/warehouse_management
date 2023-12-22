-- name: CreateProduct :one
INSERT INTO products (
  name,
  image,
  brand
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE name = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
