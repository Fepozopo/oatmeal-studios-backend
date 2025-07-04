-- name: GetCustomer :one
SELECT *
FROM customers
WHERE id = $1;

-- name: ListCustomers :many
SELECT *
FROM customers
ORDER BY business_name;

-- name: CreateCustomer :one
INSERT INTO customers (business_name, contact_name, email, phone, address_1, address_2, city, state, zip_code, country, terms, discount, commission, notes, free_shipping)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
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
    country = $11,
    terms = $12,
    discount = $13,
    commission = $14,
    notes = $15,
    free_shipping = $16,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;
