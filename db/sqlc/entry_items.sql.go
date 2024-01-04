// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: entry_items.sql

package db

import (
	"context"

	zero "gopkg.in/guregu/null.v4/zero"
)

const createEntryItem = `-- name: CreateEntryItem :one
INSERT INTO entry_items (product_id, entry_group_id, sup_code)
VALUES (
    $1,
    $2,
    $3
  )
RETURNING id, product_id, entry_group_id, sup_code, created_at
`

type CreateEntryItemParams struct {
	ProductID    int32  `json:"product_id"`
	EntryGroupID int32  `json:"entry_group_id"`
	SupCode      string `json:"sup_code"`
}

func (q *Queries) CreateEntryItem(ctx context.Context, arg CreateEntryItemParams) (EntryItem, error) {
	row := q.db.QueryRow(ctx, createEntryItem, arg.ProductID, arg.EntryGroupID, arg.SupCode)
	var i EntryItem
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.EntryGroupID,
		&i.SupCode,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntryItem = `-- name: DeleteEntryItem :exec
DELETE FROM entry_items
WHERE id = $1
`

func (q *Queries) DeleteEntryItem(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteEntryItem, id)
	return err
}

const getEntryItem = `-- name: GetEntryItem :one
SELECT id, product_id, entry_group_id, sup_code, created_at
FROM entry_items
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEntryItem(ctx context.Context, id int32) (EntryItem, error) {
	row := q.db.QueryRow(ctx, getEntryItem, id)
	var i EntryItem
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.EntryGroupID,
		&i.SupCode,
		&i.CreatedAt,
	)
	return i, err
}

const listEntryItems = `-- name: ListEntryItems :many
SELECT id, product_id, entry_group_id, sup_code, created_at
FROM entry_items
ORDER BY id DESC
LIMIT $2 OFFSET $1
`

type ListEntryItemsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListEntryItems(ctx context.Context, arg ListEntryItemsParams) ([]EntryItem, error) {
	rows, err := q.db.Query(ctx, listEntryItems, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EntryItem
	for rows.Next() {
		var i EntryItem
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.EntryGroupID,
			&i.SupCode,
			&i.CreatedAt,
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

const updateEntryItem = `-- name: UpdateEntryItem :one
UPDATE entry_items
SET product_id = COALESCE($1, product_id),
  entry_group_id = COALESCE($2, entry_group_id),
  sup_code = COALESCE($3, sup_code)
WHERE id = $4
RETURNING id, product_id, entry_group_id, sup_code, created_at
`

type UpdateEntryItemParams struct {
	ProductID    zero.Int    `json:"product_id"`
	EntryGroupID zero.Int    `json:"entry_group_id"`
	SupCode      zero.String `json:"sup_code"`
	ID           int32       `json:"id"`
}

func (q *Queries) UpdateEntryItem(ctx context.Context, arg UpdateEntryItemParams) (EntryItem, error) {
	row := q.db.QueryRow(ctx, updateEntryItem,
		arg.ProductID,
		arg.EntryGroupID,
		arg.SupCode,
		arg.ID,
	)
	var i EntryItem
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.EntryGroupID,
		&i.SupCode,
		&i.CreatedAt,
	)
	return i, err
}
