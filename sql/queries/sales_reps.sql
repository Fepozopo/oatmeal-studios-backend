-- name: GetSalesRep:one
SELECT *
FROM sales_reps
WHERE id = $1;

-- name: ListSalesReps :many
SELECT *
FROM sales_reps
ORDER BY company;

-- name: CreateSalesRep :one
INSERT INTO sales_reps (status, first_name, last_name, company, address_1, address_2, city, state, zip_code)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateSalesRep :one
UPDATE sales_reps
SET status = $2,
    first_name = $3,
    last_name = $4,
    company = $5,
    address_1 = $6,
    address_2 = $7,
    city = $8,
    state = $9,
    zip_code = $10,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteSalesRep :exec
DELETE FROM sales_reps
WHERE id = $1;