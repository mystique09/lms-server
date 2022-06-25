-- name: GetUser :one
SELECT *
FROM "user"
WHERE id = $1;

-- name: GetUsers :many
SELECT (id, username, email, user_role)
FROM "user";

-- name: GetUserByUsername :one
SELECT *
FROM "user"
WHERE username = $1;

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

-- name: UpdateUser :exec
UPDATE "user"
SET password = $1, email = $2
WHERE id =  $3;

-- name: UpdateUserPassword :exec
UPDATE "user"
SET password = $1
WHERE id =  $2;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
