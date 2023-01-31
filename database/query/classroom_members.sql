-- name: GetAllClassroomMembers :many
SELECT *
FROM classroom_members
WHERE class_id = $1
ORDER BY created_at
LIMIT 100
OFFSET $2;

-- name: GetClassroomMemberById :one
SELECT *
FROM classroom_members
WHERE user_id = $1 AND class_id = $2
LIMIT 1;

-- name: GetAllJoinedClassrooms :many
SELECT c.*
FROM classrooms c
WHERE c.id = (SELECT cm.class_id FROM classroom_members cm WHERE cm.class_id = c.id AND cm.user_id = $1)
ORDER BY c.created_at
LIMIT 10
OFFSET $2;

-- name: AddNewClassroomMember :one
INSERT INTO classroom_members
(
  id, class_id, user_id
)
VALUES
(
  $1, $2, $3
)
RETURNING *;

-- name: LeaveClassroom :one
DELETE FROM classroom_members
WHERE user_id = $1 AND class_id = $2
RETURNING *;