package db

import (
	"context"
	models2 "github.com/jana-o/subscription/models"
)

type Querier interface {
	GetProductByID(ctx context.Context, id int64) (models2.Product, error)
	GetProducts(ctx context.Context) ([]models2.Product, error)
}

var _ Querier = (*Queries)(nil)
