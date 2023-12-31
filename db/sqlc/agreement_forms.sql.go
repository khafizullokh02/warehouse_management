// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: agreement_forms.sql

package db

import (
	"context"
)

const createAgreementForms = `-- name: CreateAgreementForms :one
INSERT INTO agreement_forms (
    id,
    from_account,
    to_account,
    product_ids,
    action_type_for_agreement
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
  )
RETURNING id, from_account, to_account, product_ids, action_type_for_agreement, wholesale_price, retail_price
`

type CreateAgreementFormsParams struct {
	ID                     int32      `json:"id"`
	FromAccount            int32      `json:"from_account"`
	ToAccount              int32      `json:"to_account"`
	ProductIds             []int32    `json:"product_ids"`
	ActionTypeForAgreement ActionType `json:"action_type_for_agreement"`
}

func (q *Queries) CreateAgreementForms(ctx context.Context, arg CreateAgreementFormsParams) (AgreementForm, error) {
	row := q.db.QueryRow(ctx, createAgreementForms,
		arg.ID,
		arg.FromAccount,
		arg.ToAccount,
		arg.ProductIds,
		arg.ActionTypeForAgreement,
	)
	var i AgreementForm
	err := row.Scan(
		&i.ID,
		&i.FromAccount,
		&i.ToAccount,
		&i.ProductIds,
		&i.ActionTypeForAgreement,
		&i.WholesalePrice,
		&i.RetailPrice,
	)
	return i, err
}

const deleteAgreementForm = `-- name: DeleteAgreementForm :exec
DELETE FROM agreement_forms
WHERE id = $1
`

func (q *Queries) DeleteAgreementForm(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAgreementForm, id)
	return err
}

const getAgreementForm = `-- name: GetAgreementForm :one
SELECT id, from_account, to_account, product_ids, action_type_for_agreement, wholesale_price, retail_price
FROM agreement_forms
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAgreementForm(ctx context.Context, id int32) (AgreementForm, error) {
	row := q.db.QueryRow(ctx, getAgreementForm, id)
	var i AgreementForm
	err := row.Scan(
		&i.ID,
		&i.FromAccount,
		&i.ToAccount,
		&i.ProductIds,
		&i.ActionTypeForAgreement,
		&i.WholesalePrice,
		&i.RetailPrice,
	)
	return i, err
}

const listAgreementForms = `-- name: ListAgreementForms :many
SELECT id, from_account, to_account, product_ids, action_type_for_agreement, wholesale_price, retail_price
FROM agreement_forms
ORDER BY id DESC
LIMIT $2 OFFSET $1
`

type ListAgreementFormsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListAgreementForms(ctx context.Context, arg ListAgreementFormsParams) ([]AgreementForm, error) {
	rows, err := q.db.Query(ctx, listAgreementForms, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AgreementForm
	for rows.Next() {
		var i AgreementForm
		if err := rows.Scan(
			&i.ID,
			&i.FromAccount,
			&i.ToAccount,
			&i.ProductIds,
			&i.ActionTypeForAgreement,
			&i.WholesalePrice,
			&i.RetailPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAgreementForm = `-- name: UpdateAgreementForm :one
UPDATE agreement_forms
SET action_type_for_agreement = $1
WHERE id = $2
RETURNING id, from_account, to_account, product_ids, action_type_for_agreement, wholesale_price, retail_price
`

type UpdateAgreementFormParams struct {
	ActionTypeForAgreement ActionType `json:"action_type_for_agreement"`
	ID                     int32      `json:"id"`
}

func (q *Queries) UpdateAgreementForm(ctx context.Context, arg UpdateAgreementFormParams) (AgreementForm, error) {
	row := q.db.QueryRow(ctx, updateAgreementForm, arg.ActionTypeForAgreement, arg.ID)
	var i AgreementForm
	err := row.Scan(
		&i.ID,
		&i.FromAccount,
		&i.ToAccount,
		&i.ProductIds,
		&i.ActionTypeForAgreement,
		&i.WholesalePrice,
		&i.RetailPrice,
	)
	return i, err
}
