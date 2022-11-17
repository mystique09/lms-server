// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: class.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createClass = `-- name: CreateClass :one
INSERT INTO classrooms(
  id, admin_id, name, description, section, room, subject, invite_code, visibility
)
VALUES(
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type CreateClassParams struct {
	ID          uuid.UUID  `json:"id"`
	AdminID     uuid.UUID  `json:"admin_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Section     string     `json:"section"`
	Room        string     `json:"room"`
	Subject     string     `json:"subject"`
	InviteCode  uuid.UUID  `json:"invite_code"`
	Visibility  Visibility `json:"visibility"`
}

func (q *Queries) CreateClass(ctx context.Context, arg CreateClassParams) (Classroom, error) {
	row := q.queryRow(ctx, q.createClassStmt, createClass,
		arg.ID,
		arg.AdminID,
		arg.Name,
		arg.Description,
		arg.Section,
		arg.Room,
		arg.Subject,
		arg.InviteCode,
		arg.Visibility,
	)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteClass = `-- name: DeleteClass :one
DELETE FROM classrooms
WHERE id = $1
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

func (q *Queries) DeleteClass(ctx context.Context, id uuid.UUID) (Classroom, error) {
	row := q.queryRow(ctx, q.deleteClassStmt, deleteClass, id)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllClassFromUser = `-- name: GetAllClassFromUser :many
SELECT id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
FROM classrooms
WHERE admin_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type GetAllClassFromUserParams struct {
	AdminID uuid.UUID `json:"admin_id"`
	Offset  int32     `json:"offset"`
}

func (q *Queries) GetAllClassFromUser(ctx context.Context, arg GetAllClassFromUserParams) ([]Classroom, error) {
	rows, err := q.query(ctx, q.getAllClassFromUserStmt, getAllClassFromUser, arg.AdminID, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Classroom
	for rows.Next() {
		var i Classroom
		if err := rows.Scan(
			&i.ID,
			&i.AdminID,
			&i.Name,
			&i.Description,
			&i.Section,
			&i.Room,
			&i.Subject,
			&i.InviteCode,
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

const getClass = `-- name: GetClass :one
SELECT id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
FROM classrooms
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetClass(ctx context.Context, id uuid.UUID) (Classroom, error) {
	row := q.queryRow(ctx, q.getClassStmt, getClass, id)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getClassroomWithInviteCode = `-- name: GetClassroomWithInviteCode :one
SELECT id
FROM classrooms
WHERE invite_code = $1
LIMIT 1
`

func (q *Queries) GetClassroomWithInviteCode(ctx context.Context, inviteCode uuid.UUID) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.getClassroomWithInviteCodeStmt, getClassroomWithInviteCode, inviteCode)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const listAllPublicClass = `-- name: ListAllPublicClass :many
SELECT id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
FROM classrooms
WHERE visibility = 'PUBLIC'
ORDER BY created_at
LIMIT 10
OFFSET $1
`

func (q *Queries) ListAllPublicClass(ctx context.Context, offset int32) ([]Classroom, error) {
	rows, err := q.query(ctx, q.listAllPublicClassStmt, listAllPublicClass, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Classroom
	for rows.Next() {
		var i Classroom
		if err := rows.Scan(
			&i.ID,
			&i.AdminID,
			&i.Name,
			&i.Description,
			&i.Section,
			&i.Room,
			&i.Subject,
			&i.InviteCode,
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

const updateClassroomDescription = `-- name: UpdateClassroomDescription :one
UPDATE classrooms
SET description = $1
WHERE id = $2
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type UpdateClassroomDescriptionParams struct {
	Description string    `json:"description"`
	ID          uuid.UUID `json:"id"`
}

func (q *Queries) UpdateClassroomDescription(ctx context.Context, arg UpdateClassroomDescriptionParams) (Classroom, error) {
	row := q.queryRow(ctx, q.updateClassroomDescriptionStmt, updateClassroomDescription, arg.Description, arg.ID)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateClassroomInviteCode = `-- name: UpdateClassroomInviteCode :one
UPDATE classrooms
SET invite_code = $1
WHERE id = $2
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type UpdateClassroomInviteCodeParams struct {
	InviteCode uuid.UUID `json:"invite_code"`
	ID         uuid.UUID `json:"id"`
}

func (q *Queries) UpdateClassroomInviteCode(ctx context.Context, arg UpdateClassroomInviteCodeParams) (Classroom, error) {
	row := q.queryRow(ctx, q.updateClassroomInviteCodeStmt, updateClassroomInviteCode, arg.InviteCode, arg.ID)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateClassroomName = `-- name: UpdateClassroomName :one
UPDATE classrooms
SET name = $1
WHERE id = $2
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type UpdateClassroomNameParams struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func (q *Queries) UpdateClassroomName(ctx context.Context, arg UpdateClassroomNameParams) (Classroom, error) {
	row := q.queryRow(ctx, q.updateClassroomNameStmt, updateClassroomName, arg.Name, arg.ID)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateClassroomRoom = `-- name: UpdateClassroomRoom :one
UPDATE classrooms
SET room = $1
WHERE id = $2
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type UpdateClassroomRoomParams struct {
	Room string    `json:"room"`
	ID   uuid.UUID `json:"id"`
}

func (q *Queries) UpdateClassroomRoom(ctx context.Context, arg UpdateClassroomRoomParams) (Classroom, error) {
	row := q.queryRow(ctx, q.updateClassroomRoomStmt, updateClassroomRoom, arg.Room, arg.ID)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateClassroomSection = `-- name: UpdateClassroomSection :one
UPDATE classrooms
SET section = $1
WHERE id = $2
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type UpdateClassroomSectionParams struct {
	Section string    `json:"section"`
	ID      uuid.UUID `json:"id"`
}

func (q *Queries) UpdateClassroomSection(ctx context.Context, arg UpdateClassroomSectionParams) (Classroom, error) {
	row := q.queryRow(ctx, q.updateClassroomSectionStmt, updateClassroomSection, arg.Section, arg.ID)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateClassroomSubject = `-- name: UpdateClassroomSubject :one
UPDATE classrooms
SET subject = $1
WHERE id = $2
RETURNING id, admin_id, name, description, section, room, subject, invite_code, visibility, created_at, updated_at
`

type UpdateClassroomSubjectParams struct {
	Subject string    `json:"subject"`
	ID      uuid.UUID `json:"id"`
}

func (q *Queries) UpdateClassroomSubject(ctx context.Context, arg UpdateClassroomSubjectParams) (Classroom, error) {
	row := q.queryRow(ctx, q.updateClassroomSubjectStmt, updateClassroomSubject, arg.Subject, arg.ID)
	var i Classroom
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Name,
		&i.Description,
		&i.Section,
		&i.Room,
		&i.Subject,
		&i.InviteCode,
		&i.Visibility,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
