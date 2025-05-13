-- name: CreateProduct :one
INSERT INTO products (
    type, sku, upc, status, cost, price, envelope, artist, category, release_date, last_bought_date, description, text_front, text_inside
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

-- name: GetProductByID :one
SELECT * FROM products WHERE id = $1;

-- name: GetProductBySKU :one
SELECT * FROM products WHERE sku = $1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- name: UpdateProduct :one
UPDATE products SET
    type = $2,
    sku = $3,
    upc = $4,
    status = $5,
    cost = $6,
    price = $7,
    envelope = $8,
    artist = $9,
    category = $10,
    release_date = $11,
    last_bought_date = $12,
    description = $13,
    text_front = $14,
    text_inside = $15,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;
