// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package batch

import (
	"context"

	"github.com/jackc/pgconn"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name) VALUES ($1)
RETURNING author_id, name, biography
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRow(ctx, createAuthor, name)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name, &i.Biography)
	return i, err
}

const deleteBookExecResult = `-- name: DeleteBookExecResult :execresult
DELETE FROM books
WHERE book_id = $1
`

func (q *Queries) DeleteBookExecResult(ctx context.Context, bookID int32) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteBookExecResult, bookID)
}

const getAuthor = `-- name: GetAuthor :one
SELECT author_id, name, biography FROM authors
WHERE author_id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, authorID int32) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, authorID)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name, &i.Biography)
	return i, err
}
