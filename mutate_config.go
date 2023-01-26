package autogql

import (
	"fmt"
	"log"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/fasibio/autogql/structure"
	"github.com/vektah/gqlparser/v2/ast"
)

func (ggs *AutoGqlPlugin) MutateConfig(cfg *config.Config) error {
	log.Println("MutateConfig")

	cfg.Directives[string(structure.DirectiveSQL)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLPrimary)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLIndex)] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives[string(structure.DirectiveSQLGorm)] = config.DirectiveConfig{SkipRuntime: true}
	for k := range ggs.Handler.List {
		makeResolverFor := []string{fmt.Sprintf("Add%sPayload", k), fmt.Sprintf("Update%sPayload", k), fmt.Sprintf("Delete%sPayload", k)}
		for _, r := range makeResolverFor {
			e := cfg.Models[r]
			e.Fields = make(map[string]config.TypeMapField)
			e.Fields[templates.LcFirst(k)] = config.TypeMapField{
				Resolver: true,
			}
			cfg.Models[r] = e
		}

	}
	return ggs.remapInputType2Type(cfg)
	// return nil
}
func ConstraintFieldHook(ggs *AutoGqlPlugin) func(td *ast.Definition, fd *ast.FieldDefinition, f *modelgen.Field) (*modelgen.Field, error) {
	return func(td *ast.Definition, fd *ast.FieldDefinition, f *modelgen.Field) (*modelgen.Field, error) {
		if o, ok := ggs.Handler.List[td.Name]; ok {
			if o.HasSqlDirective() {
				for _, e := range o.Entities {
					if e.Name() == fd.Name {
						if e.IsPrimary() {
							f.Tag += " gorm:\"primaryKey\""
						}
						if fd.Directives.ForName(string(structure.DirectiveSQLIndex)) != nil {
							f.Tag += " gorm:\"index\""
						}
						if d := fd.Directives.ForName(string(structure.DirectiveSQLGorm)); d != nil {
							f.Tag += fmt.Sprintf(" gorm:\"%s\"", d.Arguments.ForName("value").Value.Raw)
						}
						if !e.BuiltIn && !e.GqlTypeObj().HasSqlDirective() {
							f.Tag += " gorm:\"-\""
						}
					}
				}
			}
		}
		return f, nil
	}
}
func MutateHook(ggs *AutoGqlPlugin) func(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	return func(b *modelgen.ModelBuild) *modelgen.ModelBuild {

		return b
	}

}

func (ggs *AutoGqlPlugin) remapInputType2Type(cfg *config.Config) error {
	return nil
}

func getModelStruct(cfg *config.Config, name string) string {
	return fmt.Sprintf("%s/model.%s", cfg.Resolver.ImportPath(), name)
}
