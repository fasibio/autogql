package autogql

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/Masterminds/sprig/v3"

	_ "embed"

	"github.com/fasibio/autogql/structure"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *AutoGqlPlugin) InjectSourceLate(schema *ast.Schema) *ast.Source {
	log.Println("InjectSourceLate")
	builderHelper := structure.NewSqlBuilderHelper()
	for _, c := range schema.Types {
		if sqlDirective := c.Directives.ForName(string(structure.DirectiveSQL)); sqlDirective != nil {
			object := structure.NewObject(c)
			a := make(structure.SqlBuilderList)
			e := getSqlBuilderFields(c.Fields, schema, a)
			object.Entities = e
			for k, v := range a {
				if v.Raw.Kind == ast.Scalar {
					continue
				}
				if _, ok := builderHelper.List[k]; !ok {
					builderHelper.List[k] = v
				}
			}
			builderHelper.List[object.Name()] = &object
		}
	}
	result := getExtendsSource(builderHelper)
	ggs.Handler = builderHelper
	return &ast.Source{
		Name:    fmt.Sprintf("%s/autogql.graphql", ggs.Name()),
		Input:   result,
		BuiltIn: false,
	}
}

func fillSqlBuilderByName(schema *ast.Schema, name string, knownValues structure.SqlBuilderList) {

	val := schema.Types[name]
	if val.BuiltIn {
		return
	}
	if _, isOk := knownValues[val.Name]; isOk {
		return
	} else {
		tmp := structure.NewObject(val)
		knownValues[val.Name] = &tmp
		f := getSqlBuilderFields(val.Fields, schema, knownValues)
		tmp.Entities = f

	}
}

func getSqlBuilderFields(fields ast.FieldList, schema *ast.Schema, knownValues structure.SqlBuilderList) []structure.Entity {
	res := make([]structure.Entity, 0)
	for _, field := range fields {
		var m *structure.Object = nil
		if !schema.Types[field.Type.Name()].BuiltIn {
			tmp := structure.NewObject(schema.Types[field.Type.Name()])
			m = &tmp
		}
		tempE := structure.Entity{
			BuiltIn:    schema.Types[field.Type.Name()].BuiltIn,
			TypeObject: m,
			Raw:        field,
			RawObject:  schema.Types[field.Type.Name()],
		}
		res = append(res, tempE)
		fillSqlBuilderByName(schema, field.Type.Name(), knownValues)
	}
	return res
}

//go:embed inject_source_late.gql.go.tpl
var gqltemplate string

func getExtendsSource(builder structure.SqlBuilderHelper) string {
	funcs := make(map[string]any)

	codegenTemplatesFunc := templates.Funcs()
	sprigTemplatesFunc := sprig.FuncMap()
	for k, v := range codegenTemplatesFunc {
		funcs[k] = v
	}
	for k, v := range sprigTemplatesFunc {
		funcs[k] = v
	}

	tmpl, _ := template.New("sourcebuilder").Funcs(funcs).Parse(gqltemplate)
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, builder)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
