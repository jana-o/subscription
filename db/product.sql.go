package db

import (
	"context"
	models2 "github.com/jana-o/subscription/models"
)

const getProductByID = `-- name: GetProduct :one
SELECT id, name, duration, price, description FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductByID(ctx context.Context, id int64) (models2.Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i models2.Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Duration,
		&i.Price,
		&i.Description,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, duration, price, description FROM products
WHERE account_id = $1
ORDER BY id
`
// maybe add limit here

func (q *Queries) GetProducts(ctx context.Context) ([]models2.Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models2.Product{}
	for rows.Next() {
		var i models2.Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Duration,
			&i.Price,
			&i.Description,
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