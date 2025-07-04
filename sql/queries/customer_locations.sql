-- name: CreateCustomerLocation :one
INSERT INTO customer_locations (
    customer_id, business_name, contact_name, address_1, address_2, city, state, zip_code, country, phone, sales_rep, notes, location_num
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: GetCustomerLocationByID :one
SELECT * FROM customer_locations WHERE id = $1;

-- name: ListCustomerLocationsByCustomer :many
SELECT * FROM customer_locations WHERE customer_id = $1 ORDER BY created_at DESC;

-- name: UpdateCustomerLocation :one
UPDATE customer_locations SET
    business_name = $2,
    contact_name = $3,
    address_1 = $4,
    address_2 = $5,
    city = $6,
    state = $7,
    zip_code = $8,
    country = $9,
    phone = $10,
    sales_rep = $11,
    notes = $12,
    location_num = $13,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCustomerLocation :exec
DELETE FROM customer_locations WHERE id = $1;
