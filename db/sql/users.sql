-- name: FindAllUser :many
SELECT id, name, email
FROM users;

-- name: InsertUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email;

-- name: FindUserById :one
SELECT id, name, email 
FROM users 
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT id, name, email, password FROM users WHERE email = $1;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1 RETURNING id, name, email, created_at, updated_at;
