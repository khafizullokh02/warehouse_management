-- name: CreateAgreementForms :one
INSERT INTO agreement_forms (
  id,
  from_account,
  to_account,
  product_ids,
  action_type_for_agreement
) VALUES (
  sqlc.arg(id),
  sqlc.arg(from_account),
  sqlc.arg(to_account),
  sqlc.arg(product_ids),
  sqlc.arg(action_type_for_agreement)
) RETURNING *;

-- name: GetAgreementForm :one
SELECT * 
FROM agreement_forms
WHERE id = sqlc.arg(id) 
LIMIT 1;

-- name: ListAgreementForms :many
SELECT * 
FROM agreement_forms
ORDER BY id
LIMIT sqlc.arg(id)
OFFSET sqlc.arg(from_account);

-- name: UpdateAgreementForm :one
UPDATE agreement_forms
SET action_type_for_agreement = sqlc.arg(action_type_for_agreement)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAgreementForm :exec
DELETE FROM agreement_forms
WHERE id = sqlc.arg(id);
