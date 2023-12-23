-- name: CreateProduct :one
INSERT INTO products (
    name,
    sup_code,
    bar_code,
    image,
    brand,
    wholesale_price,
    retail_price
  )
VALUES (
    sqlc.arg(name),
    sqlc.arg(sup_code),
    sqlc.arg(bar_code),
    sqlc.arg(image),
    sqlc.arg(brand),
    sqlc.arg(wholesale_price),
    sqlc.arg(retail_price)
  )
RETURNING *;

-- name: GetProduct :one
SELECT *
FROM products
WHERE name = $1
LIMIT 1;

-- name: ListProducts :many
SELECT *
FROM products
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;