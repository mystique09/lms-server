-- name: GetClassWork :one
--description: Get a class work by id
--parameters: id
--returns: class_work
SELECT * 
FROM class_work 
WHERE id = $1, class_id = $2;

-- name: ListClassworkAdmin :many
--description: List all class works
--parameters: none
--returns: class_work
SELECT *
FROM class_work
WHERE class_id = $1
ORDER BY created_at
DESC;

-- name: ListSubmittedClassworks :many
--description: List all submitted classworks of a user
--parameters: user_id, class_id
--returns: class_work
SELECT *
FROM class_work
WHERE class_id = $1, user_id = $2
ORDER BY created_at
DESC;
