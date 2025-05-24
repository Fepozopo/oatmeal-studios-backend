-- name: GetOrderItem :one
SELECT *
FROM order_items
WHERE id = $1;

-- name: ListOrderItemsBySKU :many
SELECT *
FROM order_items
WHERE sku = $1;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, sku, quantity, price, discount, item_total, pocket_number)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateOrderItem :one
UPDATE order_items
SET sku = $2,
    quantity = $3,
    price = $4,
    discount = $5,
    item_total = $6,
    pocket_number = $7
WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;
