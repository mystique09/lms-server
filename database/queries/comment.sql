-- name: GetAllCommentsFromPost :many
SELECT *
FROM comments
WHERE post_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2;

-- name: InsertNewCommentInPost :one
INSERT INTO comments (
  id, content, author_id, post_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateCommentContentInPost :one
UPDATE comments
SET content = $1
WHERE id = $2 AND author_id = $3 AND post_id = $4
RETURNING (id, content, author_id);

-- name: DeleteCommentFromPost :one
DELETE FROM comments
WHERE id = $1 AND author_id = $2 AND post_id = $3
RETURNING (id, content, author_id);

-- name: GetAllCommentLikes :many
SELECT *
FROM comment_likes
WHERE comment_id = $1;

-- name: LikeComment :one
INSERT INTO comment_likes (
  id, comment_id, user_id
) VALUES (
$1, $2, $3
) RETURNING *;

-- name: UnlikeComment :one
DELETE FROM comment_likes
WHERE id = $1 AND comment_id = $2
RETURNING *;
