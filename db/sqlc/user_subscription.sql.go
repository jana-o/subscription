// Code generated by sqlc. DO NOT EDIT.
// source: user_subscription.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const cancelSubscription = `-- name: CancelSubscription :one
UPDATE user_subscriptions
SET Status = 'Canceled' AND Updated_at= CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, product_id, trial_start, trial_end, start_date, end_date, discount, tax, status, created_at, paused_at, updated_at, deleted_at
`

func (q *Queries) CancelSubscription(ctx context.Context, id uuid.UUID) (UserSubscription, error) {
	row := q.db.QueryRowContext(ctx, cancelSubscription, id)
	var i UserSubscription
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.TrialStart,
		&i.TrialEnd,
		&i.StartDate,
		&i.EndDate,
		&i.Discount,
		&i.Tax,
		&i.Status,
		&i.CreatedAt,
		&i.PausedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createSubscription = `-- name: CreateSubscription :one
INSERT INTO user_subscriptions (User_id,
                      Product_id,
                      Start_date,
                      End_date,
                      Tax) VALUES (
                                    $1,$2,$3,$4,$5
) RETURNING id, user_id, product_id, trial_start, trial_end, start_date, end_date, discount, tax, status, created_at, paused_at, updated_at, deleted_at
`

type CreateSubscriptionParams struct {
	UserID    uuid.UUID       `json:"user_id"`
	ProductID uuid.UUID       `json:"product_id"`
	StartDate sql.NullTime    `json:"start_date"`
	EndDate   sql.NullTime    `json:"end_date"`
	Tax       sql.NullFloat64 `json:"tax"`
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) (UserSubscription, error) {
	row := q.db.QueryRowContext(ctx, createSubscription,
		arg.UserID,
		arg.ProductID,
		arg.StartDate,
		arg.EndDate,
		arg.Tax,
	)
	var i UserSubscription
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.TrialStart,
		&i.TrialEnd,
		&i.StartDate,
		&i.EndDate,
		&i.Discount,
		&i.Tax,
		&i.Status,
		&i.CreatedAt,
		&i.PausedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getSubscriptionByID = `-- name: GetSubscriptionByID :one
SELECT id, user_id, product_id, trial_start, trial_end, start_date, end_date, discount, tax, status, created_at, paused_at, updated_at, deleted_at FROM user_subscriptions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSubscriptionByID(ctx context.Context, id uuid.UUID) (UserSubscription, error) {
	row := q.db.QueryRowContext(ctx, getSubscriptionByID, id)
	var i UserSubscription
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.TrialStart,
		&i.TrialEnd,
		&i.StartDate,
		&i.EndDate,
		&i.Discount,
		&i.Tax,
		&i.Status,
		&i.CreatedAt,
		&i.PausedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const pauseSubscription = `-- name: PauseSubscription :one
UPDATE user_subscriptions
SET Paused_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, product_id, trial_start, trial_end, start_date, end_date, discount, tax, status, created_at, paused_at, updated_at, deleted_at
`

func (q *Queries) PauseSubscription(ctx context.Context, id uuid.UUID) (UserSubscription, error) {
	row := q.db.QueryRowContext(ctx, pauseSubscription, id)
	var i UserSubscription
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.TrialStart,
		&i.TrialEnd,
		&i.StartDate,
		&i.EndDate,
		&i.Discount,
		&i.Tax,
		&i.Status,
		&i.CreatedAt,
		&i.PausedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
