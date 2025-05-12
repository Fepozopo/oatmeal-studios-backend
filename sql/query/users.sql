-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY last_name, first_name;

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2,
    last_name = $3,
    password = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, first_name, last_name;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserFromRefreshToken :one
SELECT *
FROM users
    INNER JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1;
