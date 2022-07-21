-- name: GetAllFollowers :many
SELECT *
FROM user_follows
WHERE following = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: GetAllFollowing :many
SELECT *
FROM user_follows
WHERE follower = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- namw: FollowUser :one
INSERT INTO user_follows
(id, follower, following)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UnfollowUser :one
DELETE FROM user_follows
WHERE id = $1
RETURNING *;
