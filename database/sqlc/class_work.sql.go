// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: class_work.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const deleteClassworkFromClass = `-- name: DeleteClassworkFromClass :one
DELETE FROM class_works
WHERE id = $1 AND user_id = $2 AND class_id = $3
RETURNING id, url, user_id, class_id, mark, created_at, updated_at
`

type DeleteClassworkFromClassParams struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	ClassID uuid.UUID `json:"class_id"`
}

func (q *Queries) DeleteClassworkFromClass(ctx context.Context, arg DeleteClassworkFromClassParams) (ClassWork, error) {
	row := q.queryRow(ctx, q.deleteClassworkFromClassStmt, deleteClassworkFromClass, arg.ID, arg.UserID, arg.ClassID)
	var i ClassWork
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.ClassID,
		&i.Mark,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getClassWork = `-- name: GetClassWork :one
SELECT id, url, user_id, class_id, mark, created_at, updated_at 
FROM class_works 
WHERE id = $1 AND user_id = $2
LIMIT 1
`

type GetClassWorkParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) GetClassWork(ctx context.Context, arg GetClassWorkParams) (ClassWork, error) {
	row := q.queryRow(ctx, q.getClassWorkStmt, getClassWork, arg.ID, arg.UserID)
	var i ClassWork
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.ClassID,
		&i.Mark,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertNewClasswork = `-- name: InsertNewClasswork :one
INSERT INTO class_works (
  id, url, user_id, class_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, url, user_id, class_id, mark, created_at, updated_at
`

type InsertNewClassworkParams struct {
	ID      uuid.UUID `json:"id"`
	Url     string    `json:"url"`
	UserID  uuid.UUID `json:"user_id"`
	ClassID uuid.UUID `json:"class_id"`
}

func (q *Queries) InsertNewClasswork(ctx context.Context, arg InsertNewClassworkParams) (ClassWork, error) {
	row := q.queryRow(ctx, q.insertNewClassworkStmt, insertNewClasswork,
		arg.ID,
		arg.Url,
		arg.UserID,
		arg.ClassID,
	)
	var i ClassWork
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.ClassID,
		&i.Mark,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listClassworkAdmin = `-- name: ListClassworkAdmin :many
SELECT id, url, user_id, class_id, mark, created_at, updated_at
FROM class_works
WHERE class_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type ListClassworkAdminParams struct {
	ClassID uuid.UUID `json:"class_id"`
	Offset  int32     `json:"offset"`
}

func (q *Queries) ListClassworkAdmin(ctx context.Context, arg ListClassworkAdminParams) ([]ClassWork, error) {
	rows, err := q.query(ctx, q.listClassworkAdminStmt, listClassworkAdmin, arg.ClassID, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClassWork
	for rows.Next() {
		var i ClassWork
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.UserID,
			&i.ClassID,
			&i.Mark,
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

const listSubmittedClassworks = `-- name: ListSubmittedClassworks :many
SELECT id, url, user_id, class_id, mark, created_at, updated_at
FROM class_works
WHERE user_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type ListSubmittedClassworksParams struct {
	UserID uuid.UUID `json:"user_id"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListSubmittedClassworks(ctx context.Context, arg ListSubmittedClassworksParams) ([]ClassWork, error) {
	rows, err := q.query(ctx, q.listSubmittedClassworksStmt, listSubmittedClassworks, arg.UserID, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClassWork
	for rows.Next() {
		var i ClassWork
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.UserID,
			&i.ClassID,
			&i.Mark,
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

const updateAClassworkMark = `-- name: UpdateAClassworkMark :exec
UPDATE class_works
SET mark = $1
WHERE id = $2 AND user_id = $3 AND class_id = $4
`

type UpdateAClassworkMarkParams struct {
	Mark    sql.NullInt32 `json:"mark"`
	ID      uuid.UUID     `json:"id"`
	UserID  uuid.UUID     `json:"user_id"`
	ClassID uuid.UUID     `json:"class_id"`
}

func (q *Queries) UpdateAClassworkMark(ctx context.Context, arg UpdateAClassworkMarkParams) error {
	_, err := q.exec(ctx, q.updateAClassworkMarkStmt, updateAClassworkMark,
		arg.Mark,
		arg.ID,
		arg.UserID,
		arg.ClassID,
	)
	return err
}
