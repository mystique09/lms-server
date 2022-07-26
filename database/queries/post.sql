-- name: GetOnePost :one
SELECT *
FROM posts
WHERE id = $1
LIMIT 1;

-- name: ListAllPostsFromClass :many
SELECT *
FROM posts
WHERE class_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: ListAllPostsByUser :many
SELECT *
FROM posts
WHERE author_id = $1 AND class_id = $2
ORDER BY created_at;


-- name: InsertNewPost :one
INSERT INTO posts (
  id, content, author_id, class_id
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: UpdatePostContent :one
UPDATE posts
SET content = $1
WHERE id = $2 AND author_id = $3 AND class_id = $4
RETURNING (id, content, author_id);

-- name: DeletePostFromClass :one
DELETE FROM posts
WHERE id = $1 AND author_id = $2 AND class_id = $3
RETURNING (id, content, author_id);
