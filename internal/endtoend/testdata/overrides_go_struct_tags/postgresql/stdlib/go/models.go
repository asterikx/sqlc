// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package override

import (
	"database/sql"
)

type Foo struct {
	ID      sql.NullString `x:"y"`
	OtherID sql.NullString
}
