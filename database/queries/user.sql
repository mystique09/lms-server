-- name: GetUser :one
SELECT *
FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUsers :many
--description: Get all user by page and offset
SELECT *
FROM "user"
ORDER BY created_at
ASC;

-- name: GetUserByUsername :one
SELECT *
FROM "user"
WHERE username = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user"(
  id, username, password, email, user_role, visibility
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUsername :one
UPDATE "user"
SET username = $1
WHERE id =  $2
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE "user"
SET email = $1
WHERE id =  $2
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE "user"
SET password = $1
WHERE id =  $2
RETURNING *;

-- name: DeleteUser :one
DELETE FROM "user"
WHERE id = $1
RETURNING *;
