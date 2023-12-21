package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Warehouse interface {
	Querier
	CreateProduct(ctx context.Context, arg CreateProductParams) (CreateProductResult, error)
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewWarehouse(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
