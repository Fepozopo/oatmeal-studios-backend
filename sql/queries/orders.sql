-- name: GetOrder :one
SELECT *
FROM orders
WHERE id = $1;

-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY order_date DESC;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, customer_location_id, order_date, status, type, method, ship_date, po_number, shipping_cost, free_shipping, apply_to_commission, sales_rep, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;

-- name: UpdateOrder :one
UPDATE orders
SET customer_id = $2,
    customer_location_id = $3,
    order_date = $4,
    status = $5,
    type = $6,
    method = $7,
    ship_date = $8,
    po_number = $9,
    shipping_cost = $10,
    free_shipping = $11,
    apply_to_commission = $12,
    sales_rep = $13,
    notes = $14,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;
