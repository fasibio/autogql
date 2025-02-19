package autogql

import (
	"fmt"
	"log"
	"strings"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/fasibio/autogql/helper"
	"github.com/fasibio/autogql/structure"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *AutoGqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")

	cfg.Directives[string(structure.DirectiveSQL)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLPrimary)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLIndex)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLGorm)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveNoMutation)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLInputTypeTags)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLInputTypeTags.InternalName())] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLInputTypeDirective)] = config.DirectiveConfig{SkipRuntime: true}
	for objectKey, object := range ggs.Handler.List {
		makeResolverFor := []string{fmt.Sprintf("Add%sPayload", objectKey), fmt.Sprintf("Update%sPayload", objectKey), fmt.Sprintf("Delete%sPayload", objectKey)}
		for _, r := range makeResolverFor {
			e := cfg.Models[r]
			e.Fields = make(map[string]config.TypeMapField)
			e.Fields[templates.LcFirst(objectKey)] = config.TypeMapField{
				Resolver: true,
			}
			cfg.Models[r] = e
		}
		for _, entity := range object.Entities {
			if entity.HasGormDirective() && entity.GormDirectiveValue() == "-" {
				m := cfg.Models[object.Name()]
				if m.Fields == nil {
					m.Fields = make(map[string]config.TypeMapField)
				}
				m.Fields[entity.Name()] = config.TypeMapField{
					FieldName: entity.Name(),
					Resolver:  true,
				}
				cfg.Models[object.Name()] = m
			}
		}
	}
	return ggs.remapInputType2Type(cfg)
}
func ConstraintFieldHook(ggs *AutoGqlPlugin) func(td *ast.Definition, fd *ast.FieldDefinition, f *modelgen.Field) (*modelgen.Field, error) {
	return func(td *ast.Definition, fd *ast.FieldDefinition, f *modelgen.Field) (*modelgen.Field, error) {
		if f, err := modelgen.DefaultFieldMutateHook(td, fd, f); err != nil {
			return f, err
		}
		if inputTagDirective := fd.Directives.ForName(structure.DirectiveSQLInputTypeTags.InternalName()); inputTagDirective != nil {
			v, err := inputTagDirective.Arguments.ForName("value").Value.Value(nil)
			if err != nil {
				return f, err
			}
			val := helper.GetArrayOfInterface[string](v)
			f.Tag += fmt.Sprintf(" %s", strings.Join(val, "; "))

		}
		if o, ok := ggs.Handler.List[td.Name]; ok {
			if o.HasSqlDirective() {
				for _, e := range o.Entities {
					if e.Name() == fd.Name {
						var sb strings.Builder
						sb.WriteString(" gorm:\"")
						if e.IsPrimary() {
							sb.WriteString("primaryKey;")
						}
						if fd.Directives.ForName(string(structure.DirectiveSQLIndex)) != nil || (e.IsSoftDeletedAtEntry()) {
							sb.WriteString("index;")
						}
						if d := fd.Directives.ForName(string(structure.DirectiveSQLGorm)); d != nil {
							sb.WriteString(d.Arguments.ForName("value").Value.Raw + ";")
						}
						if shouldIgnoredByGorm(&e) {
							sb.WriteString("-;")
						}
						sb.WriteRune('"')
						if s := sb.String(); strings.Compare(strings.Trim(s, " "), "gorm:\"\"") != 0 {
							f.Tag += s
						}
					}
				}
			}
		}
		return f, nil
	}
}

func shouldIgnoredByGorm(e *structure.Entity) bool {
	return !e.IsPrimitive() && e.RawObject.Kind != ast.Enum && !e.GqlTypeObj().HasSqlDirective() && e.RawObject.Kind != ast.Scalar
}

func MutateHook(ggs *AutoGqlPlugin) func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	return func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
		return b
	}

}

func (ggs *AutoGqlPlugin) remapInputType2Type(cfg *config.Config) error {
	return nil
}
