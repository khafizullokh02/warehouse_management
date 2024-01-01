-- name: CreateAgreementForms :one
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
