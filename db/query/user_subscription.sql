-- name: CreateSubscription :one
INSERT INTO user_subscriptions (User_id,
                      Product_id,
                      Start_date,
                      End_date,
                      Tax) VALUES (
                                    $1,$2,$3,$4,$5
) RETURNING *;

-- name: GetSubscriptionByID :one
SELECT * FROM user_subscriptions
WHERE id = $1 LIMIT 1;

-- name: PauseSubscription :one
UPDATE user_subscriptions
SET Paused_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING *;


-- name: CancelSubscription :one
UPDATE user_subscriptions
SET Status = 'Canceled' AND Updated_at= CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING *;