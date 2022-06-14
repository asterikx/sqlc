// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package querytest

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func New() *Queries {
	return &Queries{}
}

type Queries struct {
}
