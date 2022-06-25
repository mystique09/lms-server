-- name: GetClassWork :one
--description: Get a class work by id
--parameters: id
--returns: class_work
SELECT * 
FROM class_work 
WHERE id = $1;

-- name: ListClassWork :many
--description: List all class works
--parameters: none
--returns: class_work
SELECT *
FROM class_work;
