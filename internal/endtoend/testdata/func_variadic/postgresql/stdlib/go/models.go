// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package querytest

import (
	"database/sql"
	"encoding/json"
)

type Test struct {
	ID sql.NullInt32
	J  json.RawMessage
}
