// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/asterikx/sqlc/internal/sql/ast"
	"github.com/asterikx/sqlc/internal/sql/catalog"
)

func PgFreespacemap() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = []*catalog.Function{
		{
			Name: "pg_freespace",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "regclass"},
				},
				{
					Type: &ast.TypeName{Name: "bigint"},
				},
			},
			ReturnType: &ast.TypeName{Name: "smallint"},
		},
	}
	return s
}
