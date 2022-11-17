-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUsers :many
SELECT *
FROM users
ORDER BY created_at
ASC
LIMIT 10
OFFSET $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(
  id, username, password, email, user_role, visibility
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING username;

-- name: UpdateUsername :one
UPDATE users
SET username = $1
WHERE id =  $2
RETURNING id;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $1
WHERE id =  $2
RETURNING id;

-- name: UpdateUserPassword :one
UPDATE users
SET password = $1
WHERE id =  $2
RETURNING id;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;
