-- name: GetInvoice :one
SELECT *
FROM invoices
WHERE id = $1;

-- name: GetInvoicesByOrder :one
SELECT *
FROM invoices
WHERE order_id = $1;

-- name: ListInvoicesByCustomer :many
SELECT *
FROM invoices
WHERE customer_id = $1
ORDER BY invoice_date DESC;

-- name: ListInvoicesByCustomerLocation :many
SELECT *
FROM invoices
WHERE customer_location_id = $1
ORDER BY invoice_date DESC;

-- name: CreateInvoice :one
INSERT INTO invoices (invoice_date, order_id, customer_id, customer_location_id, due_date, status, total)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateInvoice :one
UPDATE invoices
SET invoice_date = $2,
    order_id = $3,
    customer_id = $4,
    customer_location_id = $5,
    due_date = $6,
    status = $7,
    total = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1;