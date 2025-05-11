-- name: GetOrderItem :one
SELECT *
FROM order_items
WHERE id = $1;

-- name: ListOrderItemsByOrderID :many
SELECT *
FROM order_items
WHERE order_id = $1;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, item, quantity, price, discount, item_total)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, order_id, item, quantity, price, discount, item_total;

-- name: UpdateOrderItem :one
UPDATE order_items
SET item = $2,
    quantity = $3,
    price = $4,
    discount = $5,
    item_total = $6
WHERE id = $1
RETURNING id, order_id, item, quantity, price, discount, item_total;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;

-- name: DeleteOrderItemsByOrderID :exec
DELETE FROM order_items
WHERE order_id = $1;
