// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package users

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name, email)
VALUES ($1, $2)
`

type CreateUserParams struct {
	Name  string
	Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Name, arg.Email)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, email
FROM users
`

type GetAllUsersRow struct {
	ID    int32
	Name  string
	Email string
}

func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
