-- name: GetOneFollower :one
SELECT *
FROM user_follows
WHERE follower = $1 AND following = $2
LIMIT 1;

-- name: GetFollowerById :one
SELECT id, follower AS user_id, created_at, updated_at
FROM user_follows
WHERE id = $1
LIMIT 1;

-- name: GetAllFollowers :many
SELECT id, follower AS user_id, created_at, updated_at
FROM user_follows
WHERE following = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: GetAllFollowing :many
SELECT id, following AS user_id, created_at, updated_at
FROM user_follows
WHERE follower = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: FollowUser :one
INSERT INTO user_follows
(id, follower, following)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UnfollowUser :one
DELETE FROM user_follows
WHERE id = $1
RETURNING *;
