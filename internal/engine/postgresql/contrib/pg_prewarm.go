// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/asterikx/sqlc/internal/sql/ast"
	"github.com/asterikx/sqlc/internal/sql/catalog"
)

var funcsPgPrewarm = []*catalog.Function{
	{
		Name:       "autoprewarm_dump_now",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "bigint"},
	},
	{
		Name:       "autoprewarm_start_worker",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "pg_prewarm",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name:       "mode",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "text"},
			},
			{
				Name:       "fork",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "text"},
			},
			{
				Name:       "first_block",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "bigint"},
			},
			{
				Name:       "last_block",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bigint"},
	},
}

func PgPrewarm() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPgPrewarm
	return s
}
