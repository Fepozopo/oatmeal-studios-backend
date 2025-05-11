-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY last_name, first_name;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, first_name, last_name, email;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2,
    last_name = $3,
    email = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, first_name, last_name, email;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
