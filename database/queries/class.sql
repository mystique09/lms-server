-- name: GetClass :one
SELECT *
FROM classrooms
WHERE id = $1
LIMIT 1;

-- name: ListAllPublicClass :many
SELECT *
FROM classrooms
WHERE visibility = 'PUBLIC'
ORDER BY created_at
LIMIT 10
OFFSET $1;

-- name: GetAllClassFromUser :many
SELECT *
FROM classrooms
WHERE admin_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: CreateClass :one
INSERT INTO classrooms(
  id, admin_id, name, description, section, room, subject, invite_code, visibility
)
VALUES(
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetClassroomWithInviteCode :one
SELECT id
FROM classrooms
WHERE invite_code = $1
LIMIT 1;

-- name: UpdateClassroomName :one
UPDATE classrooms
SET name = $1
WHERE id = $2
RETURNING *;

-- name: UpdateClassroomDescription :one
UPDATE classrooms
SET description = $1
WHERE id = $2
RETURNING *;

-- name: UpdateClassroomSubject :one
UPDATE classrooms
SET subject = $1
WHERE id = $2
RETURNING *;

-- name: UpdateClassroomSection :one
UPDATE classrooms
SET section = $1
WHERE id = $2
RETURNING *;

-- name: UpdateClassroomRoom :one
UPDATE classrooms
SET room = $1
WHERE id = $2
RETURNING *;

-- name: UpdateClassroomInviteCode :one
UPDATE classrooms
SET invite_code = $1
WHERE id = $2
RETURNING *;

-- name: DeleteClass :one
DELETE FROM classrooms
WHERE id = $1
RETURNING *;
