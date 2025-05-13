-- name: CreateCustomerLocation :one
INSERT INTO customer_locations (
    customer_id, address_1, address_2, city, state, zip_code, phone, notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetCustomerLocationByID :one
SELECT * FROM customer_locations WHERE id = $1;

-- name: ListCustomerLocationsByCustomer :many
SELECT * FROM customer_locations WHERE customer_id = $1 ORDER BY created_at DESC;

-- name: UpdateCustomerLocation :one
UPDATE customer_locations SET
    address_1 = $2,
    address_2 = $3,
    city = $4,
    state = $5,
    zip_code = $6,
    phone = $7,
    notes = $8,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCustomerLocation :exec
DELETE FROM customer_locations WHERE id = $1;
