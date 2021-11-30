// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package authors

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO
  authors (name, bio)
VALUES
  ($1, $2) RETURNING id, name, bio
`

type CreateAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const createAuthors = `-- name: CreateAuthors :many
INSERT INTO
  authors
SELECT
  unnest($1::text[]) AS name,
  unnest($2::text[]) AS bio RETURNING id, name, bio
`

type CreateAuthorsParams struct {
	Names []string
	Bios  []string
}

func (q *Queries) CreateAuthors(ctx context.Context, arg CreateAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, createAuthors, pq.Array(arg.Names), pq.Array(arg.Bios))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM
  authors
WHERE
  id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT
  id, name, bio
FROM
  authors
WHERE
  id = $1
LIMIT
  1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT
  id, name, bio
FROM
  authors
ORDER BY
  name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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
