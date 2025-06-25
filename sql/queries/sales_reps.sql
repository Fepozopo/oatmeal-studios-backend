-- name: GetSalesRep :one
SELECT *
FROM sales_reps
WHERE id = $1;

-- name: ListSalesReps :many
SELECT *
FROM sales_reps
ORDER BY company;

-- name: CreateSalesRep :one
INSERT INTO sales_reps (status, rep_code, first_name, last_name, company, address_1, address_2, city, state, country, zip_code, phone, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW(), NOW())
RETURNING *;

-- name: UpdateSalesRep :one
UPDATE sales_reps
SET status = $2,
    rep_code = $13,
    first_name = $3,
    last_name = $4,
    company = $5,
    address_1 = $6,
    address_2 = $7,
    city = $8,
    state = $9,
    zip_code = $10,
    country = $11,
    phone = $12,
    email = $14,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteSalesRep :exec
DELETE FROM sales_reps
WHERE id = $1;