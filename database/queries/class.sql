--name: GetClass :one
--description: Get a class by id
--parameters: id
--returns: class
SELECT *
FROM class
WHERE id = $1;

--name: ListClass :many
--description: List all classes
--parameters: none
--returns: classes
SELECT *
FROM class;

--name: CreateClass :one
--description: Create a class
--parameters: id(uuid), admin_id, name, description, section, room, subject, invite_code, created_at, updated_at
--returns: class
INSERT INTO class (id, admin_id, name, description, section, room, subject, invite_code, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

--name: UpdateClass :exec
--description: Update a class
--parameters: name, description, section, room, subject, invite_code, updated_at
UPDATE class
SET name = $1, description = $2, section = $3, room = $4, subject = $5, invite_code = $6, updated_at = $7
WHERE id = $8;

--name: DeleteClass :exec
--description: Delete a class
--parameters: id
DELETE FROM class
WHERE id = $1;
