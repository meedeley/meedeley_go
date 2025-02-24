-- name: FindAllExample :many
SELECT id, name, description, created_at, updated_at FROM examples;

-- name: FindExampleById :many