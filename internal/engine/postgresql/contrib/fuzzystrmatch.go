// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/asterikx/sqlc/internal/sql/ast"
	"github.com/asterikx/sqlc/internal/sql/catalog"
)

var funcsFuzzystrmatch = []*catalog.Function{
	{
		Name: "difference",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "dmetaphone",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "dmetaphone_alt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "levenshtein",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "levenshtein",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "levenshtein_less_equal",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "levenshtein_less_equal",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "metaphone",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "soundex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "text_soundex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
}

func Fuzzystrmatch() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsFuzzystrmatch
	return s
}
