-- name: CreateProduct :one
INSERT INTO products (Name,
                      Duration,
                      Price,
                      Description) VALUES (
                                           $1,$2,$3,$4
) RETURNING *;

-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products
ORDER BY id;