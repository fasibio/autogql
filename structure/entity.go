package structure

import (
	"fmt"
	"strings"

	"github.com/fasibio/autogql/helper"
	"github.com/huandu/xstrings"
	"github.com/vektah/gqlparser/v2/ast"
)

type Entity struct {
	BuiltIn    bool
	Raw        *ast.FieldDefinition
	RawObject  *ast.Definition
	TypeObject *Object
}

func (e Entity) Name() string {
	return e.Raw.Name
}

func (e Entity) IsUnion() bool {
	return e.RawObject.Kind == ast.Union
}

func (e Entity) HasGormDirective() bool {
	return e.GormDirectiveValue() != ""
}

func (e Entity) HasNoMutationDirective() bool {
	d := e.Raw.Directives.ForName(string(DirectiveNoMutation))
	return d != nil
}

func (e Entity) GormDirectiveValue() string {
	d := e.Raw.Directives.ForName(string(DirectiveSQLGorm))
	if d == nil {
		return ""
	}
	return d.Arguments.ForName("value").Value.Raw
}

func (e Entity) DatabaseFieldName() string {
	if e.HasGormDirective() {
		value, err := helper.GetGormValue(e.GormDirectiveValue(), "column")
		if err == nil {
			return value
		}
	}
	return xstrings.ToSnakeCase(e.Name())
}

func (e Entity) HasMany2ManyDirective() bool {
	return e.HasGormDirective() && strings.Contains(e.GormDirectiveValue(), "many2many:")
}

func (e Entity) Many2ManyDirectiveTable() string {
	if !e.HasMany2ManyDirective() {
		return ""
	}
	gs := strings.Split(e.GormDirectiveValue(), ";")
	table := ""
	for _, gd := range gs {
		if strings.Contains(gd, "many2many:") {
			table = strings.Replace(gd, "many2many:", "", 1)
		}
	}
	return table
}

func (e Entity) IsSoftDeletedAtEntry() bool {
	return e.Name() == "deletedAt" && e.RawObject.Name == "SoftDelete"
}

func (e Entity) Ignore() bool {
	if e.HasGormDirective() {
		return e.GormDirectiveValue() == "-"
	}
	if e.IsSoftDeletedAtEntry() {
		return true
	}

	return false
}

func (e Entity) GqlTypeName() string {
	return e.Raw.Type.Name()
}

func (e Entity) GqlTypePrimaryDataType() string {
	return e.GqlTypeObj().PrimaryKeyField().GqlTypeName()
}

func (e Entity) GqlTypeObj() *Object {
	return e.TypeObject
}

func (e Entity) GqlType(suffix string) string {
	if e.RawObject.Kind == ast.Enum || e.RawObject.Kind == ast.Scalar {
		suffix = ""
	}
	name := e.Raw.Type.Name()

	if e.IsPrimitive() {
		if e.IsArray() {
			return fmt.Sprintf("[%s!]", name)
		}
		return name
	}
	if e.IsArray() {
		return fmt.Sprintf("[%s%s!]", name, suffix)
	}
	return fmt.Sprintf("%s%s", name, suffix)
}

func (e *Entity) Required() bool {
	return e.Raw.Type.NonNull
}

func (e Entity) RequiredChar() string {
	requiredChar := ""
	if e.Required() {
		requiredChar = "!"
	}
	return requiredChar
}

func (e *Entity) IsArray() bool {
	return e.Raw.Type.String()[0] == '['
}
func (e *Entity) IsArrayElementRequired() bool {
	if !e.IsArray() {
		return false
	}
	return e.Raw.Type.Elem.NonNull
}

func (e *Entity) IsEnum() bool {
	return e.RawObject.Kind == ast.Enum
}

func (e *Entity) HasInputTypeTagsDirective() bool {
	return e.InputTypeTagsDirective() != nil
}

func (e *Entity) InputTypeTagsDirective() []string {
	d := e.Raw.Directives.ForName(string(DirectiveSQLInputTypeTags))
	if d == nil {
		return nil
	}
	a := d.Arguments.ForName("value")
	v, _ := a.Value.Value(nil)
	switch v.(type) {
	case []interface{}:
		{
			res := helper.GetArrayOfInterface[string](v)
			return res
		}
	case interface{}:
		return []string{v.(string)}
	}
	return nil
}

func (e *Entity) HasInputTypeDirective() bool {
	return e.InputTypeDirective() != nil
}

// returns the Directive string to extends to template
func (e *Entity) InputTypeDirectiveGql() string {
	tags := e.InputTypeTagsDirective()
	inputTags := ""
	if tags != nil {
		t := make([]string, len(tags))
		for i, v := range tags {
			t[i] = strings.ReplaceAll(v, "\"", "\\\"")
		}

		inputTags = fmt.Sprintf("@%s(value: [\"%s\"])", DirectiveSQLInputTypeTags.InternalName(), strings.Join(t, ","))
	}

	inputDirectives := e.InputTypeDirective()
	inputDirectivesStr := ""
	if inputDirectives != nil {
		inputDirectivesStr = strings.Join(inputDirectives, " ")
	}

	return fmt.Sprintf("%s %s", inputTags, inputDirectivesStr)
}

func (e *Entity) InputTypeDirective() []string {
	d := e.Raw.Directives.ForName(string(DirectiveSQLInputTypeDirective))
	if d == nil {
		return nil
	}
	a := d.Arguments.ForName("value")
	v, _ := a.Value.Value(nil)
	switch v.(type) {
	case []interface{}:
		{
			res := helper.GetArrayOfInterface[string](v)
			return res
		}
	case interface{}:
		return []string{v.(string)}
	}
	return nil
}

func (e *Entity) IsPrimitive() bool {
	return e.BuiltIn && e.RawObject.Kind == ast.Scalar
}

func (e *Entity) IsScalar() bool {
	return e.RawObject.Kind == ast.Scalar
}

func (e *Entity) IsPrimary() bool {
	return e.Raw.Directives.ForName(string(DirectiveSQLPrimary)) != nil
}

func (e *Entity) IsIndex() bool {
	return e.Raw.Directives.ForName(string(DirectiveSQLIndex)) != nil
}

func (e *Entity) WhereAble() bool {
	switch e.Raw.Type.Name() {
	case "String", "DateTime", "Int", "Float", "ID", "Boolean":
		return true
	}
	return false
}

func (e *Entity) OrderAble() bool {
	if e.Ignore() {
		return false
	}
	switch e.Raw.Type.Name() {
	case "String", "DateTime", "Int", "Float", "ID", "Boolean":
		return true
	}
	return false
}
