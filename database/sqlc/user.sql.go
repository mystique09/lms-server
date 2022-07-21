// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user"(
  id, username, password, email, user_role, visibility
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, username, password, email, user_role, visibility, created_at, updated_at
`

type CreateUserParams struct {
	ID         uuid.UUID  `json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	UserRole   Role       `json:"user_role"`
	Visibility Visibility `json:"visibility"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.UserRole,
		arg.Visibility,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM "user"
WHERE id = $1
RETURNING id, username, password, email, user_role, visibility, created_at, updated_at
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.deleteUserStmt, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, email, user_role, visibility, created_at, updated_at
FROM "user"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, email, user_role, visibility, created_at, updated_at
FROM "user"
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserByUsernameStmt, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, username, password, email, user_role, visibility, created_at, updated_at
FROM "user"
ORDER BY created_at
ASC
`

//description: Get all user by page and offset
func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.getUsersStmt, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Email,
			&i.UserRole,
			&i.Visibility,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserEmail = `-- name: UpdateUserEmail :one
UPDATE "user"
SET email = $1
WHERE id =  $2
RETURNING id, username, password, email, user_role, visibility, created_at, updated_at
`

type UpdateUserEmailParams struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserEmailStmt, updateUserEmail, arg.Email, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
UPDATE "user"
SET password = $1
WHERE id =  $2
RETURNING id, username, password, email, user_role, visibility, created_at, updated_at
`

type UpdateUserPasswordParams struct {
	Password string    `json:"password"`
	ID       uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserPasswordStmt, updateUserPassword, arg.Password, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUsername = `-- name: UpdateUsername :one
UPDATE "user"
SET username = $1
WHERE id =  $2
RETURNING id, username, password, email, user_role, visibility, created_at, updated_at
`

type UpdateUsernameParams struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (User, error) {
	row := q.queryRow(ctx, q.updateUsernameStmt, updateUsername, arg.Username, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.UserRole,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
