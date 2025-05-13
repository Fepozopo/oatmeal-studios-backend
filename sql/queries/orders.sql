-- name: GetOrder :one
SELECT *
FROM orders
WHERE id = $1;

-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY order_date DESC;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, order_date, status, type, method, ship_date, po_number, shipping_cost, free_shipping, apply_to_commission, notes, sales_rep)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, created_at, updated_at, customer_id, order_date, status, type, method, ship_date, po_number, shipping_cost, free_shipping, apply_to_commission, notes, sales_rep;

-- name: UpdateOrder :one
UPDATE orders
SET customer_id = $2,
    order_date = $3,
    status = $4,
    type = $5,
    method = $6,
    ship_date = $7,
    po_number = $8,
    shipping_cost = $9,
    free_shipping = $10,
    apply_to_commission = $11,
    notes = $12,
    sales_rep = $13,
    updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, customer_id, order_date, status, type, method, ship_date, po_number, shipping_cost, free_shipping, apply_to_commission, notes, sales_rep;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;
