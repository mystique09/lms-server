-- name: GetClassWork :one
--description: Get a class work by id
--parameters: id
--returns: class_work
SELECT * 
FROM "class_work" 
WHERE id = $1, user_id = $2, class_id = $3;

-- name: ListClassworkAdmin :many
--description: List all class works
--parameters: none
--returns: class_work
SELECT *
FROM "class_work"
WHERE class_id = $2
ORDER BY created_at
DESC;

-- name: ListSubmittedClassworks :many
--description: List all submitted classworks of a user
--parameters: user_id, class_id
--returns: class_work
SELECT *
FROM "class_work"
WHERE class_id = $1, user_id = $2
ORDER BY created_at
DESC;

-- name: InsertNewClasswork :one
INSERT INTO "class_work" (
  id, name, user_id, class_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateAClassworkMark :exec
UPDATE "class_work"
SET mark = $1
WHERE id = $1, user_id = $2, class_id = $3;

-- name: DeleteClassworkFromClass :exec
DELETE FROM "class_work"
WHERE id = $1, user_id = $2, class_id = $3;
