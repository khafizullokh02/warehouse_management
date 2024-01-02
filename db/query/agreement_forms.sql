-- name: CreateAgreementForm :one
INSERT INTO agreement_forms (
    from_account,
    to_account,
    product_ids,
    action_type_for_agreement,
    wholesale_price,
    retail_price
  )
VALUES (
    sqlc.arg(from_account),
    sqlc.arg(to_account),
    sqlc.arg(product_ids),
    sqlc.arg(action_type_for_agreement),
    sqlc.arg(wholesale_price),
    sqlc.arg(retail_price)
  )
RETURNING *;

-- name: GetAgreementForm :one
SELECT *
FROM agreement_forms
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListAgreementForms :many
SELECT *
FROM agreement_forms
WHERE 
    from_account = $1 OR
    to_account = $2
ORDER BY id DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: UpdateAgreementForm :one
UPDATE agreement_forms
SET 
product_ids = COALESCE(sqlc.arg(product_ids), product_ids),
action_type_for_agreement = COALESCE(sqlc.arg(action_type_for_agreement), action_type_for_agreement),
wholesale_price = COALESCE(sqlc.arg(wholesale_price), wholesale_price),
retail_price = COALESCE(sqlc.arg(retail_price), retail_price)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAgreementForm :exec
DELETE FROM agreement_forms
WHERE id = sqlc.arg(id);
