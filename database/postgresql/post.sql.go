// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: post.sql

package postgresql

import (
	"context"

	"github.com/google/uuid"
)

const deletePostFromClass = `-- name: DeletePostFromClass :one
DELETE FROM posts
WHERE id = $1 AND author_id = $2 AND class_id = $3
RETURNING (id, content, author_id)
`

type DeletePostFromClassParams struct {
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) DeletePostFromClass(ctx context.Context, arg DeletePostFromClassParams) (interface{}, error) {
	row := q.queryRow(ctx, q.deletePostFromClassStmt, deletePostFromClass, arg.ID, arg.AuthorID, arg.ClassID)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getAllPostLikes = `-- name: GetAllPostLikes :many
SELECT id, post_id, user_id, created_at, updated_at
FROM post_likes
WHERE post_id = $1
`

func (q *Queries) GetAllPostLikes(ctx context.Context, postID uuid.UUID) ([]PostLike, error) {
	rows, err := q.query(ctx, q.getAllPostLikesStmt, getAllPostLikes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostLike
	for rows.Next() {
		var i PostLike
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
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

const getOnePost = `-- name: GetOnePost :one
SELECT id, content, author_id, class_id, created_at, updated_at
FROM posts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetOnePost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.queryRow(ctx, q.getOnePostStmt, getOnePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.ClassID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertNewPost = `-- name: InsertNewPost :one
INSERT INTO posts (
  id, content, author_id, class_id
) VALUES ( $1, $2, $3, $4 )
RETURNING id, content, author_id, class_id, created_at, updated_at
`

type InsertNewPostParams struct {
	ID       uuid.UUID `json:"id"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) InsertNewPost(ctx context.Context, arg InsertNewPostParams) (Post, error) {
	row := q.queryRow(ctx, q.insertNewPostStmt, insertNewPost,
		arg.ID,
		arg.Content,
		arg.AuthorID,
		arg.ClassID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.ClassID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const likePost = `-- name: LikePost :one
INSERT INTO post_likes (
  id, post_id, user_id
) VALUES (
$1, $2, $3
) RETURNING id, post_id, user_id, created_at, updated_at
`

type LikePostParams struct {
	ID     uuid.UUID `json:"id"`
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) LikePost(ctx context.Context, arg LikePostParams) (PostLike, error) {
	row := q.queryRow(ctx, q.likePostStmt, likePost, arg.ID, arg.PostID, arg.UserID)
	var i PostLike
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllPostsByUser = `-- name: ListAllPostsByUser :many
SELECT id, content, author_id, class_id, created_at, updated_at
FROM posts
WHERE author_id = $1 AND class_id = $2
ORDER BY created_at
`

type ListAllPostsByUserParams struct {
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) ListAllPostsByUser(ctx context.Context, arg ListAllPostsByUserParams) ([]Post, error) {
	rows, err := q.query(ctx, q.listAllPostsByUserStmt, listAllPostsByUser, arg.AuthorID, arg.ClassID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.ClassID,
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

const listAllPostsFromClass = `-- name: ListAllPostsFromClass :many
SELECT id, content, author_id, class_id, created_at, updated_at
FROM posts
WHERE class_id = $1
ORDER BY created_at
LIMIT 10
OFFSET $2
`

type ListAllPostsFromClassParams struct {
	ClassID uuid.UUID `json:"class_id"`
	Offset  int32     `json:"offset"`
}

func (q *Queries) ListAllPostsFromClass(ctx context.Context, arg ListAllPostsFromClassParams) ([]Post, error) {
	rows, err := q.query(ctx, q.listAllPostsFromClassStmt, listAllPostsFromClass, arg.ClassID, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.ClassID,
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

const unlikePost = `-- name: UnlikePost :one
DELETE FROM post_likes
WHERE id = $1 AND post_id = $2 AND user_id = $3
RETURNING id, post_id, user_id, created_at, updated_at
`

type UnlikePostParams struct {
	ID     uuid.UUID `json:"id"`
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) UnlikePost(ctx context.Context, arg UnlikePostParams) (PostLike, error) {
	row := q.queryRow(ctx, q.unlikePostStmt, unlikePost, arg.ID, arg.PostID, arg.UserID)
	var i PostLike
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePostContent = `-- name: UpdatePostContent :one
UPDATE posts
SET content = $1
WHERE id = $2 AND author_id = $3 AND class_id = $4
RETURNING (id, content, author_id)
`

type UpdatePostContentParams struct {
	Content  string    `json:"content"`
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	ClassID  uuid.UUID `json:"class_id"`
}

func (q *Queries) UpdatePostContent(ctx context.Context, arg UpdatePostContentParams) (interface{}, error) {
	row := q.queryRow(ctx, q.updatePostContentStmt, updatePostContent,
		arg.Content,
		arg.ID,
		arg.AuthorID,
		arg.ClassID,
	)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}
