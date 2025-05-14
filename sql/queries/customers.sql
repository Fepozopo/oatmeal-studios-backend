-- name: GetCustomer :one
SELECT *
FROM customers
WHERE id = $1;

-- name: ListCustomers :many
SELECT *
FROM customers
ORDER BY business_name;

-- name: CreateCustomer :one
INSERT INTO customers (business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, terms, discount, commission, sales_rep, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
SET business_name = $2,
    contact_name = $3,
    email = $4,
    phone = $5,
    address_1 = $6,
    address_2 = $7,
    city = $8,
    state = $9,
    zip_code = $10,
    terms = $11,
    discount = $12,
    commission = $13,
    sales_rep = $14,
    notes = $15,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;
