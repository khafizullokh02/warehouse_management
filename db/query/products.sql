-- name: CreateProduct :one
INSERT INTO products (
  name,
  sup_code,
  bar_code,
  image,
  brand,
  wholesale_price,
  retail_price
) VALUES (
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
SET name = name = COALESCE(sqlc.narg(name), name),
sup_code = COALESCE(sqlc.narg(sup_code), sup_code),
bar_code = COALESCE(sqlc.narg(bar_code), bar_code),
image = COALESCE(sqlc.narg(image), image),
brand = COALESCE(sqlc.narg(brand), brand),
wholesale_price = COALESCE(sqlc.narg(wholesale_price), wholesale_price),
retail_price = COALESCE(sqlc.narg(retail_price), retail_price),
discount = COALESCE(sqlc.narg(discount), discount),
created_at = COALESCE(sqlc.narg(created_at), created_at)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = sqlc.arg(id);
