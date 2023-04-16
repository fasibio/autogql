package structure

import (
	"github.com/vektah/gqlparser/v2/ast"
)

type Directive string

const (
	DirectiveSQL                 Directive = "SQL"
	DirectiveSQLPrimary          Directive = "SQL_PRIMARY"
	DirectiveSQLIndex            Directive = "SQL_INDEX"
	DirectiveSQLGorm             Directive = "SQL_GORM"
	DirectiveNoMutation          Directive = "SQL_SKIP_MUTATION"
	DirectiveSQLArgumentQuery              = "query"
	DirectiveSQLArgumentMutation           = "mutation"
	DirectiveSQLArgumentOrder              = "order"
)

type SqlBuilderList map[string]*Object

type SqlBuilderHelper struct {
	List SqlBuilderList
}

func (sbh SqlBuilderList) Objects() map[string]Object {
	res := make(map[string]Object)
	for k, v := range sbh {
		if v.Raw.Kind == ast.Object {
			res[k] = *v
		}
	}
	return res
}

func (sbh SqlBuilderList) Enums() map[string]Object {
	res := make(map[string]Object)
	for k, v := range sbh {
		if v.Raw.Kind == ast.Enum {
			res[k] = *v
		}
	}
	return res
}

func (sbh SqlBuilderList) PrimaryEntityOfObject(objectName string) *Entity {
	return sbh.Objects()[objectName].PrimaryKeyField()
}

func NewSqlBuilderHelper() SqlBuilderHelper {
	return SqlBuilderHelper{
		List: make(SqlBuilderList),
	}
}
