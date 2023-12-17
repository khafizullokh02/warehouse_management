// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: entry_items.sql

package db

import (
	"context"
)

const createEntryItems = `-- name: CreateEntryItems :one
INSERT INTO entry_items (
  id,
  product_id,
  entry_group_id,
  sup_code
) VALUES (
  $1, $2, $3, $4
) RETURNING id, product_id, entry_group_id, sup_code
`

type CreateEntryItemsParams struct {
	ID           int32  `json:"id"`
	ProductID    int32  `json:"product_id"`
	EntryGroupID int32  `json:"entry_group_id"`
	SupCode      string `json:"sup_code"`
}

func (q *Queries) CreateEntryItems(ctx context.Context, arg CreateEntryItemsParams) (EntryItem, error) {
	row := q.db.QueryRowContext(ctx, createEntryItems,
		arg.ID,
		arg.ProductID,
		arg.EntryGroupID,
		arg.SupCode,
	)
	var i EntryItem
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.EntryGroupID,
		&i.SupCode,
	)
	return i, err
}

const deleteEntryItems = `-- name: DeleteEntryItems :exec
DELETE FROM entry_items
WHERE id = $1
`

func (q *Queries) DeleteEntryItems(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEntryItems, id)
	return err
}

const getEntryItems = `-- name: GetEntryItems :one
SELECT id, product_id, entry_group_id, sup_code FROM entry_items
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntryItems(ctx context.Context, id int32) (EntryItem, error) {
	row := q.db.QueryRowContext(ctx, getEntryItems, id)
	var i EntryItem
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.EntryGroupID,
		&i.SupCode,
	)
	return i, err
}

const listEntryItems = `-- name: ListEntryItems :many
SELECT id, product_id, entry_group_id, sup_code FROM entry_items
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntryItemsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntryItems(ctx context.Context, arg ListEntryItemsParams) ([]EntryItem, error) {
	rows, err := q.db.QueryContext(ctx, listEntryItems, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EntryItem{}
	for rows.Next() {
		var i EntryItem
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.EntryGroupID,
			&i.SupCode,
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

const updateEntryItems = `-- name: UpdateEntryItems :one
UPDATE entry_items
SET product_id = $2
WHERE id = $1
RETURNING id, product_id, entry_group_id, sup_code
`

type UpdateEntryItemsParams struct {
	ID        int32 `json:"id"`
	ProductID int32 `json:"product_id"`
}

func (q *Queries) UpdateEntryItems(ctx context.Context, arg UpdateEntryItemsParams) (EntryItem, error) {
	row := q.db.QueryRowContext(ctx, updateEntryItems, arg.ID, arg.ProductID)
	var i EntryItem
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.EntryGroupID,
		&i.SupCode,
	)
	return i, err
}
