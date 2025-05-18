-- name: GetUserByID :one
SELECT id, created_at, updated_at, email, first_name, last_name, password
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, first_name, last_name, password
FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT id, created_at, updated_at, email, first_name, last_name
FROM users
ORDER BY last_name, first_name;

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password, first_name, last_name)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: UpdateUserName :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;
