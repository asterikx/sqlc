// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package override

import (
	"github.com/asterikx/sqlc-testdata/pkg"
)

type Bar struct {
	Other       string
	Total       int64
	AlsoRetyped pkg.CustomType
}

type Baz struct {
	Other       string
	Total       int64
	AlsoRetyped pkg.CustomType
}

type Foo struct {
	Other   string
	Total   int64
	Retyped pkg.CustomType
}
