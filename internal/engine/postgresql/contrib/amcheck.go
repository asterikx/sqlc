// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/asterikx/sqlc/internal/sql/ast"
	"github.com/asterikx/sqlc/internal/sql/catalog"
)

var funcsAmcheck = []*catalog.Function{
	{
		Name: "bt_index_check",
		Args: []*catalog.Argument{
			{
				Name: "index",
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "bt_index_check",
		Args: []*catalog.Argument{
			{
				Name: "index",
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "heapallindexed",
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "bt_index_parent_check",
		Args: []*catalog.Argument{
			{
				Name: "index",
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "bt_index_parent_check",
		Args: []*catalog.Argument{
			{
				Name: "index",
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "heapallindexed",
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "bt_index_parent_check",
		Args: []*catalog.Argument{
			{
				Name: "index",
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "heapallindexed",
				Type: &ast.TypeName{Name: "boolean"},
			},
			{
				Name: "rootdescend",
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "verify_heapam",
		Args: []*catalog.Argument{
			{
				Name: "relation",
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name:       "on_error_stop",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "boolean"},
			},
			{
				Name:       "check_toast",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "boolean"},
			},
			{
				Name:       "skip",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "text"},
			},
			{
				Name:       "startblock",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "bigint"},
			},
			{
				Name:       "endblock",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
}

func Amcheck() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsAmcheck
	return s
}
