-- name: GetSubscriptionByID :one
SELECT * FROM user_subscriptions
WHERE id = $1 LIMIT 1;
