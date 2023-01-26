package structure

import (
	"github.com/fasibio/autogql/helper"
	"github.com/vektah/gqlparser/v2/ast"
)

type SQLDirective struct {
	Query    SQLDirectiveQuery
	Mutation SQLDirectiveMutation
}

func (sd *SQLDirective) HasQueries() bool {
	return sd.Query.Query.Value || sd.Query.Get.Value
}

func (sd *SQLDirective) HasMutation() bool {
	return sd.Mutation.Add.Value || sd.Mutation.Delete.Value || sd.Mutation.Update.Value
}

type SqlCreateExtension struct {
	Value        bool     `json:"value,omitempty"`
	DirectiveExt []string `json:"directiveExt,omitempty"`
}

type SQLDirectiveHandler struct {
	DirectiveExt []string
}

type SQLDirectiveMutation struct {
	SQLDirectiveHandler
	Add    SqlCreateExtension
	Update SqlCreateExtension
	Delete SqlCreateExtension
}

type SQLDirectiveQuery struct {
	SQLDirectiveHandler
	Get   SqlCreateExtension
	Query SqlCreateExtension
}

func (s *SQLDirectiveQuery) GetDirectiveExt(name string) []string {
	res := s.DirectiveExt
	switch name {
	case "Get":
		{
			res = append(res, s.Get.DirectiveExt...)
			break
		}
	case "Query":
		{
			res = append(res, s.Query.DirectiveExt...)
			break
		}
	}
	return res
}

func (s *SQLDirectiveMutation) GetDirectiveExt(name string) []string {
	res := s.DirectiveExt
	switch name {
	case "Add":
		{
			res = append(res, s.Add.DirectiveExt...)
			break
		}
	case "Update":
		{
			res = append(res, s.Update.DirectiveExt...)
			break
		}
	case "Delete":
		{
			res = append(res, s.Delete.DirectiveExt...)
			break
		}
	}
	return res
}

func getDefaultFilledSqlBuilderMutation(defaultValue bool) SQLDirectiveMutation {
	defaultV := SqlCreateExtension{
		Value:        defaultValue,
		DirectiveExt: []string{},
	}
	return SQLDirectiveMutation{
		Add:                 defaultV,
		SQLDirectiveHandler: SQLDirectiveHandler{DirectiveExt: []string{}},
		Update:              defaultV,
		Delete:              defaultV,
	}
}

func getDefaultFilledSqlBuilderQuery(defaultValue bool) SQLDirectiveQuery {
	defaultV := SqlCreateExtension{
		Value:        defaultValue,
		DirectiveExt: []string{},
	}
	return SQLDirectiveQuery{
		Get:                 defaultV,
		SQLDirectiveHandler: SQLDirectiveHandler{DirectiveExt: []string{}},
		Query:               defaultV,
	}
}

func getSqlCreateExtension(v map[string]interface{}) SqlCreateExtension {
	return SqlCreateExtension{
		Value:        v["value"].(bool),
		DirectiveExt: helper.GetArrayOfInterface[string](v["directiveExt"]),
	}
}

func customizeSqlBuilderQuery(a *ast.Argument) SQLDirectiveQuery {
	res := getDefaultFilledSqlBuilderQuery(true)

	if v := a.Value.Children.ForName("directiveExt"); v != nil {
		v1, _ := v.Value(nil)
		res.DirectiveExt = helper.GetArrayOfInterface[string](v1)
	}
	if v := a.Value.Children.ForName("query"); v != nil {
		v1, _ := v.Value(nil)
		res.Query = getSqlCreateExtension(v1.(map[string]interface{}))
	}
	if v := a.Value.Children.ForName("get"); v != nil {
		v1, _ := v.Value(nil)
		res.Get = getSqlCreateExtension(v1.(map[string]interface{}))
	}
	return res
}

func customizeSqlBuilderMutation(a *ast.Argument) SQLDirectiveMutation {
	res := getDefaultFilledSqlBuilderMutation(true)
	if v := a.Value.Children.ForName("directiveExt"); v != nil {
		v1, _ := v.Value(nil)
		res.DirectiveExt = helper.GetArrayOfInterface[string](v1)
	}

	if v := a.Value.Children.ForName("add"); v != nil {
		v1, _ := v.Value(nil)
		res.Add = getSqlCreateExtension(v1.(map[string]interface{}))
	}
	if v := a.Value.Children.ForName("update"); v != nil {
		v1, _ := v.Value(nil)
		res.Update = getSqlCreateExtension(v1.(map[string]interface{}))
	}

	if v := a.Value.Children.ForName("delete"); v != nil {
		v1, _ := v.Value(nil)
		res.Delete = getSqlCreateExtension(v1.(map[string]interface{}))
	}
	return res
}
