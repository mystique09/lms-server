-- name: GetUser :one
SELECT *
FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUsers :many
SELECT *
FROM "user"
ORDER BY created_at DESC;

-- name: GetUserByUsername :one
SELECT *
FROM "user"
WHERE username = $1
LIMIT 1;

-- name: GetUserWithPosts :one
SELECT (id, username, email, user_role)
FROM "user"
LEFT JOIN "post" ON "user".id = "post".user_id;

-- name: CreateUser :one
INSERT INTO "user"(
  id, username, password, email, user_role
) VALUES (
  $1, $2, $3, $4, $5
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
