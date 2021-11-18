package db

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	assert "github.com/matryer/is"
	"regexp"
	"testing"
)

func TestQueries_GetProductByID(t *testing.T) {
	is := assert.New(t)
	query:= `SELECT id, name, duration, price, description, created_at, updated_at, deleted_at FROM products WHERE id = $1 LIMIT 1`
	query = regexp.QuoteMeta(query)
	id := uuid.MustParse("188e6a6b-6f3b-435c-94bf-94c5f748494a")

	product := Product{
		ID:          uuid.MustParse("188e6a6b-6f3b-435c-94bf-94c5f748494a"),
		Name:        "Product1",
		Duration:    3,
		Price:       12.99,
		Description: "3-months subscription",
	}

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		sqlMock func() (*sql.DB, sqlmock.Sqlmock)
		args    args
		want    Product
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := tt.sqlMock()
			q := &Queries{
				db: db,
			}
			got, err := q.GetProductByID(tt.args.ctx, tt.args.id)
			is.Equal(got, tt.want)           // result and expectation are different
			is.Equal(err != nil, tt.wantErr) // error and expectation are different
			err = mock.ExpectationsWereMet()
			is.NoErr(err)
		})
	}
}