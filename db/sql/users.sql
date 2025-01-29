-- name: GetAllUsers :many
SELECT id, name, email
FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3);
