package runtimehelper

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/graphql"
	"github.com/fasibio/autogql/structure"
	"github.com/huandu/xstrings"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

type Clausel = string

const (
	Clausel_Eq           Clausel = "eq"
	Clausel_Eqi          Clausel = "eqi"
	Clausel_Ne           Clausel = "ne"
	Clausel_StartsWith   Clausel = "startsWith"
	Clausel_EndsWith     Clausel = "endsWith"
	Clausel_Contains     Clausel = "contains"
	Clausel_NotContains  Clausel = "notContains"
	Clausel_NotContainsI Clausel = "notContainsi"
	Clausel_ContainsI    Clausel = "containsi"
	Clausel_Null         Clausel = "null"
	Clausel_NotNull      Clausel = "notNull"
	Clausel_IN           Clausel = "in"
	Clausel_NotIN        Clausel = "noIn"
)

type PreloadFields struct {
	Fields      []string
	TableName   string
	SubTables   []PreloadFields
	PreloadName string
}

func GetPreloadsMap(ctx context.Context, tableName string) PreloadFields {
	return GetNestedPreloadsMap(graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		tableName, "")
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloadsMap(ctx *graphql.OperationContext, fields []graphql.CollectedField, tableName, parentTableName string) PreloadFields {
	res := PreloadFields{
		TableName: tableName,

		Fields: make([]string, 0),
	}

	ownIdAdded := false

	for _, column := range fields {
		// prefixColumn := GetPreloadString(prefix, column.Name)
		if column.Name == "__typename" || column.Name == "__schema" || column.Name == "__type" { // to remove buildIn fields
			continue
		}
		if !ownIdAdded {
			res.Fields = append(res.Fields, GetDbIdFields(column.ObjectDefinition, ""))
			ownIdAdded = true
		}

		for _, a := range column.ObjectDefinition.Fields {
			if a.Name == parentTableName+"ID" {
				res.Fields = append(res.Fields, xstrings.ToSnakeCase(a.Name))
			}
		}
		if len(column.Field.SelectionSet) != 0 { // To remove all parent objects
			if res.SubTables == nil {
				res.SubTables = make([]PreloadFields, 0)
			}
			res.Fields = append(res.Fields, GetDbIdFields(column.ObjectDefinition, column.Name))
			tmp := GetNestedPreloadsMap(ctx, graphql.CollectFields(ctx, column.Selections, nil), column.Field.Definition.Type.Name(), strings.ToLower(tableName))
			tmp.PreloadName = column.Name
			res.SubTables = append(res.SubTables, tmp)

		} else if !ShouldFieldBeIgnored(column.ObjectDefinition, column.Name) {
			res.Fields = append(res.Fields, xstrings.ToSnakeCase(column.Name))
		}

	}
	// res.Fields = append(res.Fields, GetDbIdFields())
	res.Fields = removeDuplicateStr(res.Fields)
	return res
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if item == "" {
			continue
		}
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func ShouldFieldBeIgnored(d *ast.Definition, fieldName string) bool {
	for _, v := range d.Fields {
		tmp := structure.Entity{
			BuiltIn: false,
			Raw:     v,
		}
		if tmp.Name() == fieldName {
			return tmp.Ignore()
		}
	}
	return false
}

func GetDbIdFields(d *ast.Definition, fieldName string) string {
	entities := make([]structure.Entity, len(d.Fields))

	for i, v := range d.Fields {
		entities[i] = structure.Entity{
			BuiltIn: false,
			Raw:     v,
		}
	}

	o := structure.Object{
		Raw:      d,
		Entities: entities,
	}

	if !o.HasSqlDirective() {
		return ""
	}

	if fieldName == "" {
		return o.PrimaryKeyField().Name()
	}

	return o.ForeignNameKeyName(fieldName)

}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}
func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

func GetPreloadSelection(ctx context.Context, dbObj *gorm.DB, data PreloadFields) *gorm.DB {
	return GetNestedPreloadSelection(data, dbObj)
}

func GetNestedPreloadSelection(data PreloadFields, dbObj *gorm.DB) *gorm.DB {
	fields := make([]string, len(data.Fields))
	for i, v := range data.Fields {
		fields[i] = fmt.Sprintf("%s.%s", dbObj.Config.NamingStrategy.TableName(data.TableName), v)
	}
	res := dbObj.Select(fields)
	if data.SubTables != nil {
		for _, t := range data.SubTables {
			res = addPreloadCondition(t, res)
		}
	}
	return res
}

func addPreloadCondition(pFields PreloadFields, d *gorm.DB) *gorm.DB {
	return d.Preload(templates.UcFirst(pFields.PreloadName), func(db *gorm.DB) *gorm.DB {
		return GetNestedPreloadSelection(pFields, db)
	})
}
