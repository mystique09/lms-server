-- name: GetClassWork :one
SELECT * 
FROM class_works 
WHERE id = $1 AND user_id = $2
LIMIT 1;

-- name: ListClassworkAdmin :many
SELECT *
FROM class_works
WHERE class_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: ListSubmittedClassworks :many
SELECT *
FROM class_works
WHERE user_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: InsertNewClasswork :one
INSERT INTO class_works (
  id, url, user_id, class_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateAClassworkMark :exec
UPDATE class_works
SET mark = $1
WHERE id = $2 AND user_id = $3 AND class_id = $4;

-- name: DeleteClassworkFromClass :exec
DELETE FROM class_works
WHERE id = $1 AND user_id = $2 AND class_id = $3;
