-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products
ORDER BY id;