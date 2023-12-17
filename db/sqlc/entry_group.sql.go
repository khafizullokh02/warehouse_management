// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: entry_group.sql

package db

import (
	"context"
)

const createEntryGroup = `-- name: CreateEntryGroup :one
INSERT INTO entry_group (
  id,
  quantity,
  action_type,
  pricing_type,
  price,
  currency
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, quantity, action_type, pricing_type, price, currency, entry_groups_status, created_at, updated_at, deleted_at
`

type CreateEntryGroupParams struct {
	ID          int32        `json:"id"`
	Quantity    int32        `json:"quantity"`
	ActionType  ActionType   `json:"action_type"`
	PricingType PricingType  `json:"pricing_type"`
	Price       float64      `json:"price"`
	Currency    CurrencyCode `json:"currency"`
}

func (q *Queries) CreateEntryGroup(ctx context.Context, arg CreateEntryGroupParams) (EntryGroup, error) {
	row := q.db.QueryRowContext(ctx, createEntryGroup,
		arg.ID,
		arg.Quantity,
		arg.ActionType,
		arg.PricingType,
		arg.Price,
		arg.Currency,
	)
	var i EntryGroup
	err := row.Scan(
		&i.ID,
		&i.Quantity,
		&i.ActionType,
		&i.PricingType,
		&i.Price,
		&i.Currency,
		&i.EntryGroupsStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteEntryGroup = `-- name: DeleteEntryGroup :exec
DELETE FROM entry_group
WHERE id = $1
`

func (q *Queries) DeleteEntryGroup(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEntryGroup, id)
	return err
}

const getEntryGroup = `-- name: GetEntryGroup :one
SELECT id, quantity, action_type, pricing_type, price, currency, entry_groups_status, created_at, updated_at, deleted_at FROM entry_group
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntryGroup(ctx context.Context, id int32) (EntryGroup, error) {
	row := q.db.QueryRowContext(ctx, getEntryGroup, id)
	var i EntryGroup
	err := row.Scan(
		&i.ID,
		&i.Quantity,
		&i.ActionType,
		&i.PricingType,
		&i.Price,
		&i.Currency,
		&i.EntryGroupsStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listEntryGroups = `-- name: ListEntryGroups :many
SELECT id, quantity, action_type, pricing_type, price, currency, entry_groups_status, created_at, updated_at, deleted_at FROM entry_group
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntryGroupsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntryGroups(ctx context.Context, arg ListEntryGroupsParams) ([]EntryGroup, error) {
	rows, err := q.db.QueryContext(ctx, listEntryGroups, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EntryGroup{}
	for rows.Next() {
		var i EntryGroup
		if err := rows.Scan(
			&i.ID,
			&i.Quantity,
			&i.ActionType,
			&i.PricingType,
			&i.Price,
			&i.Currency,
			&i.EntryGroupsStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntryGroup = `-- name: UpdateEntryGroup :one
UPDATE entry_group
SET price = $2
WHERE id = $1
RETURNING id, quantity, action_type, pricing_type, price, currency, entry_groups_status, created_at, updated_at, deleted_at
`

type UpdateEntryGroupParams struct {
	ID    int32   `json:"id"`
	Price float64 `json:"price"`
}

func (q *Queries) UpdateEntryGroup(ctx context.Context, arg UpdateEntryGroupParams) (EntryGroup, error) {
	row := q.db.QueryRowContext(ctx, updateEntryGroup, arg.ID, arg.Price)
	var i EntryGroup
	err := row.Scan(
		&i.ID,
		&i.Quantity,
		&i.ActionType,
		&i.PricingType,
		&i.Price,
		&i.Currency,
		&i.EntryGroupsStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
