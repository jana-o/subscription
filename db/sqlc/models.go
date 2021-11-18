// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive   Status = "Active"
	StatusPaused   Status = "Paused"
	StatusCanceled Status = "Canceled"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type Product struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Duration    int32        `json:"duration"`
	Price       float64      `json:"price"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Active    sql.NullBool `json:"active"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type UserSubscription struct {
	ID         uuid.UUID       `json:"id"`
	UserID     uuid.UUID       `json:"user_id"`
	ProductID  uuid.UUID       `json:"product_id"`
	TrialStart sql.NullTime    `json:"trial_start"`
	TrialEnd   sql.NullTime    `json:"trial_end"`
	StartDate  sql.NullTime    `json:"start_date"`
	EndDate    sql.NullTime    `json:"end_date"`
	Discount   sql.NullFloat64 `json:"discount"`
	Tax        sql.NullFloat64 `json:"tax"`
	Status     Status          `json:"status"`
	CreatedAt  sql.NullTime    `json:"created_at"`
	PausedAt   sql.NullTime    `json:"paused_at"`
	UpdatedAt  sql.NullTime    `json:"updated_at"`
	DeletedAt  sql.NullTime    `json:"deleted_at"`
}
