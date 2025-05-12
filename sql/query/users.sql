-- name: GetUser :one
SELECT id, created_at, updated_at, email, first_name, last_name
FROM users
WHERE id = $1;

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
RETURNING id, created_at, updated_at, email, first_name, last_name;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    first_name = $3,
    last_name = $4,
    password = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, email, first_name, last_name;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserFromRefreshToken :one
SELECT users.id, users.created_at, users.updated_at, users.email, users.first_name, users.last_name
FROM users
    INNER JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1;
