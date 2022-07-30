// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: comment.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteCommentFromPost = `-- name: DeleteCommentFromPost :one
DELETE FROM comments
WHERE id = $1 AND author_id = $2 AND post_id = $3
RETURNING (id, content, author_id)
`

type DeleteCommentFromPostParams struct {
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}

func (q *Queries) DeleteCommentFromPost(ctx context.Context, arg DeleteCommentFromPostParams) (interface{}, error) {
	row := q.queryRow(ctx, q.deleteCommentFromPostStmt, deleteCommentFromPost, arg.ID, arg.AuthorID, arg.PostID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getAllCommentLikes = `-- name: GetAllCommentLikes :many
SELECT id, comment_id, user_id, created_at, updated_at
FROM comment_likes
WHERE comment_id = $1
`

func (q *Queries) GetAllCommentLikes(ctx context.Context, commentID uuid.UUID) ([]CommentLike, error) {
	rows, err := q.query(ctx, q.getAllCommentLikesStmt, getAllCommentLikes, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CommentLike
	for rows.Next() {
		var i CommentLike
		if err := rows.Scan(
			&i.ID,
			&i.CommentID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllCommentsFromPost = `-- name: GetAllCommentsFromPost :many
SELECT id, content, author_id, post_id, created_at, updated_at
FROM comments
WHERE post_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type GetAllCommentsFromPostParams struct {
	PostID uuid.UUID `json:"post_id"`
	Offset int32     `json:"offset"`
}

func (q *Queries) GetAllCommentsFromPost(ctx context.Context, arg GetAllCommentsFromPostParams) ([]Comment, error) {
	rows, err := q.query(ctx, q.getAllCommentsFromPostStmt, getAllCommentsFromPost, arg.PostID, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.PostID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertNewCommentInPost = `-- name: InsertNewCommentInPost :one
INSERT INTO comments (
  id, content, author_id, post_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, content, author_id, post_id, created_at, updated_at
`

type InsertNewCommentInPostParams struct {
	ID       uuid.UUID `json:"id"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}

func (q *Queries) InsertNewCommentInPost(ctx context.Context, arg InsertNewCommentInPostParams) (Comment, error) {
	row := q.queryRow(ctx, q.insertNewCommentInPostStmt, insertNewCommentInPost,
		arg.ID,
		arg.Content,
		arg.AuthorID,
		arg.PostID,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.PostID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const likeComment = `-- name: LikeComment :one
INSERT INTO comment_likes (
  id, comment_id, user_id
) VALUES (
$1, $2, $3
) RETURNING id, comment_id, user_id, created_at, updated_at
`

type LikeCommentParams struct {
	ID        uuid.UUID `json:"id"`
	CommentID uuid.UUID `json:"comment_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func (q *Queries) LikeComment(ctx context.Context, arg LikeCommentParams) (CommentLike, error) {
	row := q.queryRow(ctx, q.likeCommentStmt, likeComment, arg.ID, arg.CommentID, arg.UserID)
	var i CommentLike
	err := row.Scan(
		&i.ID,
		&i.CommentID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const unlikeComment = `-- name: UnlikeComment :one
DELETE FROM comment_likes
WHERE id = $1 AND comment_id = $2
RETURNING id, comment_id, user_id, created_at, updated_at
`

type UnlikeCommentParams struct {
	ID        uuid.UUID `json:"id"`
	CommentID uuid.UUID `json:"comment_id"`
}

func (q *Queries) UnlikeComment(ctx context.Context, arg UnlikeCommentParams) (CommentLike, error) {
	row := q.queryRow(ctx, q.unlikeCommentStmt, unlikeComment, arg.ID, arg.CommentID)
	var i CommentLike
	err := row.Scan(
		&i.ID,
		&i.CommentID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCommentContentInPost = `-- name: UpdateCommentContentInPost :one
UPDATE comments
SET content = $1
WHERE id = $2 AND author_id = $3 AND post_id = $4
RETURNING (id, content, author_id)
`

type UpdateCommentContentInPostParams struct {
	Content  string    `json:"content"`
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}

func (q *Queries) UpdateCommentContentInPost(ctx context.Context, arg UpdateCommentContentInPostParams) (interface{}, error) {
	row := q.queryRow(ctx, q.updateCommentContentInPostStmt, updateCommentContentInPost,
		arg.Content,
		arg.ID,
		arg.AuthorID,
		arg.PostID,
	)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}
