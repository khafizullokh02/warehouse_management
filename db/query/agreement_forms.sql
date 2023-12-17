-- name: CreateAgreementForms :one
INSERT INTO agreement_forms (
  id,
  from_account,
  to_account,
  product_ids,
  action_type_for_agreement,
  wholesale_price,
  retail_price
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetAgreementForm :one
SELECT * FROM agreement_forms
WHERE id = $1 LIMIT 1;

-- name: ListAgreementForms :many
SELECT * FROM agreement_forms
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAgreementForm :one
UPDATE agreement_forms
SET action_type_for_agreement = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAgreementForm :exec
DELETE FROM agreement_forms
WHERE id = $1;
