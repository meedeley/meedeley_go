-- name: FindAllUser :many
SELECT id, name, email, created_at, updated_at FROM users;

-- name: InsertUser :one
INSERT INTO
    users (name, email, password)
VALUES ($1, $2, $3)
RETURNING
    id,
    name,
    email,
    created_at,
    updated_at;

-- name: FindUserByEmail :one
SELECT id, name, email, password,created_at, updated_at FROM users WHERE email = $1;

-- name: FindUserById :one
SELECT
    id,
    name,
    email,
    created_at,
    updated_at
FROM users
WHERE
    id = $1;

-- name: UpdateUserById :exec
UPDATE users
SET
    name = $2,
    email = $3,
    updated_at = $4
WHERE
    id = $1
RETURNING
    id,
    name,
    email,
    created_at,
    updated_at;

-- name: DeleteUserById :exec
DELETE FROM users
WHERE
    id = $1
RETURNING
    id,
    name,
    email,
    created_at,
    updated_at;