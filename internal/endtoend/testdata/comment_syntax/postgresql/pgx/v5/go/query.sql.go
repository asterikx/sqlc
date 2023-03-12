// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const doubleDash = `-- name: DoubleDash :one
SELECT bar FROM foo LIMIT 1
`

func (q *Queries) DoubleDash(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, doubleDash)
	var bar pgtype.Text
	err := row.Scan(&bar)
	return bar, err
}

const slashStar = `-- name: SlashStar :one
SELECT bar FROM foo LIMIT 1
`

func (q *Queries) SlashStar(ctx context.Context) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, slashStar)
	var bar pgtype.Text
	err := row.Scan(&bar)
	return bar, err
}
