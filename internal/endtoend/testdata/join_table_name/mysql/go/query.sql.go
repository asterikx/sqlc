// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"
)

const tableName = `-- name: TableName :one
SELECT foo.id
FROM foo
JOIN bar ON foo.bar = bar.id
WHERE bar.id = ? AND foo.id = ?
`

type TableNameParams struct {
	ID   int64
	ID_2 int64
}

func (q *Queries) TableName(ctx context.Context, arg TableNameParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, tableName, arg.ID, arg.ID_2)
	var id int64
	err := row.Scan(&id)
	return id, err
}
