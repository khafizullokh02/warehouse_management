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
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListProducts :many
SELECT * 
FROM products
WHERE name like '%' || sqlc.arg('search') || '%'
ORDER BY id
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateProduct :one
UPDATE products
SET
name = COALESCE(sqlc.arg(name), name), 
sup_code = COALESCE(sqlc.arg(sup_code), sup_code),
bar_code = COALESCE(sqlc.arg(bar_code), bar_code),
image = COALESCE(sqlc.arg(image), image),
brand = COALESCE(sqlc.arg(brand), brand),
wholesale_price = COALESCE(sqlc.arg(wholesale_price), wholesale_price),
retail_price = COALESCE(sqlc.arg(retail_price), retail_price),
discount = COALESCE(sqlc.arg(discount), discount),
created_at = COALESCE(sqlc.arg(created_at), created_at)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = sqlc.arg(id);
