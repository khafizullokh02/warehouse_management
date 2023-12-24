-- name: CreateProduct :one
INSERT INTO products (
  id,
  name,
  sup_code,
  bar_code,
  image,
  brand,
  wholesale_price,
  retail_price
) VALUES (
  sqlc.arg(id),
  sqlc.arg(name),
  sqlc.arg(sup_code),
  sqlc.arg(bar_code),
  sqlc.arg(image),
  sqlc.arg(brand),
  sqlc.arg(wholesale_price),
  sqlc.arg(retail_price)
) RETURNING *;

-- name: GetProduct :one
SELECT * 
FROM products
WHERE name = sqlc.arg(id) 
LIMIT 1;

-- name: ListProducts :many
SELECT * 
FROM products
ORDER BY name
LIMIT sqlc.arg(id)
OFFSET sqlc.arg(name);

-- name: UpdateProduct :one
UPDATE products
SET name = sqlc.arg(name)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = sqlc.arg(id);
