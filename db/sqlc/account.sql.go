// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: account.sql

package db

import (
	"context"

	zero "gopkg.in/guregu/null.v4/zero"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    name
)
VALUES (
    $1,
    $2
)
RETURNING id, user_id, name, created_at
`

type CreateAccountParams struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.UserID, arg.Name)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, user_id, name, created_at
FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, user_id, name, created_at
FROM accounts
ORDER BY id DESC
LIMIT $1 OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
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

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET name = COALESCE($1, name)
WHERE id = $2
RETURNING id, user_id, name, created_at
`

type UpdateAccountParams struct {
	Name zero.String `json:"name"`
	ID   int32       `json:"id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.Name, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}
