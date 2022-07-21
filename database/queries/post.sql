-- name: GetOnePost :one
SELECT *
FROM "post"
WHERE id = $1 AND class_id = $2
LIMIT 1;

-- name: ListAllPostsFromClass :many
SELECT *
FROM "post"
WHERE class_id = $1
ORDER BY created_at
ASC;

-- name: ListAllPostsByUser :many
SELECT *
FROM "post"
WHERE author_id = $1 AND class_id = $2
ORDER BY created_at
ASC;

-- name: InsertNewPost :one
INSERT INTO "post" (
  id, content, author_id, class_id
) VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: UpdatePostContent :one
UPDATE "post"
SET content = $1
WHERE id = $2 AND author_id = $3 AND class_id = $4
RETURNING (id, content, author_id);

-- name: DeletePostFromClass :one
DELETE FROM "post"
WHERE id = $1 AND author_id = $2 AND class_id = $3
RETURNING (id, content, author_id);

-- name: GetAllCommentsFromPost :many
SELECT *
FROM "comment"
WHERE post_id = $1
ORDER BY created_at
ASC;

-- name: InsertNewCommentInPost :one
INSERT INTO "comment" (
  id, content, author_id, post_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateCommentContentInPost :one
UPDATE "comment"
SET content = $1
WHERE id = $2 AND author_id = $3 AND post_id = $4
RETURNING (id, content, author_id);

-- name: DeleteCommentFromPost :one
DELETE FROM "comment"
WHERE id = $1 AND author_id = $2 AND post_id = $3
RETURNING (id, content, author_id);
